package models

import (
	pb "github.com/ebilsanta/social-network/backend/complex_services/post_service/services/proto/generated"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreatePostResponse struct {
	Data *pb.Post `json:"data"`
}

type GetPostResponse struct {
	Data *Post `json:"data"`
}

type Post struct {
	Id        int64                  `json:"id,omitempty"`
	Caption   string                 `json:"caption,omitempty"`
	Image     string                 `json:"image,omitempty"`
	User      *pb.User               `json:"user,omitempty"`
	CreatedAt *timestamppb.Timestamp `json:"createdAt,omitempty"`
	DeletedAt *timestamppb.Timestamp `json:"deletedAt,omitempty"`
}

type GetPostsResponse struct {
	Data       []*pb.Post          `json:"data"`
	Pagination *PaginationMetadata `json:"pagination"`
}

type PaginationMetadata struct {
	TotalRecords int64  `json:"totalRecords"`
	CurrentPage  int64  `json:"currentPage"`
	TotalPages   int64  `json:"totalPages"`
	NextPage     *int64 `json:"nextPage"`
	PrevPage     *int64 `json:"prevPage"`
}
