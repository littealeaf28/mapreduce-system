package mr

import (
	"net"
	"fmt"
	"log"
	"context"
	"sync"

	"google.golang.org/grpc"
	"mr_system/pb"
	"github.com/golang/protobuf/ptypes/empty"
)

// Could do read write lock
type Coordinator struct {
	pb.UnimplementedCoordinatorServer
	l sync.Mutex
	mapTasks []*pb.MapTask
	isMapComplete bool
	reduceTasks []*pb.ReduceTask
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

func (c *Coordinator) GetTask(ctx context.Context, in *empty.Empty) (*pb.TaskReply, error) {
	c.l.Lock()
	defer c.l.Unlock()

	if (!c.isMapComplete) {
		for _, mapTask := range c.mapTasks {
			if (!mapTask.GetStatus()) {
				return &pb.TaskReply{Type: 1, MapTask: mapTask}, nil
			}
		}
	}
	c.isMapComplete = true

	for _, reduceTask := range c.reduceTasks {
		if (!reduceTask.GetStatus()) {
			return &pb.TaskReply{Type: 2, ReduceTask: reduceTask}, nil
		}
	}

	return &pb.TaskReply{Type: 0}, nil
}

func (c *Coordinator) CompleteTask(ctx context.Context, in *pb.TaskComplete) (*empty.Empty, error) {
	c.l.Lock()
	defer c.l.Unlock()

	if (in.GetType() == 1) {
		c.mapTasks[in.GetNum()].Status = true
	} else if (in.GetType() == 2) {
		c.reduceTasks[in.GetNum()].Status = true
	}

	return &empty.Empty{}, nil
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	// c.l.Lock()
	// defer c.l.Unlock()

	// for _, mapTask := range c.mapTasks {
	// 	if (!mapTask.GetStatus()) {
	// 		return false
	// 	}
	// }
	return true
}

func MakeCoordinator(files []string, nReduce int32) *Coordinator {
	mapTasks := make([]*pb.MapTask, 0)
	for i, file := range files {
		mapTasks = append(mapTasks, &pb.MapTask{Status: false, FileName: file, MapNum: int32(i), NReduce: nReduce})
	}

	reduceTasks := make([]*pb.ReduceTask, nReduce)
	for i := 0; i < len(reduceTasks); i++ {
		reduceTasks[i] = &pb.ReduceTask{Status: false, ReduceNum: int32(i)}
	}

	c := Coordinator{isMapComplete: false, mapTasks: mapTasks, reduceTasks: reduceTasks}

	c.serve()
	return &c
}
