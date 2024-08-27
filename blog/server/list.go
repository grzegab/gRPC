package main

import (
	"context"
	pb "github.com/grzegab/gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
	log.Println("Received list blogs")

	cur, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		bi := &BlogItem{}
		err := cur.Decode(bi)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		stream.Send(documentToBlog(bi))
	}

	if err := cur.Err(); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
