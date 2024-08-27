package main

import (
	"context"
	pb "github.com/grzegab/gRPC/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("Deleting Blog %s\n", id)

	_, err := c.DeletBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Failed to delete Blog: %v\n", err)
	}

	log.Println("Blog deleted")
}
