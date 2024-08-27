package main

import "context"
import pb "github.com/grzegab/gRPC/greet/proto"

func (s *Server) Sum(ctx context.Context, sum *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: sum.First + sum.Second}, nil
}
