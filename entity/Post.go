package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostedAt *time.Time
}
