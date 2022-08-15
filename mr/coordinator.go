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

func (c *Coordinator) waitMapTask(i int) {
	// time.Sleep(10 * time.Second)
	timer := time.NewTimer(10 * time.Second)
	<-timer.C

	// FIrst thing I'll do is switch from sleep to timer
	// THen what I'll do is I'll put console swithces to check for bugs

	log.Printf("Waited 10 seconds for map task %d at %v", i, time.Now())

	c.l.Lock()
	defer c.l.Unlock()

	log.Printf("Attempting to write to map task %d at %v", i, time.Now())
	if (c.mapTasks[i].Status == pb.StatusType_Processing) {
		c.mapTasks[i].Status = pb.StatusType_Incomplete
	}
}

func (c *Coordinator) waitReduceTask(i int) {
	// time.Sleep(10 * time.Second)
	timer := time.NewTimer(10 * time.Second)
	<-timer.C

	log.Printf("Waited 10 seconds for map task %d at %v", i, time.Now())

	c.l.Lock()
	defer c.l.Unlock()

	log.Printf("Attempting to write to map task %d at %v", i, time.Now())
	if (c.reduceTasks[i].Status == pb.StatusType_Processing) {
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

				log.Printf("Finished giving map task %d at %v", i, time.Now())

				taskCopy := pb.MapTask{}
				copier.Copy(&taskCopy, &c.mapTasks[i])
				return &pb.TaskReply{Type: 1, MapTask: &taskCopy}, nil
			}
		}
	}
	for _, mapTask := range c.mapTasks {
		if (mapTask.GetStatus() == pb.StatusType_Processing) {
			return &pb.TaskReply{Type: 0}, nil
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
				return &pb.TaskReply{Type: 2, ReduceTask: &taskCopy}, nil
			}
		}
	}
	for _, reduceTask := range c.reduceTasks {
		if (reduceTask.GetStatus() == pb.StatusType_Processing) {
			return &pb.TaskReply{Type: 0}, nil
		}
	}
	c.isReduceComplete = true

	return &pb.TaskReply{Type: 0}, nil
}

func (c *Coordinator) CompleteTask(ctx context.Context, in *pb.TaskComplete) (*empty.Empty, error) {
	c.l.Lock()
	defer c.l.Unlock()

	log.Printf("coordinator, complete task %d", in.GetNum())
	if (in.GetType() == 1) {
		c.mapTasks[in.GetNum()].Status = pb.StatusType_Complete
	} else if (in.GetType() == 2) {
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
