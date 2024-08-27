package main

import (
	"context"
	pb "github.com/grzegab/gRPC/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("call client.ReadBlog")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatal("Error! These was: ", err)
	}

	log.Println("Blog found:", res.Id)

	return res
}
