package main

import (
	"log"
	"net"

	pb "github.com/tomschdev/go-grpc-api/unary/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)

	// start up a vanilla gRPC server
	s := grpc.NewServer()
	// register a GreetService with required params (according to auto generated proto code) of gRPC server and Server struct
	// that contains a pb.GreetServiceServer
	pb.RegisterGreetServiceServer(s, &Server{})

	// start the gRPC server (that is now registered as a GreetService with pb) on the address we're listening to
	// i.e., direct requests to that address to this server
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
