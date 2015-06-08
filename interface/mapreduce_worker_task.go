package mapreduce

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/coreos/go-etcd/etcd"
	"github.com/golang/protobuf/proto"
	pb "github.com/plutoshe/mr/proto"
	"github.com/taskgraph/taskgraph"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type workerTask struct {
	framework          taskgraph.Framework
	epoch              uint64
	logger             *log.Logger
	taskID             uint64
	etcdClient         *etcd.Client
	userSeverPort      uint64
	mapperWriteCloser  []bufio.Writer
	reducerWriteCloser bufio.Writer
	shuffleContainer   map[string][]string //store temporary shuffle results

	//channels
	epochChange               chan *mapreduceEvent
	dataReady                 chan *mapreduceEvent
	metaReady                 chan *mapreduceEvent
	notifyChan                chan *mapreduceEvent
	exitChan                  chan struct{}
	stopGrabTaskForEveryEpoch chan bool

	config MapreduceConfig
}

type mapperEmitKV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (t *workerTask) Init(taskID uint64, framework taskgraph.Framework) {
	t.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	t.taskID = taskID
	t.framework = framework
	t.etcdClient = etcd.NewClient(t.config.EtcdURLs)
	t.userSeverPort = 10000 + taskID

	//channel init
	t.stopGrabTaskForEveryEpoch = make(chan bool, 1)
	t.epochChange = make(chan *mapreduceEvent, 1)
	t.dataReady = make(chan *mapreduceEvent, 1)
	t.metaReady = make(chan *mapreduceEvent, 1)
	t.notifyChan = make(chan *mapreduceEvent, 1)
	t.exitChan = make(chan struct{})

	t.initializeTaskEnv()
	go t.run()
}

func (t *workerTask) run() {
	for {
		select {
		case ec := <-t.epochChange:
			go t.doEnterEpoch(ec.ctx, ec.epoch)

		case notify := <-t.notifyChan:
			t.framework.FlagMeta(notify.ctx, notify.linkType, notify.meta)

		case <-t.exitChan:
			return

		case dataReady := <-t.dataReady:
			go t.processWork(dataReady.ctx, dataReady.fromID, dataReady.method, dataReady.output)
		case <-t.metaReady:

		}
	}
}

func (t *workerTask) startNewUserServer(cmdline []string) {
	argv := []string{"-port", strconv.FormatUint(t.userSeverPort, 10)}
	cmd := exec.Command(cmdline[0], argv...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

}

func (t *workerTask) getNewMapperUserServer(address string) pb.MapperClient {
	t.logger.Println(address + fmt.Sprintf(":%d", t.userSeverPort))
	conn, err := grpc.Dial(address + fmt.Sprintf(":%d", t.userSeverPort))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewMapperClient(conn)
}

func (t *workerTask) getNewReducerUserServer(address string) pb.ReducerClient {
	conn, err := grpc.Dial(address + fmt.Sprintf(":%d", t.userSeverPort))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewReducerClient(conn)
}

func (t *workerTask) processWork(ctx context.Context, fromID uint64, method string, output proto.Message) {

	resp, ok := output.(*pb.WorkConfigResponse)
	t.logger.Println(resp)
	if !ok {
		t.logger.Panicf("doDataRead, corruption in proto.Message.")
	}
	if resp == nil {
		return
	}
	workID := t.getWorkID()
	if workID == "non" {
		t.logger.Println("Expect worker possessing a work")
		return
	}

	workConfig := WorkConfig{}
	for i := range resp.Key {
		switch resp.Key[i] {
		case "InputFilePath":
			workConfig.InputFilePath = strings.Split(resp.Value[i], delim)
		case "OutputFilePath":
			workConfig.OutputFilePath = strings.Split(resp.Value[i], delim)
		case "UserProgram":
			workConfig.UserProgram = strings.Split(resp.Value[i], delim)
		case "UserServerAddress":
			workConfig.UserServerAddress = resp.Value[i]
		case "WorkType":
			workConfig.WorkType = resp.Value[i]
		case "SupplyContent":
			workConfig.SupplyContent = strings.Split(resp.Value[i], delim)
		}
	}

	// start user grpc server by cmd line,
	go t.startNewUserServer(workConfig.UserProgram)

	t.logger.Println("begin process work", workConfig)
	t.logger.Println(workConfig.WorkType)
	// Determined by the work type, start relative processing procedure
	switch workConfig.WorkType {
	case "Mapper":
		userClient := t.getNewMapperUserServer(workConfig.UserServerAddress)
		go t.mapperProcedure(ctx, workID, workConfig, userClient)
	case "Reducer":
		userClient := t.getNewReducerUserServer(workConfig.UserServerAddress)
		go t.reducerProcedure(ctx, workID, workConfig, userClient)
	}
}

func (t *workerTask) initializeTaskEnv() error {
	_, err := t.etcdClient.Create(MapreduceTaskStatusPath(t.config.AppName, t.taskID, "workStatus"), "non", 0)
	if err != nil {
		if strings.Contains(err.Error(), "Key already exists") {
			return nil
		}
		return err
	}
	return nil
}

func (t *workerTask) datarequestForWork(ctx context.Context, method string) {
	t.logger.Println("data request")
	master := t.framework.GetTopology()["Master"].GetNeighbors(t.epoch)

	for _, node := range master {
		t.framework.DataRequest(ctx, node, method, &pb.WorkRequest{TaskID: t.taskID})
	}
}

func (t *workerTask) grabWork(ctx context.Context, method string, stop chan bool) {
	t.logger.Println("in grab work")
	t.datarequestForWork(ctx, method)

	// afterwards, watch etcd worker attribute "workStatus"
	// if exist "set' operation, and the value is "non"
	receiver := make(chan *etcd.Response, 1)
	t.logger.Println(MapreduceTaskStatusPath(t.config.AppName, t.taskID, "workStatus"))
	go t.etcdClient.Watch(MapreduceTaskStatusPath(t.config.AppName, t.taskID, "workStatus"), 0, true, receiver, stop)
	for resp := range receiver {
		if resp.Action != "set" {
			continue
		}
		if resp.Node.Value == "non" {
			t.datarequestForWork(ctx, method)
		} else {
			continue
		}
	}
}

func (t *workerTask) Emit(key, val string) {
	if t.config.ReducerNum == 0 {
		return
	}
	h := fnv.New32a()
	h.Write([]byte(key))
	var KV mapperEmitKV
	KV.Key = key
	KV.Value = val
	toShuffle := h.Sum32() % uint32(t.config.ReducerNum)
	data, err := json.Marshal(KV)
	data = append(data, '\n')
	if err != nil {
		t.logger.Fatalf("json marshal error : ", err)
	}
	t.mapperWriteCloser[toShuffle].Write(data)
}

func (t *workerTask) Collect(key string, val string) {
	t.reducerWriteCloser.Write([]byte(key + " " + val + "\n"))
}

func (t *workerTask) Clean(path string) {
	err := t.config.FilesystemClient.Remove(path)
	if err != nil {
		t.logger.Fatal(err)
	}
}

func (t *workerTask) emitKvPairs(userClient pb.MapperClient, str string, value string, stop bool) {
	stream, err := userClient.GetEmitResult(context.Background(), &pb.MapperRequest{Key: str, Value: value})
	if err != nil {
		t.logger.Fatalf("could not access the user program server : %v", err)
	}
	if !stop {
		for {
			feature, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.GetEmitResult, %v", userClient, err)
				return
			}
			t.Emit(feature.Key, feature.Value)
		}
	}
}

