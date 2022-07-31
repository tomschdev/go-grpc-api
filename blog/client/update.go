package main

import (
	"context"
	"log"

	pb "github.com/tomschdev/go-grpc-api/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("update blog was invoked on blog with id: %s\n", id)
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "stevedev",
		Title:    "stevedev's preferred title",
		Content:  "I'm stevedev and i'm updating this blog",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("error occured while updating: %v\n", err)
	}

	log.Println("blog was updated")
}
