package main

import (
	"log"

	pb "github.com/d782/GogRPC/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v \n", err)
	}

	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	doPrimes(c)
	defer conn.Close()
}