func (t *workerTask) getWorkID() string {
	requestWorkStatus, err := t.etcdClient.Get(MapreduceTaskStatusPath(t.config.AppName, t.taskID, "workStatus"), false, false)
	if err != nil {
		log.Fatal("etcdutil: can not get worker status from etcd")
	}
	return requestWorkStatus.Node.Value
}

func (t *workerTask) mapperProcedure(ctx context.Context, workID string, workConfig WorkConfig, userClient pb.MapperClient) {
	t.logger.Println("In mapper procedure")
	var i uint64
	t.mapperWriteCloser = make([]bufio.Writer, 0)

	for i = 0; i < t.config.ReducerNum; i++ {
		path := workConfig.OutputFilePath[0] + "/" + strconv.FormatUint(i, 10) + "from" + workID
		t.logger.Println("Output Path ", path)
		t.Clean(path)
		tmpWrite, err := t.config.FilesystemClient.OpenWriteCloser(path)
		if err != nil {
			t.logger.Printf("MapReduce : get mapreduce filesystem client writer failed, ", err)
		}
		t.mapperWriteCloser = append(t.mapperWriteCloser, *bufio.NewWriterSize(tmpWrite, t.config.WriterBufferSize))
	}

	// Input file loading
	for readFileID := 0; readFileID < len(workConfig.InputFilePath); readFileID++ {
		mapperReaderCloser, err := t.config.FilesystemClient.OpenReadCloser(workConfig.InputFilePath[readFileID])
		if err != nil {
			t.logger.Printf("MapReduce : get mapreduce filesystem client reader failed, ", err)
		}

		var str string
		bufioReader := bufio.NewReaderSize(mapperReaderCloser, t.config.ReaderBufferSize)

		for err != io.EOF {
			str, err = bufioReader.ReadString('\n')

			if err != io.EOF && err != nil {
				t.logger.Fatalf("MapReduce : mapper read Error, ", err)
			}
			if err != io.EOF {
				str = str[:len(str)-1]
			}
			t.emitKvPairs(userClient, str, "", false)
		}
		// stop the reader of corresponding file
		mapperReaderCloser.Close()
	}

	// stop user program grpc client
	t.emitKvPairs(userClient, "Stop", "Stop", true)

	//flush output result
	for i = 0; i < t.config.ReducerNum; i++ {
		t.mapperWriteCloser[i].Flush()
	}
	t.logger.Println("FileRead finished")

	// notify the master mapper work has been done
	t.notifyChan <- &mapreduceEvent{ctx: ctx, fromID: t.taskID, linkType: "Master", meta: "WorkFinished" + workID}
}

