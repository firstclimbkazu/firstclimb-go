package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title  string `gorm:"type:varchar(255);"`
	Detail string `gorm:"type:varchar(255);"`
	Post   *Post
	PostID uint
}
