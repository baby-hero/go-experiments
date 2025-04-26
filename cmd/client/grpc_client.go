package main

import (
	"context"
	"log"
	"time"

	pb "go-experiments/pb"

	"google.golang.org/grpc"
)

func main() {
	// Set up gRPC dial options
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure()) // Use TLS in production!

	// Dial the server
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create client stub
	client := pb.NewGreeterClient(conn)

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the SayHello RPC
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go Learner"})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}

	log.Printf("Response from server: %s", resp.GetMessage())
}
