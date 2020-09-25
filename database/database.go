package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/babyplug/api-challenge-gin-gorm/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DisableForeignKeyConstraintWhenMigrating: true,

	if err != nil {
		panic("Failed to connect to database!")
	}

	var user models.User
	var author models.Author
	var photo models.Photo
	var photoMetadata models.PhotoMetadata

	db.AutoMigrate(
		&user,
		&author,
		&photo,
		&photoMetadata,
	)

	DB = db
}