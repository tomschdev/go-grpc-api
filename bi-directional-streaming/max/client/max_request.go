package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/tomschdev/go-grpc-api/bi-directional-streaming/max/proto"
)

func doMax(c pb.MaxServiceClient) {
	log.Println("doBiGreet was invoked")

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 3},
		{Number: 9},
		{Number: 3},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error calling Max: %v\n", err)
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving message from client: %v\n", err)
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
