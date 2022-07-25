package main

import (
	"fmt"
	"log"

	pb "github.com/tomschdev/go-grpc-api/server-streaming/greet/proto"
)

// this precedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) GreetMultiple(in *pb.GreetRequest, stream pb.GreetService_GreetMultipleServer) error {
	log.Printf("greet function was invoked with: %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %v (%s)", in.Firstname, i)
		// received by stream.Recv on client side
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
