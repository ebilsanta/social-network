package models

import "google.golang.org/protobuf/types/known/timestamppb"

type CreatePostRequest struct {
	Caption string `json:"caption" binding:"required"`
	Image   string `json:"image" binding:"required,uri"`
	UserId  string `json:"userId" binding:"required"`
}

type GetPostsByUserRequest struct {
	Page  int64 `form:"page,default=1" binding:"min=1"`
	Limit int64 `form:"limit,default=10" binding:"min=1,max=100"`
}

type Post struct {
	Id        int64                  `json:"id"`
	Caption   string                 `json:"caption"`
	Image     string                 `json:"image"`
	UserId    string                 `json:"userId"`
	CreatedAt *timestamppb.Timestamp `json:"createdAt"`
	DeletedAt *timestamppb.Timestamp `json:"deletedAt"`
}
