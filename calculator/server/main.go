package main

import (
	pb "github.com/grzegab/gRPC/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.CalcOperationServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())

	server := grpc.NewServer()
	pb.RegisterCalcOperationServer(server, &Server{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Operation(req *pb.OperationRequest, res grpc.ServerStreamingServer[pb.OperationResponse]) error {

	return nil
}

//type Server struct {
//	pb.NnwServiceServer
//}

//func main() {
//	l, e := net.Listen("tcp", addr)
//	if e != nil {
//		log.Fatal("listen error:", e)
//	}
//
//	log.Println("listening on", addr)
//
//	s := grpc.NewServer()
//	pb.RegisterNnwServiceServer(s, &Server{})
//
//	err := s.Serve(l)
//	if err != nil {
//		log.Fatal("failed to serve")
//	}
//}
//
//func (s *Server) Nnw(in *pb.NnwRequest, stream grpc.ServerStreamingServer[pb.NnwResponse]) error {
//	log.Println("Nnw called")
//
//	num := in.Number
//	div := int64(2)
//
//	for num > 1 {
//		if num%div == 0 {
//			stream.Send(&pb.NnwResponse{
//				Result: div,
//			})
//
//			num /= div
//		} else {
//			div++
//		}
//	}
//
//	return nil
//}
