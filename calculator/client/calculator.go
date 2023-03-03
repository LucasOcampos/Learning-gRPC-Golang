package main

import (
	"context"
	pb "gRPC-Golang/calculator/proto"
	"log"
)

func calculateSum(client pb.SumServiceClient) {
	log.Println("calculateSum was invoked")

	response, err := client.Sum(context.Background(), &pb.SumRequest{
		Number1: 10,
		Number2: 3,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v", err)
	}

	log.Printf("Calculated Sum: %v", response.Result)
}
