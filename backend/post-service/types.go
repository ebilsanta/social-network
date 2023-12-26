package main

import (
	"math/rand"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Caption   string    `json:"caption"`
	ImageURL  string    `json:"imageURL"`
	PosterID  string    `json:"posterID"`
	Timestamp time.Time `json:"timestamp"`
}

func NewPost(caption, imageURL, posterID string) *Post {
	return &Post{
		ID:        rand.Intn(10000), // to change
		Caption:   caption,
		ImageURL:  imageURL,
		PosterID:  posterID,
		Timestamp: time.Now(),
	}
}
