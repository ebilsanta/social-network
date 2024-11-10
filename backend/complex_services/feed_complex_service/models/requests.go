package models

type GetFeedRequest struct {
	Page  int64 `form:"page,default=1" binding:"min=1"`
	Limit int64 `form:"limit,default=10" binding:"min=1,max=100"`
}
