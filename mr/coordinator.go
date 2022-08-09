package mr

import (
	"net"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type Coordinator struct {
	UnimplementedCoordinatorServer
}

func (c *Coordinator) serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1234))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterCoordinatorServer(s, &Coordinator{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	return ret
}

func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	c.serve()
	return &c
}
