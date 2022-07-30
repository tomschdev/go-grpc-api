package main

import (
	"log"

	pb "github.com/tomschdev/go-grpc-api/advanced/ssl/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	options := []grpc.DialOption{}
	if tls {
		certFile := "cert/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("error while loading CA trust certificate: %v\n", err)
		}

		options = append(options, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, options...)
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	} else {
		log.Printf("connected: %v\n", conn)
	}
	defer conn.Close() // executes at end of this block of code
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
}
