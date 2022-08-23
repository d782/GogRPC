package main

import (
	"context"
	"log"

	pb "github.com/d782/GogRPC/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatal("Could not sum")
	}

	log.Printf("Sum: %d", res.Result)
}
