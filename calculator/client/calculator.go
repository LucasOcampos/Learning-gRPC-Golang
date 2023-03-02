package main

import (
	"context"
	pb "gRPC-Golang/calculator/proto"
	"log"
)

func calculate(client pb.CalculatorServiceClient) {
	log.Println("calculate was invoked")

	response, err := client.Calculator(context.Background(), &pb.CalculatorRequest{
		Number1: 10,
		Number2: 3,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v", err)
	}

	log.Printf("Calculated: %v", response.Result)
}
