package main

import (
	"context"
	pb "go-experiments/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %s", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})

	log.Println("gRPC server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
