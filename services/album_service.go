package services

import (
	"errors"

	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/dto"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func FindAllAlbum(c *gin.Context) ([]models.Album, error) {
	var albums []models.Album
	database.DB.Find(&albums)
	return albums, nil
}

func CreateAlbum(form *dto.AlbumRequestForm) (models.Album, error) {
	album := models.Album{
		Name: form.Name,
	}
	database.DB.Create(&album)
	return album, nil
}

func getAlbumById(id string) (models.Album, error) {
	var album models.Album
	if err := database.DB.Where("id = ?", id).First(&album).Error; err != nil {
		return album, errors.New("Album not found!")
	}
	return album, nil
}

func FindAlbumById(id string) (models.Album, error) {
	album, err := getAlbumById(id)
	if err != nil {
		return album, err
	}
	return album, nil
}

func UpdateAlbumById(id string, form *dto.AlbumRequestForm) (models.Album, error) {
	album, err := getAlbumById(id)
	if err != nil {
		return album, err
	}

	album.Name = form.Name
	database.DB.Save(&album)
	return album, nil
}

func DeleteAlbumById(id string) error {
	album, err := getAlbumById(id)
	if err != nil {
		return err
	}

	database.DB.Delete(album)
	return nil
}
