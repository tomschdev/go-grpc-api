package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("error while deleting blog: %v\n", err)
	}
	log.Println("blog was deleted")
}
