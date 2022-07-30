package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/tomschdev/go-grpc-api/advanced/errorhandle/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// this procedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function was invoked with: %v\n", in)
	operand := in.Number

	if operand > 0 {
		return &pb.SqrtResponse{
			Result: math.Sqrt(float64(operand)),
		}, nil
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %d", operand),
		)
	}
}
