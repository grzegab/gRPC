package main

import (
	"context"
	pb "github.com/grzegab/gRPC/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("start greet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Greg",
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Print("Greeting:", res.Result)
}
