package main

import (
	"context"
	pb "github.com/grzegab/gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

var collection *mongo.Collection
var addr = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("failed to create mongo client db: %v", err)
	}
	client.Connect(context.Background())
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
