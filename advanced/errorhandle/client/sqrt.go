package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/advanced/errorhandle/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.SqrtServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	// below is the remote procedure call i.e.,
	// we've received a GreetServiceClient (from proto dir) as a parameter
	// with our client, we directly call a procedure (OOP style) that is implemented and invoked on the server (Greet)
	// grpc takes care of the networking
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, std := status.FromError(err)
		if std {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Printf("Input is probably a negative number or NaN")
			}
		} else {
			log.Fatalf("error was not nil, rpc on server failed: %v\n", err)
		}
	} else {
		log.Printf("Sqrt result: %f\n", res.Result)
	}

}
