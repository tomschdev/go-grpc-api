package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/greet/proto"
)

// this precedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet function was invoked with: %v\n", in)
	return &pb.GreetResponse{
		Result: "Howzit " + in.Firstname,
	}, nil
}
