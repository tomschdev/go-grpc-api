package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")
	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("blog read error: %v\n", err)
	}

	log.Printf("read blog: %v\n", res)
	return res
}
