package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Username  string             `bson:"username"`
	ImageURL  string             `bson:"imageURL"`
	CreatedAt time.Time          `bson:"createdAt"`
	DeletedAt *time.Time         `bson:"deletedAt"`
}

func NewUser(email, username, imageURL string) *User {
	return &User{
		Id:        primitive.NewObjectID(),
		Email:     email,
		Username:  username,
		ImageURL:  imageURL,
		CreatedAt: time.Now(),
		DeletedAt: nil,
	}
}
