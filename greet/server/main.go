package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/grzegab/gRPC/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
	pb.CalcServiceServer
}

func main() {
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}

	log.Println("listening on", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	pb.RegisterCalcServiceServer(s, &Server{})

	if err := s.Serve(l); err != nil {
		log.Fatal("failed to serve")
	}
}
