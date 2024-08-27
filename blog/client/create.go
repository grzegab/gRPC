package main

import (
	"context"
	pb "github.com/grzegab/gRPC/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("call client.CreateBlog")

	blog := &pb.Blog{
		AuthorId: "Greg",
		Title:    "Blog 1",
		Content:  "Content of first one",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatal("Error! These was: ", err)
	}

	log.Println("CreateBlog created:", res.Id)

	return res.Id
}
