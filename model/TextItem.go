package model

import (
	"firstclimb-go/entity"
	"fmt"
)

type TextItem struct {
	ID        string `json:"id"`
	Order     uint   `json:"order"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ArticleID string `json:"article"`
}

func NewTextItemFromEntity(e *entity.TextItem) *TextItem {
	return &TextItem{
		ID:        fmt.Sprintf("%d", e.ID),
		Order:     e.Order,
		Value:     e.Value,
		CreatedAt: e.CreatedAt,
		ArticleID: fmt.Sprintf("%d", e.ArticleID),
	}
}
