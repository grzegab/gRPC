package main

import (
	"context"
	pb "github.com/grzegab/gRPC/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("call client.List blogs")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("ListBlogs(err) = _, %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListBlogs(err) = fatal, %v", err)
		}

		log.Println(res)
	}
}
