package main

import (
	pb "github.com/grzegab/gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"authorId"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(bi *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       bi.ID.Hex(),
		AuthorId: bi.AuthorId,
		Title:    bi.Title,
		Content:  bi.Content,
	}
}
