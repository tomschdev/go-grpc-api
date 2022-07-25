package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/unary/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	// below is the remote procedure call i.e.,
	// we've received a GreetServiceClient (from proto dir) as a parameter
	// with our client, we directly call a procedure (OOP style) that is implemented and invoked on the server (Greet)
	// grpc takes care of the networking
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		Firstname: "tomschdev",
	})

	if err != nil {
		log.Fatalf("error was not nil, rpc on server failed: %v\n", err)
	}

	log.Printf("greeting: %s\n", res.Result)
}
