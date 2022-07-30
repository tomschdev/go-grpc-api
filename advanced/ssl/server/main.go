package main

import (
	"log"
	"net"

	pb "github.com/tomschdev/go-grpc-api/advanced/ssl/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	options := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "cert/server.crt"
		keyFile := "cert/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed loading certificates: %v\n", err)

		}

		options = append(options, grpc.Creds(creds))
	}
	// start up a vanilla gRPC server
	s := grpc.NewServer(options...)
	// register a GreetService with required params (according to auto generated proto code) of gRPC server and Server struct
	// that contains a pb.GreetServiceServer
	pb.RegisterGreetServiceServer(s, &Server{})

	// start the gRPC server (that is now registered as a GreetService with pb) on the address we're listening to
	// i.e., direct requests to that address to this server
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
