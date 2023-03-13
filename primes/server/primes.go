package main

import (
	"fmt"
	pb "gRPC-Golang/primes/proto"
	"log"
)

func (server *Server) PrimeDecomposition(in *pb.PrimeRequest, stream pb.PrimeService_PrimeDecompositionServer) error {
	log.Printf("PrimeDecomposition function was invoked with %v\n", in)
	var factor int32
	factor = 2
	number := in.Number

	for number > 1 {
		if number%factor == 0 {
			fmt.Printf("Number: %d.  Factor: %d", number, factor)
			stream.Send(&pb.PrimeResponse{Result: factor})

			number = number / factor
		} else {
			factor++
		}
	}

	return nil
}
