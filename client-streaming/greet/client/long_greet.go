package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tomschdev/go-grpc-api/client-streaming/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{Firstname: "tomschdev"},
		{Firstname: "stevedev"},
		{Firstname: "momdev"},
		{Firstname: "daddev"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet after stream of requests: %v\n", err)
	}
	log.Printf("doLongGreet result: \n%s\n", res.Result)
}
