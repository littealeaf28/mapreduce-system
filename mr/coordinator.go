package mr

import (
	"net"
	"fmt"
	"log"
	"context"

	"google.golang.org/grpc"

	"mr_system/pb"
)

type Coordinator struct {
	pb.UnimplementedCoordinatorServer
	files []string
	nReduce int32
}

func (c *Coordinator) serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1234))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCoordinatorServer(s, c)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func (c *Coordinator) GetTask(ctx context.Context, in *pb.TaskRequest) (*pb.TaskReply, error) {
	log.Printf("received")
	// Keep track of task number. Should only give one task at a time
	return &pb.TaskReply{FileName: c.files[0], MapNum: 0, NReduce: c.nReduce}, nil
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	return ret
}

func MakeCoordinator(files []string, nReduce int32) *Coordinator {
	c := Coordinator{files: files, nReduce: nReduce}
	c.serve()
	return &c
}
