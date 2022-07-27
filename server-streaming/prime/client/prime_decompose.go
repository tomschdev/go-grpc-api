package main

import (
	"context"
	"io"
	"log"

	pb "github.com/tomschdev/go-grpc-api/server-streaming/prime/proto"
)

func doPrimeDecomposition(c pb.PrimeServiceClient) {
	log.Printf("doPrimeDecomposition has been invoked")
	req := &pb.PrimeRequest{
		Number: 120,
	}
	stream, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling PrimeDecompostion: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)
		}
		log.Printf("Prime: %s\n", msg.Prime)
	}
}
