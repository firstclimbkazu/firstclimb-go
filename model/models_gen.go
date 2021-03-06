// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ArticleData struct {
	ArticleID *string `json:"articleId"`
	Title     string  `json:"title"`
	Detail    string  `json:"detail"`
	PostID    string  `json:"postId"`
}

type MovieItemData struct {
	MovieItemID *string `json:"movieItemId"`
	Order       string  `json:"order"`
	Value       string  `json:"value"`
	ArticleID   string  `json:"articleId"`
}

type PostData struct {
	PostID   *string `json:"postId"`
	PostedAt *string `json:"postedAt"`
}

type TextItemData struct {
	TextItemID *string `json:"textItemId"`
	Order      string  `json:"order"`
	Value      string  `json:"value"`
	ArticleID  string  `json:"articleId"`
}

type UserData struct {
	UserID   *string `json:"userId"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
}

type OrderBy string

const (
	OrderByCreatedAtAsc  OrderBy = "createdAt_ASC"
	OrderByCreatedAtDesc OrderBy = "createdAt_DESC"
)

var AllOrderBy = []OrderBy{
	OrderByCreatedAtAsc,
	OrderByCreatedAtDesc,
}

func (e OrderBy) IsValid() bool {
	switch e {
	case OrderByCreatedAtAsc, OrderByCreatedAtDesc:
		return true
	}
	return false
}

func (e OrderBy) String() string {
	return string(e)
}

func (e *OrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderBy", str)
	}
	return nil
}

func (e OrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
