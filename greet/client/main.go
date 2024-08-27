package main

import (
	"context"
	pb "github.com/grzegab/gRPC/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	//conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//
	//c := pb.NewGreetServiceClient(conn)
	//
	//doGreet(c)

	c, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer c.Close()

	cs := pb.NewCalcServiceClient(c)
	doSum(cs)
}

//func doGreet(c pb.GreetServiceClient) {
//	log.Println("start greet invoked")
//	res, err := c.Greet(context.Background(), &pb.GreetRequest{
//		FirstName: "Greg",
//	})
//
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//
//	log.Print("Greeting:", res.Result)
//}

func doSum(c pb.CalcServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		First:  12.4,
		Second: 13.5,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Result of the sum: %s", res.Result)
}
