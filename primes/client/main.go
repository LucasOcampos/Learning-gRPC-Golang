package main

import (
	pb "gRPC-Golang/primes/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "localhost:50051"

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer connection.Close()

	client := pb.NewPrimeServiceClient(connection)

	primeDecomposition(client)
}
