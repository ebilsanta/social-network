package models

type CreateUserRequest struct {
	Id       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Image    string `json:"image" binding:"required,uri"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
}
