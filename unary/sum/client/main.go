package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/unary/sum/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Printf("doSum invoked")
	res, err := c.Addition(context.Background(), &pb.SumRequest{
		X: 17,
		Y: 2,
	})

	if err != nil {
		log.Fatalf("error was not nil, rpc on server failed: %v\n", err)
	}

	log.Printf("addition: %v\n", res.Result)
}
