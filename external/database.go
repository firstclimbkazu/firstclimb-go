package external

import (
	"firstclimb-go/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Article{},
		&entity.TextItem{},
		&entity.MovieItem{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
