package main

import (
	"context"
	"io"
	"log"

	pb "github.com/d782/GogRPC/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("do primes invoked")

	req := &pb.PrimeRequest{
		Number: 1234,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while Calling primes %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Printf("primes: %d\n", res.Result)
	}
}
