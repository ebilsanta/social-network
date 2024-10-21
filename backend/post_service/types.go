package main

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
)

type CreatePostRequest struct {
	Caption  string `json:"caption"`
	ImageURL string `json:"imageURL"`
	PosterID string `json:"posterID"`
}

func NewPost(caption, imageURL string, userID int64) *pb.Post {
	return &pb.Post{
		Caption:   caption,
		ImageURL:  imageURL,
		UserID:    userID,
		CreatedAt: timestamppb.New(time.Now()),
	}
}
