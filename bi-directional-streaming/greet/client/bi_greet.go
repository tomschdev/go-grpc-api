package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/tomschdev/go-grpc-api/bi-directional-streaming/greet/proto"
)

func doBiGreet(c pb.GreetServiceClient) {
	log.Println("doBiGreet was invoked")

	reqs := []*pb.GreetRequest{
		{Firstname: "tomschdev"},
		{Firstname: "stevedev"},
		{Firstname: "momdev"},
		{Firstname: "daddev"},
	}

	stream, err := c.BiGreet(context.Background())
	if err != nil {
		log.Fatalf("error calling BiGreet: %v\n", err)
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
