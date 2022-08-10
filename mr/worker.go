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

	"mr_system/pb"
)

func Worker(mapf func(string, string) []*pb.KeyValue,
	reducef func(string, []string) string) {
	conn := getConn()
	c := pb.NewCoordinatorClient(conn)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Infinite loop until all tasks are done. Must let workers know somehow

	r, err := c.GetTask(ctx, &pb.TaskRequest{})
	if err != nil {
		log.Fatalf("could not retrieve task %v", err)
	}

	kva := getKeyValuesFromFile(mapf, r.GetFileName())
	organizedKvaChunks := make([][]*pb.KeyValue, r.GetNReduce())
	for _, kv := range kva {
		reduceNum := ihash(kv.Key) % r.GetNReduce()
		organizedKvaChunks[reduceNum] = append(organizedKvaChunks[reduceNum], kv)
	}
	for reduceNum := int32(0); reduceNum < r.GetNReduce(); reduceNum++ {
		serializeKvaChunk(organizedKvaChunks[reduceNum], r.GetMapNum(), reduceNum)
	}

	// // Try to read one of them out
	// in, err := ioutil.ReadFile(getFileName(r.GetMapNum(), 0))
	// if err != nil {
	// 	log.Fatalf("failed to read key values file: %v", err)
	// }
	// kvaReduce := &pb.KeyValuesFile{}
	// if err := proto.Unmarshal(in, kvaReduce); err != nil {
	// 	log.Fatalf("failed to decode key values file: %v", err)
	// }
	// for _, kv := range kvaReduce.KeyValues {
	// 	log.Printf("%v %v", kv.Key, kv.Value)
	// }
}

func getConn() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func getKeyValuesFromFile(mapf func(string, string) []*pb.KeyValue, fileName string) []*pb.KeyValue {
	log.Printf("reading file: %s", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("cannot open %v", fileName)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", fileName)
	}
	defer file.Close()
	return mapf(fileName, string(content))
}

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int32(h.Sum32() & 0x7fffffff)
}

func serializeKvaChunk(kvaChunk []*pb.KeyValue, mapNum int32, reduceNum int32) {
	getFileName := func(mapNum int32, reduceNum int32) string {
		return fmt.Sprintf("mr_temp/mr-%d-%d", mapNum, reduceNum)
	}

	f, err := os.Create(getFileName(mapNum, reduceNum))
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	kvaChunkFile := &pb.KeyValuesFile{KeyValues: kvaChunk}
	out, err := proto.Marshal(kvaChunkFile)
	if err != nil {
		log.Fatalf("failed to encode key values file: %v", err)
	}
	if err := ioutil.WriteFile(getFileName(mapNum, reduceNum), out, 0644); err != nil {
		log.Fatalf("failed to write key values file: %v", err)
	}
}