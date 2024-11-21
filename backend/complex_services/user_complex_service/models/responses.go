package models

import "google.golang.org/protobuf/types/known/timestamppb"

type CreateUserResponse struct {
	Data *User `json:"data"`
}

type GetUsersResponse struct {
	Data       []*User             `json:"data"`
	Pagination *PaginationMetadata `json:"pagination"`
}

type GetUserResponse struct {
	Data *User `json:"data"`
}

type User struct {
	Id             string                 `json:"id"`
	Email          string                 `json:"email,omitempty"`
	Name           string                 `json:"name"`
	Username       string                 `json:"username"`
	PostCount      uint32                 `json:"postCount"`
	FollowerCount  uint32                 `json:"followerCount"`
	FollowingCount uint32                 `json:"followingCount"`
	Image          string                 `json:"image"`
	CreatedAt      *timestamppb.Timestamp `json:"createdAt"`
	DeletedAt      *timestamppb.Timestamp `json:"deletedAt,omitempty"`
}

type PaginationMetadata struct {
	TotalRecords int64  `json:"totalRecords"`
	CurrentPage  int64  `json:"currentPage"`
	TotalPages   int64  `json:"totalPages"`
	NextPage     *int64 `json:"nextPage"`
	PrevPage     *int64 `json:"prevPage"`
}

type CheckFollowingResponse struct {
	Following bool `json:"following"`
}
