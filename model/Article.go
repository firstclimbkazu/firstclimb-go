package model

import (
	"firstclimb-go/entity"
	"fmt"
)

type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	PostID    string `json:"post"`
}

func NewArticleFromEntity(e *entity.Article) *Article {
	return &Article{
		ID:        fmt.Sprintf("%d", e.ID),
		Title:     e.Title,
		Detail:    e.Detail,
		CreatedAt: e.CreatedAt.String(),
		PostID:    fmt.Sprintf("%d", e.PostID),
	}
}
