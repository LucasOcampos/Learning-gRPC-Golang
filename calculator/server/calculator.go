package main

import (
	"context"
	pb "gRPC-Golang/calculator/proto"
	"log"
)

func (server *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.CalculatorResponse{Result: in.Number1 + in.Number2}, nil
}
