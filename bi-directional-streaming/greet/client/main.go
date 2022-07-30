package main

import (
	"log"

	pb "github.com/tomschdev/go-grpc-api/bi-directional-streaming/greet/proto"
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
	c := pb.NewGreetServiceClient(conn)
	doBiGreet(c)
}
