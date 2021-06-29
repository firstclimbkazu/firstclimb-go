package model

import (
	"firstclimb-go/entity"
	"fmt"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func UserFromEntity(e *entity.User) *User {
	return &User{
		ID:        fmt.Sprintf("%d", e.ID),
		Name:      e.Name,
		Email:     e.Email,
		CreatedAt: e.CreatedAt.String(),
	}
}
