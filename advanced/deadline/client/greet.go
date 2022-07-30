package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tomschdev/go-grpc-api/advanced/deadline/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	//set a timeout that invokes cancel function on timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		Firstname: "tomschdev",
	})

	if err != nil {
		e, std := status.FromError(err)
		if std {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("server exceeded client deadline")
			} else {
				log.Fatalf("unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("error was not nil, rpc on server failed: %v\n", err)
		}
	} else {
		log.Printf("greeting: %s\n", res.Result)
	}

}
