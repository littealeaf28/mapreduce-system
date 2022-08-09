package mr

import (
	"log"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type KeyValue struct {
	Key string
	Value string
}

func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {
	connect()
}

func connect() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewCoordinatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetTask(ctx, &TaskRequest{})
	if err != nil {
		log.Fatalf("could not greet %v", err)
	}
	log.Printf("greeting: %s", r.GetFileName())
}