package main

import (
	pb "github.com/grzegab/gRPC/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	log.Printf("Created blog with id: %v", id)
	readBlog(c, id)
	//readBlog(c, "fakeID")
	updateBlog(c, id)
	listBlogs(c)
	deleteBlog(c, id)
}
