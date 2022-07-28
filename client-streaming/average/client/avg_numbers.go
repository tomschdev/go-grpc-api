package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tomschdev/go-grpc-api/client-streaming/average/proto"
)

func doAvgNumbers(c pb.GreetServiceClient) {
	log.Println("doAvgNumbers was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 3},
		{Number: 56},
		{Number: 23},
	}
	stream, err := c.AvgNumbers(context.Background())
	if err != nil {
		log.Fatalf("error calling AvgNumbers: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from AvgNumbers after stream of requests: %v\n", err)
	}
	log.Printf("Average: %s\n", res.Result)
}
