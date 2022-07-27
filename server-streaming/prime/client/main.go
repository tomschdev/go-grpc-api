package main

import (
	"log"

	pb "github.com/tomschdev/go-grpc-api/server-streaming/prime/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50054"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	} else {
		log.Printf("connected: %v\n", conn)
	}
	defer conn.Close()
	c := pb.NewPrimeServiceClient(conn)
	doPrimeDecomposition(c)
}
