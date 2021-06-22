package model

import (
	"firstclimb-go/entity"
	"fmt"
)

type Post struct {
	ID        string `json:"id"`
	PostedAt  string `json:"postedAt"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewPostFromEntity(e *entity.Post) *Post {
	return &Post{
		ID:        fmt.Sprintf("%d", e.ID),
		CreatedAt: e.CreatedAt.String(),
	}
}
