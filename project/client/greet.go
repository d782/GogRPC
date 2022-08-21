package main

import (
	"context"
	"log"

	pb "github.com/d782/GogRPC/project/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("do greet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Diego",
	})

	if err != nil {
		log.Fatalf("failed %v", err)
	}

	log.Printf("Greeting %v", res.Result)
}
