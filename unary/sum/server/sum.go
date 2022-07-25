package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/unary/sum/proto"
)

func (s *Server) Addition(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("addition function invoked with: %v\n", in)
	return &pb.SumResponse{
		Result: in.X + in.Y,
	}, nil
}
