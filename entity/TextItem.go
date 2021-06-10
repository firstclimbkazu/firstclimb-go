package entity

type TextItem struct {
	ID        uint
	Order     uint
	Value     string
	CreatedAt string
	UpdatedAt string
	ArticleID uint
	Article   *Article
}
