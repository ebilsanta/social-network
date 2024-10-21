package main

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
)

func NewPost(caption, imageURL string, userID int64) *pb.Post {
	return &pb.Post{
		Caption:   caption,
		ImageURL:  imageURL,
		UserID:    userID,
		CreatedAt: timestamppb.New(time.Now()),
	}
}
