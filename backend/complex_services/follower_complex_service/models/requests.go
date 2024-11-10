package models

type GetRequest struct {
	Page  int64 `form:"page,default=1" binding:"min=1"`
	Limit int64 `form:"limit,default=10" binding:"min=1,max=100"`
}

type ModifyRequest struct {
	FollowerId string `json:"followerId" binding:"required"`
	FollowingId string `json:"followingId" binding:"required"`
}
