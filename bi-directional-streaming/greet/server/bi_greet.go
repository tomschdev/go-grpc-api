package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/tomschdev/go-grpc-api/bi-directional-streaming/greet/proto"
)

// this procedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) BiGreet(stream pb.GreetService_BiGreetServer) error {
	log.Printf("BiGreet function was invoked")

	for {
		//receive all incoming request
		//only when EOF is recieved does the server respond with generated payload
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		res := fmt.Sprintf("Hello %s\n", req)
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error while sending response to client: %v\n", err)
		}
	}
}
