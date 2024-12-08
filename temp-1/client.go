package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

	"example" 
)

func main() {
	
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	
	c := example.NewGreeterClient(conn)

	
	req := &example.HelloRequest{Name: "World"}
	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	
	fmt.Printf("Greeting: %s\n", resp.GetMessage())
}
