package main

import (
	"context"
	"fmt"
	pb "github.com/grzegab/gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {

	log.Printf("Received update blog: %v\n", in)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Converting id failed: %v", err))
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Updating blog failed: %v", err))
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Blog not found: %v", err))
	}

	return &emptypb.Empty{}, nil
}
