package entity

type Article struct {
	ID          uint
	Title       string
	Detail      string
	CreatedAt   string
	UpdatedAt   string
	Post        *Post
	PostID      string
}
