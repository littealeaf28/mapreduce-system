package mr

import (
	"log"
	"context"
	"time"
	"hash/fnv"
	"fmt"

	"os"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int32(h.Sum32() & 0x7fffffff)
}

func getKeyValueFileName(mapTaskNo int32, reduceNo int32) string {
	return fmt.Sprintf("mr_temp/mr-%d-%d", mapTaskNo, reduceNo)
}

func Worker(mapf func(string, string) []*KeyValue,
	reducef func(string, []string) string) {
	connect(mapf)
	// I think I just have to defer here..
}

func connect(mapf func(string, string) []*KeyValue) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewCoordinatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Infinite loop until all tasks are done. Must let workers know somehow

	// Assuming just working with map tasks
	r, err := c.GetTask(ctx, &TaskRequest{})
	if err != nil {
		log.Fatalf("could not retrieve task %v", err)
	}
	log.Printf("file to read: %s", r.GetFileName())

	fileName := r.GetFileName()

	// Want to better understand how to defer - refactor

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("cannot open %v", fileName)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", fileName)
	}
	file.Close()
	kva := mapf(fileName, string(content))

	intermediate := make([][]*KeyValue, r.GetNReduce())
	for _, kv := range kva {
		reduceNo := ihash(kv.Key) % r.GetNReduce()
		intermediate[reduceNo] = append(intermediate[reduceNo], kv)
	}

	// Create files for map task
	for i := int32(0); i < r.GetNReduce(); i++ {
		os.Create(getKeyValueFileName(r.GetMapTaskNo(), i))

		kvaReduce := &KeyValuesFile{KeyValues: intermediate[i]}
		out, err := proto.Marshal(kvaReduce)
		if err != nil {
			log.Fatalf("failed to encode key values file: %v", err)
		}
		if err := ioutil.WriteFile(getKeyValueFileName(r.GetMapTaskNo(), i), out, 0644); err != nil {
			log.Fatalf("failed to write key values file: %v", err)
		}
	}

	// Try to read one of them out
	in, err := ioutil.ReadFile(getKeyValueFileName(r.GetMapTaskNo(), 0))
	if err != nil {
		log.Fatalf("failed to read key values file: %v", err)
	}
	kvaReduce := &KeyValuesFile{}
	if err := proto.Unmarshal(in, kvaReduce); err != nil {
		log.Fatalf("failed to decode key values file: %v", err)
	}
	for _, kv := range kvaReduce.KeyValues {
		log.Printf("%v %v", kv.Key, kv.Value)
	}
}