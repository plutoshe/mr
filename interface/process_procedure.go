package mapreduce

import (
	"io"
	"log"
	"strconv"

	pb "../proto"

	"golang.org/x/net/context"
)

// mapperProcedure is called by the worker node to process mapper work.
// It only transmits fundamental mapper work configuration to corresponding user processing grpc server
func (t *workerTask) mapperProcedure(ctx context.Context, workID string, workConfig WorkConfig, userClient pb.MapperStreamClient) {
	t.logger.Println("In mapper procedure")

	waitc := make(chan struct{})
	stream, err := userClient.GetStreamEmitResult(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// grpc receieve endpoint
	go func() {
		for {
			_, err := stream.Recv()

			if err != nil && err != io.EOF {
				log.Fatalf("Failed to receive a note : %v", err)
				return
			}
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
		}
	}()

	var buf []*pb.KvPair

	// Input file loading
	for readFileID := 0; readFileID < len(workConfig.InputFilePath); readFileID++ {
		buf = append(buf, &pb.KvPair{Key: "InputFile", Value: workConfig.InputFilePath[readFileID]})
		buf = append(buf, &pb.KvPair{Key: "ReducerNum", Value: strconv.FormatUint(t.config.ReducerNum, 10)})
		err := stream.Send(&pb.MapperRequest{buf})
		if err != nil {
			t.logger.Fatalln(err)
		}
		buf = nil
	}

	// stop user program grpc client
	stream.CloseSend()
	t.logger.Println("Send finshed")

	<-waitc

	t.logger.Println("Mapper work finished")

	// notify the master mapper work has been done
	t.notifyChan <- &mapreduceEvent{ctx: ctx, fromID: t.taskID, linkType: "Master", meta: "WorkFinished" + workID}
}

// reducerProcedure is called by worker node to process reducer work
// It only transmits fundamental reducer work configuration to corresponding user process grpc server
func (t *workerTask) reducerProcedure(ctx context.Context, workID string, workConfig WorkConfig, userClient pb.ReducerStreamClient) {
	waitc := make(chan struct{})
	stream, err := userClient.GetStreamCollectResult(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			_, err := stream.Recv()

			if err != nil && err != io.EOF {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			if err == io.EOF {
				close(waitc)
				return
			}
		}
	}()

	var buf []*pb.KvsPair
	buf = append(buf, &pb.KvsPair{Key: "InputPath", Value: workConfig.InputFilePath})
	buf = append(buf, &pb.KvsPair{Key: "Outputpaht", Value: workConfig.OutputFilePath})
	buf = append(buf, &pb.KvsPair{Key: "SupplentConfig", Value: workConfig.SupplyContent})
	err = stream.Send(&pb.ReducerRequest{buf})
	if err != nil {
		t.logger.Fatalln(err)
	}
	stream.CloseSend()
	t.logger.Println("Send finshed")
	<-waitc

	t.logger.Println("Reducer work finished")
	t.notifyChan <- &mapreduceEvent{ctx: ctx, epoch: t.epoch, linkType: "Master", meta: "WorkFinished" + workID}
}
