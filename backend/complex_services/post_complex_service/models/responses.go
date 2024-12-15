package models

import (
	pb "github.com/ebilsanta/social-network/backend/complex_services/post_service/services/proto/generated"
)

type CreatePostResponse struct {
	Data *pb.Post `json:"data"`
}

type GetPostResponse struct {
	Data *pb.Post `json:"data"`
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
