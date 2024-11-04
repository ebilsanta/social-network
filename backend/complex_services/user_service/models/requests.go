package models

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	ImageURL string `json:"imageURL" binding:"required,uri"`
	Username string `json:"username" binding:"required"`
}
