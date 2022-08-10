package mr

import (
	"log"
	"context"
	"time"
	"hash/fnv"
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"sort"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"

	"mr_system/pb"
)

type ByKey []*pb.KeyValue

func (a ByKey) Len() int { return len(a) }
func (a ByKey) Swap(i int, j int) { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i int, j int) bool { return a[i].Key < a[j].Key }

func Worker(mapf func(string, string) []*pb.KeyValue,
	reducef func(string, []string) string) {
	conn := getConn()
	c := pb.NewCoordinatorClient(conn)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for {
		r, _ := c.GetTask(ctx, &empty.Empty{})
		// if err != nil {
		// 	log.Fatalf("could not retrieve task %v", err)
		// }

		// Switch
		if r.GetType() == 0 {
			break
		} else if (r.GetType() == 1) {
			task := r.GetMapTask()

			kva := getKeyValuesFromFile(mapf, task.GetFileName())
			kvaBuckets := make([][]*pb.KeyValue, task.GetNReduce())
			for _, kv := range kva {
				reduceNum := ihash(kv.Key) % task.GetNReduce()
				kvaBuckets[reduceNum] = append(kvaBuckets[reduceNum], kv)
			}
			for reduceNum := int32(0); reduceNum < task.GetNReduce(); reduceNum++ {
				serializeKvaBucket(kvaBuckets[reduceNum], task.GetMapNum(), reduceNum)
			}

			log.Printf("completed map task for %v", task.GetFileName())
			c.CompleteTask(ctx, &pb.TaskComplete{Type: 1, Num: task.GetMapNum()})
		} else if (r.GetType() == 2) {
			task := r.GetReduceTask()

			kvaBucketFileNames, err := filepath.Glob(fmt.Sprintf("mr_temp/mr-*-%d", task.GetReduceNum()))
			if err != nil {
				log.Fatalf("could not find key value bucket files: %v", err)
			}

			kvaOutput := make([]*pb.KeyValue, 0)
			for _, kvaBucketFileName := range kvaBucketFileNames {
				in, err := ioutil.ReadFile(kvaBucketFileName)
				if err != nil {
					log.Fatalf("failed to read key values file: %v", err)
				}
				kvaReduce := &pb.KeyValuesFile{}
				if err := proto.Unmarshal(in, kvaReduce); err != nil {
					log.Fatalf("failed to decode key values file: %v", err)
				}
				kvaOutput = append(kvaOutput, kvaReduce.KeyValues...)
			}

			sort.Sort(ByKey(kvaOutput))

			oname := fmt.Sprintf("mr_temp/mr-out-%d", task.GetReduceNum())
			ofile, err := os.Create(oname)
			if err != nil {
				log.Fatalf("could not create output file: %v", err)
			}
			defer ofile.Close()

			i := 0
			for i < len(kvaOutput) {
				j := i + 1
				for j < len(kvaOutput) && kvaOutput[j].Key == kvaOutput[i].Key {
					j++
				}
				values := []string{}
				for k := i; k < j; k++ {
					values = append(values, kvaOutput[k].Value)
				}
				output := reducef(kvaOutput[i].Key, values)

				fmt.Fprintf(ofile, "%v %v\n", kvaOutput[i].Key, output)

				i = j
			}

			log.Printf("completed reduce task %d", task.GetReduceNum())
			c.CompleteTask(ctx, &pb.TaskComplete{Type: 2, Num: task.GetReduceNum()})
		}
	}
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

func serializeKvaBucket(kvaBucket []*pb.KeyValue, mapNum int32, reduceNum int32) {
	getFileName := func(mapNum int32, reduceNum int32) string {
		return fmt.Sprintf("mr_temp/mr-%d-%d", mapNum, reduceNum)
	}

	f, err := os.Create(getFileName(mapNum, reduceNum))
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	kvaBucketFile := &pb.KeyValuesFile{KeyValues: kvaBucket}
	out, err := proto.Marshal(kvaBucketFile)
	if err != nil {
		log.Fatalf("failed to encode key values file: %v", err)
	}
	if err := ioutil.WriteFile(getFileName(mapNum, reduceNum), out, 0644); err != nil {
		log.Fatalf("failed to write key values file: %v", err)
	}
}