package database

import (
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := viper.GetString("mysql.dsn")
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
