package main

import (
	"context"
	"log"

	pb "github.com/d782/GogRPC/project/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet func was called with %v", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
