package main

import (
	"context"
	pb "github.com/grzegab/gRPC/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("call client. update blog was invokled")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Greg 2",
		Title:    "Blog 2",
		Content:  "Content of first one modified",
	}

	res, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatal("Error! These was: ", err)
	}

	log.Printf("CreateBlog updated: %v\n", res)
}
