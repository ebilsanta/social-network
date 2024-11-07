package models

type CreateUserRequest struct {
	Id       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Image    string `json:"image" binding:"required,uri"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type GetUsersRequest struct {
	Page  int64  `form:"page,default=1" binding:"min=1"`
	Limit int64  `form:"limit,default=10" binding:"min=1,max=100"`
	Query string `form:"query"`
}
