package main

import (
	pb "gRPC-Golang/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const address = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", err)
	}

	log.Printf("Listening on %s\n", address)

	server := grpc.NewServer()
	pb.RegisterSumServiceServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
