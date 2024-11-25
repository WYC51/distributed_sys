package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "github.com/WYC51/gRPC_Server_Client/proto"
)

const (
	address = "localhost:5050"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := "Liam"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}