package main

import (
	"log"

	pb "github.com/tomschdev/go-grpc-api/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	readBlog(c, id) //should be a valid read on recently created blog
	// readBlog(c, "FakeIDMadeonMSWordUsingBrothersDriversLicence")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
	listBlog(c)
}
