package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tomschdev/go-grpc-api/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/mgo.v2/bson"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("DeleteBlog was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot parse ID: %v\n", err),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot delete object in MongoDB: %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("blog was not found: %v\n", err),
		)
	}
	return &emptypb.Empty{}, nil
}
