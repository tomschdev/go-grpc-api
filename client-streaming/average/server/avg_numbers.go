package main

import (
	"io"
	"log"

	pb "github.com/tomschdev/go-grpc-api/client-streaming/average/proto"
)

// this procedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) AvgNumbers(stream pb.GreetService_AvgNumbersServer) error {
	log.Printf("AvgNumbers function was invoked")

	res := 0
	count := 0
	for {
		//receive all incoming request
		//only when EOF is recieved does the server respond with generated payload
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float32(res / count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		res += int(req.Number)
		count++
	}
	return nil
}
