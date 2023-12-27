package main

import (
	"database/sql"
	"encoding/json"
	"time"
)

type CreatePostRequest struct {
	Caption  string `json:"caption"`
	ImageURL string `json:"imageURL"`
	PosterID string `json:"posterID"`
}

type Post struct {
	ID        int          `json:"id"`
	Caption   string       `json:"caption"`
	ImageURL  string       `json:"imageURL"`
	UserID    string       `json:"userID"`
	CreatedAt time.Time    `json:"createdAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}

func NewPost(caption, imageURL, userID string) *Post {
	return &Post{
		Caption:   caption,
		ImageURL:  imageURL,
		UserID:  userID,
		CreatedAt: time.Now().UTC(),
	}
}

func (p *Post) MarshalJSON() ([]byte, error) {
	type Alias Post
	if p.DeletedAt.Valid {
		return json.Marshal(&struct {
			*Alias
			DeletedAt string `json:"deletedAt"`
		}{
			Alias:     (*Alias)(p),
			DeletedAt: p.DeletedAt.Time.Format(time.RFC3339),
		})
	} else {
		// Omit deletedAt field if it's null
		return json.Marshal(&struct {
			*Alias
			DeletedAt interface{} `json:"deletedAt"`
		}{
			Alias:     (*Alias)(p),
			DeletedAt: nil,
		})
	}
}
