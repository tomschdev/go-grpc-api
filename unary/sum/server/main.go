package main

import (
	"log"
	"net"

	pb "github.com/tomschdev/go-grpc-api/unary/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.SumServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("server failed to serve: %v\n", err)
	}
}
