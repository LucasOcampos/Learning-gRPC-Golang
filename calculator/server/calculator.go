package main

import (
	"context"
	pb "gRPC-Golang/calculator/proto"
	"log"
)

func (server *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{Result: in.Number1 + in.Number2}, nil
}
