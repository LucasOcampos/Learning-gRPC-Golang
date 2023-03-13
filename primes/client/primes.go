package main

import (
	"context"
	pb "gRPC-Golang/primes/proto"
	"io"
	"log"
)

func primeDecomposition(client pb.PrimeServiceClient) {
	log.Println("primeDecomposition was invoked")

	request := &pb.PrimeRequest{Number: 120}

	stream, err := client.PrimeDecomposition(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("PrimeDecomposition: %d\n", message.Result)
	}
}
