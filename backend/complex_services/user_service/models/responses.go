package models

type CreateUserResponse struct {
	Data User `json:"data"`
}

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	PostCount      int    `json:"postCount"`
	FollowerCount  int    `json:"followerCount"`
	FollowingCount int    `json:"followingCount"`
	ImageURL       string `json:"imageURL"`
	Email          string `json:"email"`
	CreatedAt      string `json:"createdAt"`
}
