package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Printf("createBlog invoked\n")
	blog := &pb.Blog{
		AuthorId: "tomschdev",
		Title:    "First Blog",
		Content:  "Content of the first blog.",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("failed to create blog: %v\n", err)
	}
	log.Printf("New blog with BlogId: %s\n", res.Id)
	return res.Id

}
