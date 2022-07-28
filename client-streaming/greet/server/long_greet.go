package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/tomschdev/go-grpc-api/client-streaming/greet/proto"
)

// this procedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""
	for {
		//receive all incoming request
		//only when EOF is recieved does the server respond with generated payload
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		res += fmt.Sprintf("Hello %s\n", req)

	}
	return nil
}
