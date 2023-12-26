package main

import (
	"time"
)

type CreatePostRequest struct {
	Caption  string `json:caption`
	ImageURL string `json:imageURL`
	PosterID string `json:posterID`
}

type Post struct {
	ID        int       `json:"id"`
	Caption   string    `json:"caption"`
	ImageURL  string    `json:"imageURL"`
	PosterID  string    `json:"posterID"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPost(caption, imageURL, posterID string) *Post {
	return &Post{
		Caption:   caption,
		ImageURL:  imageURL,
		PosterID:  posterID,
		CreatedAt: time.Now().UTC(),
	}
}
