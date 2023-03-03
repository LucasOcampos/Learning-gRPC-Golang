package main

import (
	"fmt"
	pb "gRPC-Golang/greet/proto"
	"log"
)

func (server *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)
	for i := 0; i < 10; i++ {
		response := fmt.Sprintf("Hello %s, number %d", in.FirstName, i+1)
		stream.Send(&pb.GreetResponse{Result: response})
	}

	return nil
}
