package external

import (
	"firstclimb-go/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	dsn := DbUser + ":" + DbPass + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
