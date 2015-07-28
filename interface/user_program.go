package mapreduce

import (
	"log"
	"os/exec"
	"strings"

	pb "../proto"
	"google.golang.org/grpc"
)

func (t *workerTask) getNewMapperUserServer(address string) pb.MapperStreamClient {
	t.logger.Println(address)
	conn, err := grpc.Dial(address)
	if err != nil {
		t.logger.Fatalf("did not connect: %v", err)
	}
	return pb.NewMapperStreamClient(conn)
}

func (t *workerTask) getNewReducerUserServer(address string) pb.ReducerStreamClient {
	conn, err := grpc.Dial(address)
	if err != nil {
		t.logger.Fatalf("did not connect: %v", err)
	}
	return pb.NewReducerStreamClient(conn)
}

func (t *workerTask) startNewUserServer(cmdline []string) {
	for i := 0; i < len(cmdline); i++ {
		parts := strings.Fields(cmdline[i])
		background := parts[0]
		head := parts[1]
		parts = parts[2:len(parts)]
		cmd := exec.Command(head, parts...)
		t.logger.Println(head, parts)
		if background == "b" {
			err := cmd.Start()
			log.Println("background", head, parts, err)
			if err != nil {
				t.logger.Fatalln("background ", head, parts, err)
			}
		} else if background == "wc" || background == "ww" {
			err := cmd.Start()
			if err != nil {
				t.logger.Fatal("wait", head, parts, err)
			}
			err = cmd.Wait()
			if err != nil {
				if background == "ww" {
					t.logger.Fatal(err)
				} else {
					t.logger.Println(err)
				}
			}
		} else {
			err := cmd.Run()
			if err != nil {
				t.logger.Fatalln("run", head, parts, err)
			}
		}
	}

}
