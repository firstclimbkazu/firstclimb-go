package entity

import "gorm.io/gorm"

type MovieItem struct {
	gorm.Model
	Order     uint
	Value     string `gorm:"type:text;"`
	ArticleID uint
	Article   Article
}
