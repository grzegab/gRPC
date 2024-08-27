package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grzegab/gRPC/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet Received: %v with %v", time.Now(), in)

	return &pb.GreetResponse{
		Result: "Heelo" + in.FirstName,
	}, nil
}
