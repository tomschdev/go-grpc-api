package main

import (
	"io"
	"log"
	"math"

	pb "github.com/tomschdev/go-grpc-api/bi-directional-streaming/max/proto"
)

// this procedure is implemented on server side and directly called from client side, with access provided via stub (auto generated code in proto dir)
func (s *Server) Max(stream pb.MaxService_MaxServer) error {
	log.Printf("Max function was invoked")

	var max int32 = math.MinInt32
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
		if int32(req.Number) > int32(max) {
			max = req.Number
		}
		err = stream.Send(&pb.MaxResponse{
			Result: int32(max),
		})
		if err != nil {
			log.Fatalf("Error while sending response to client: %v\n", err)
		}
	}
}
