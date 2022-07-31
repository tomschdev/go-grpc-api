package main

import (
	"context"
	"log"
	"net"

	pb "github.com/tomschdev/go-grpc-api/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("failed to create client for mongodb: %v\n", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v\n", err)
	}
	collection = client.Database("blogdb").Collection("blog")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen on: %v\n", err)
	}
	log.Printf("listening on %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
