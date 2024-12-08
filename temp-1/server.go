package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"example" 
)


type server struct {
	example.UnimplementedGreeterServer
}


func (s *server) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &example.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	
	s := grpc.NewServer()

	
	example.RegisterGreeterServer(s, &server{})


	log.Printf("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
