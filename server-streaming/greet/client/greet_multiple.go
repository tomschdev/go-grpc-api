package main

import (
	"context"
	"io"
	"log"

	pb "github.com/tomschdev/go-grpc-api/server-streaming/greet/proto"
)

func doGreetMultiple(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	req := &pb.GreetRequest{
		Firstname: "tomschdev",
	}
	// rpc call to server's GreetMultiple implementation
	// returns stream which receives responses
	// note: GreetMultiple definition on server has stream as a second paramater which is made implicit by stream keyword in proto file
	stream, err := c.GreetMultiple(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling GreetMultiple: %v\n", err)
	}

	for {
		// wait for stream payload to arrive
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading the stream: %v\n", err)
		}

		log.Printf("GreetMultiple: %s\n", msg.Result)
	}
}