func (t *workerTask) processShuffleKV(str []byte) {
	var tp mapperEmitKV
	if err := json.Unmarshal([]byte(str), &tp); err == nil {
		t.shuffleContainer[tp.Key] = append(t.shuffleContainer[tp.Key], tp.Value)
	}
}

func (t *workerTask) collectKvPairs(userClient pb.ReducerClient, key string, value []string, stop bool) {
	r, err := userClient.GetCollectResult(context.Background(), &pb.ReducerRequest{Key: key, Value: value})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if !stop {
		for {
			feature, err := r.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.GetCollectResult, %v", userClient, err)
				return
			}
			t.Collect(feature.Key, feature.Value)
		}
	}
}

func (t *workerTask) reducerProcedure(ctx context.Context, workID string, workConfig WorkConfig, userClient pb.ReducerClient) {
	for ProcessID := 0; ProcessID < len(workConfig.InputFilePath); ProcessID++ {
		t.shuffleContainer = make(map[string][]string)

		arg := strings.Split(workConfig.SupplyContent[ProcessID], " ")

		mapperWorkSum, err := strconv.ParseUint(arg[0], 10, 64)
		if err != nil {
			t.logger.Fatalf("Failed to get argv mapperWorkSum : %v", err)
		}
		t.logger.Println(workConfig.InputFilePath)
		t.logger.Println(workConfig.OutputFilePath)
		for i := uint64(0); i < mapperWorkSum; i++ {
			shufflePath := workConfig.InputFilePath[0] + "/" + arg[1] + "from" + strconv.FormatUint(i, 10)
			shuffleReadCloser, err := t.config.FilesystemClient.OpenReadCloser(shufflePath)
			t.logger.Println("get shuffle data from ", shufflePath)
			if err != nil {
				if strings.Contains(err.Error(), "The specified blob does not exist") {
					continue
				}
				t.logger.Fatalf("MapReduce : get azure storage client failed, ", err)
			}
			bufioReader := bufio.NewReaderSize(shuffleReadCloser, t.config.ReaderBufferSize)
			var str []byte
			err = nil
			for err != io.EOF {
				str, err = bufioReader.ReadBytes('\n')
				if err != io.EOF && err != nil {
					t.logger.Fatalf("MapReduce : Shuffle read Error, ", err)
				}
				if err != io.EOF {
					str = str[:len(str)-1]
				}
				t.processShuffleKV(str)
			}
		}
		reducerOutputPath := workConfig.OutputFilePath[0]
		t.Clean(reducerOutputPath)

		reducerWriteCloser, err := t.config.FilesystemClient.OpenWriteCloser(reducerOutputPath)
		if err != nil {
			t.logger.Fatalf("MapReduce : get reducer writer error, %v", err)
		}

		t.reducerWriteCloser = *bufio.NewWriterSize(reducerWriteCloser, t.config.WriterBufferSize)

		for k := range t.shuffleContainer {
			t.collectKvPairs(userClient, k, t.shuffleContainer[k], false)
		}
		t.reducerWriteCloser.Flush()
	}

	t.collectKvPairs(userClient, "Stop", []string{}, true)
	t.logger.Println("finshed reducer")
	t.notifyChan <- &mapreduceEvent{ctx: ctx, epoch: t.epoch, linkType: "Master", meta: "WorkFinished" + workID}
	// for i := uint64(0); i < mapperWorkSum; i++ {
	// 	shufflePath := workConfig.InputFilePath[i] + "/" + arg[1] + "from" + strconv.FormatUint(i, 10)
	// 	t.Clean(shufflePath)
	// }
}

func (t *workerTask) EnterEpoch(ctx context.Context, epoch uint64) {
	t.epochChange <- &mapreduceEvent{ctx: ctx, epoch: epoch}
}

func (t *workerTask) doEnterEpoch(ctx context.Context, epoch uint64) {
	// stop the last epoch grab work procedure
	// start a new one
	close(t.stopGrabTaskForEveryEpoch)
	t.stopGrabTaskForEveryEpoch = make(chan bool, 1)
	go t.grabWork(ctx, "/proto.Master/GetWork", t.stopGrabTaskForEveryEpoch)
}

func (t *workerTask) Exit() {
	close(t.stopGrabTaskForEveryEpoch)
	close(t.exitChan)
}

func (t *workerTask) CreateServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterWorkerServer(server, t)
	return server

}

func (t *workerTask) CreateOutputMessage(method string) proto.Message {
	switch method {
	case "/proto.Master/GetWork":
		return new(pb.WorkConfigResponse)
	}
	panic("")
}

func (t *workerTask) DataReady(ctx context.Context, fromID uint64, method string, output proto.Message) {
	t.dataReady <- &mapreduceEvent{ctx: ctx, fromID: fromID, method: method, output: output}
}

func (t *workerTask) MetaReady(ctx context.Context, fromID uint64, LinkType, meta string) {
	t.metaReady <- &mapreduceEvent{ctx: ctx, fromID: fromID, linkType: LinkType, meta: meta}
}