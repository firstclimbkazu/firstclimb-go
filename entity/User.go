package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);"`
	Password string `gorm:"type:varchar(100);"`
	Email    string `gorm:"type:varchar(100);"`
}
