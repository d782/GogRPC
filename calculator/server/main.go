package main

import (
	"log"
	"net"

	pb "github.com/d782/GogRPC/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	log.Printf("Listen on %s \n", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v ", err)
	}

}
