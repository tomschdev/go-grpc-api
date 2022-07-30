package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tomschdev/go-grpc-api/advanced/deadline/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// this precedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet with deadline function was invoked with: %v\n", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request: deadline exceeded.")
			return nil, status.Error(codes.Canceled, "The client canceled the request.")
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Howzit " + in.Firstname,
	}, nil
}
