package main

import (
	"log"

	pb "github.com/tomschdev/go-grpc-api/advanced/errorhandle/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	} else {
		log.Printf("connected: %v\n", conn)
	}
	defer conn.Close() // executes at end of this block of code
	c := pb.NewSqrtServiceClient(conn)
	doSqrt(c, 10) //demonstrate success case
	log.Printf("Observe error handling after making request with illegal operand:\n")
	doSqrt(c, -2) //demomstrate error case
}
