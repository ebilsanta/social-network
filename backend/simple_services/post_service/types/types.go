package types

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto/generated"
)

func NewPost(caption, image, userID string) *pb.Post {
	return &pb.Post{
		Caption:   caption,
		Image:     image,
		UserId:    userID,
		CreatedAt: timestamppb.New(time.Now()),
	}
}
