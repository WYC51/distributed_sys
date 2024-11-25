package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/WYC51/gRPC_Server_Client/proto"
	"google.golang.org/grpc"
)

type MessageService struct {
	pb.UnimplementedGreeterServer // starting from v1.20.0 needing to add this line
}

func (m *MessageService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Server say hello to " + in.GetName()}, nil
}

const (
	port = ":5050"
)

func main() {
	fmt.Printf("starting gRPC server on port %s\n", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &MessageService{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
