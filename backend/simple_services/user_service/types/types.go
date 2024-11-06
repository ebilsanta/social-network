package types

import (
	"time"
)

type User struct {
	Id             string     `bson:"id"`
	Email          string     `bson:"email"`
	Name           string     `bson:"name"`
	Username       string     `bson:"username"`
	Image          string     `bson:"image"`
	PostCount      uint32     `bson:"postCount"`
	FollowerCount  uint32     `bson:"followerCount"`
	FollowingCount uint32     `bson:"followingCount"`
	CreatedAt      time.Time  `bson:"createdAt"`
	DeletedAt      *time.Time `bson:"deletedAt"`
}

func NewUser(id, email, name, username, image string) *User {
	return &User{
		Id:             id,
		Email:          email,
		Name:           name,
		Username:       username,
		Image:          image,
		PostCount:      0,
		FollowerCount:  0,
		FollowingCount: 0,
		CreatedAt:      time.Now(),
		DeletedAt:      nil,
	}
}
