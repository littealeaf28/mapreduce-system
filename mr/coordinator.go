package mr

import (
	"net"
	"fmt"
	"log"
	"context"
	"sync"
	"time"
	"os"

	"google.golang.org/grpc"
	"mr_system/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
)

type Coordinator struct {
	pb.UnimplementedCoordinatorServer
	l sync.RWMutex
	isMapComplete bool
	isReduceComplete bool
	mapTasks []*pb.MapTask
	reduceTasks []*pb.ReduceTask
}

func coordSocket() string {
	return fmt.Sprintf("/var/tmp/824-mr-%d", os.Getuid())
}

func (c *Coordinator) serve() {
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1234))4
	sockname := coordSocket()
	os.Remove(sockname)
	lis, err := net.Listen("unix", sockname)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCoordinatorServer(s, c)

	log.Printf("server listening at %v", lis.Addr())

	// Hmm: not sure best way to resolve errors if server go routine fails
	// errs := make(chan error, 1)
	// go func() { errs <- s.Serve(lis) }()
	// if err := <-errs; err != nil {
	// 	log.Fatalf("failed to server: %v", err)
	// }
	go s.Serve(lis)
}

// could use generics to simplify code
func (c *Coordinator) waitMapTask(i int) {
	// time.Sleep(10 * time.Second)
	timer := time.NewTimer(10 * time.Second)
	<-timer.C

	c.l.Lock()
	defer c.l.Unlock()

	if (c.mapTasks[i].Status == pb.StatusType_Processing) {
		log.Printf("Retrying map task %d", i)
		c.mapTasks[i].Status = pb.StatusType_Incomplete
	}
}

func (c *Coordinator) waitReduceTask(i int) {
	// time.Sleep(10 * time.Second)
	timer := time.NewTimer(10 * time.Second)
	<-timer.C

	c.l.Lock()
	defer c.l.Unlock()

	if (c.reduceTasks[i].Status == pb.StatusType_Processing) {
		log.Printf("Retrying reduce task %d", i)
		c.reduceTasks[i].Status = pb.StatusType_Incomplete
	}
}

func (c *Coordinator) GetTask(ctx context.Context, in *empty.Empty) (*pb.TaskReply, error) {
	c.l.Lock()
	defer c.l.Unlock()

	if (!c.isMapComplete) {
		for i, mapTask := range c.mapTasks {
			if (mapTask.GetStatus() == pb.StatusType_Incomplete) {
				c.mapTasks[i].Status = pb.StatusType_Processing

				go c.waitMapTask(i)

				taskCopy := pb.MapTask{}
				copier.Copy(&taskCopy, &c.mapTasks[i])
				return &pb.TaskReply{Type: pb.TaskType_Map, MapTask: &taskCopy}, nil
			}
		}
	}
	for _, mapTask := range c.mapTasks {
		if (mapTask.GetStatus() == pb.StatusType_Processing) {
			return &pb.TaskReply{Type: pb.TaskType_Empty}, nil
		}
	}
	c.isMapComplete = true

	if (!c.isReduceComplete) {
		for i, reduceTask := range c.reduceTasks {
			if (reduceTask.GetStatus() == pb.StatusType_Incomplete) {
				c.reduceTasks[i].Status = pb.StatusType_Processing
				
				go c.waitReduceTask(i)

				taskCopy := pb.ReduceTask{}
				copier.Copy(&taskCopy, &c.reduceTasks[i])
				return &pb.TaskReply{Type: pb.TaskType_Reduce, ReduceTask: &taskCopy}, nil
			}
		}
	}
	for _, reduceTask := range c.reduceTasks {
		if (reduceTask.GetStatus() == pb.StatusType_Processing) {
			return &pb.TaskReply{Type: pb.TaskType_Empty}, nil
		}
	}
	c.isReduceComplete = true

	return &pb.TaskReply{Type: pb.TaskType_Empty}, nil
}

func (c *Coordinator) CompleteTask(ctx context.Context, in *pb.TaskComplete) (*empty.Empty, error) {
	c.l.Lock()
	defer c.l.Unlock()

	log.Printf("coordinator, complete task %d", in.GetNum())
	if (in.GetType() == pb.TaskType_Map) {
		c.mapTasks[in.GetNum()].Status = pb.StatusType_Complete
	} else if (in.GetType() == pb.TaskType_Reduce) {
		c.reduceTasks[in.GetNum()].Status = pb.StatusType_Complete
	}

	return &empty.Empty{}, nil
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	c.l.RLock()
	defer c.l.RUnlock()

	return c.isMapComplete && c.isReduceComplete
}

func MakeCoordinator(files []string, nReduce int32) *Coordinator {
	mapTasks := make([]*pb.MapTask, 0)
	for i, file := range files {
		mapTasks = append(mapTasks, &pb.MapTask{Status: pb.StatusType_Incomplete, FileName: file, MapNum: int32(i), NReduce: nReduce})
	}

	reduceTasks := make([]*pb.ReduceTask, nReduce)
	for i := 0; i < len(reduceTasks); i++ {
		reduceTasks[i] = &pb.ReduceTask{Status: pb.StatusType_Incomplete, ReduceNum: int32(i)}
	}

	c := Coordinator{isMapComplete: false, isReduceComplete: false, mapTasks: mapTasks, reduceTasks: reduceTasks}

	c.serve()
	return &c
}
