package types

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto/generated"
)

func NewPost(caption, imageURL, userID string) *pb.Post {
	return &pb.Post{
		Caption:   caption,
		ImageURL:  imageURL,
		UserId:    userID,
		CreatedAt: timestamppb.New(time.Now()),
	}
}
