package main

import (
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Failed to start server %v", err)
	}

	log.Printf("Listen on %s \n", addr)
}
