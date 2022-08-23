package main

import (
	"context"
	"io"
	"log"

	pb "github.com/d782/GogRPC/project/proto"
)

func doGreeManyTimes(c pb.GreetServiceClient) {
	log.Println("do many times invoked!")

	req := &pb.GreetRequest{
		FirstName: "Diego",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error  while calling GreetManyTimes")
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading the stream")
		}

		log.Printf("Stream resolved %s", msg.Result)
	}
}
