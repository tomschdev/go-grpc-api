package main

import (
	"log"
	"net"

	pb "github.com/tomschdev/go-grpc-api/server-streaming/prime/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50054"

type Server struct {
	pb.PrimeServiceServer
}

func main() {
	//listen on addr
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}
	log.Printf("listening on: %s\n", addr)

	//create grpc server
	s := grpc.NewServer()

	//register grpc server with proto service
	pb.RegisterPrimeServiceServer(s, &Server{})

	//serve server at addr
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
