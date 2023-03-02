package main

import (
	"context"
	pb "gRPC-Golang/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	response, err := client.Greet(context.Background(), &pb.GreetRequest{FirstName: "Lucas"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", response.Result)
}
