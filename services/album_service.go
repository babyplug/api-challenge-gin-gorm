package services

import (
	"errors"
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

type AlbumRequestform struct {
	Name   string `json:"name" binding:"required"`
	Photos []uint `json:"photoId,omitempty"`
}

func FindAllAlbum(c *gin.Context) ([]models.Album, error) {
	var albums []models.Album
	database.DB.Find(&albums)
	return albums, nil
}

func CreateAlbum(c *gin.Context) (models.Album, error) {
	var form AlbumRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return models.Album{}, err
	}

	// Create book
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

func FindAlbumById(c *gin.Context) (models.Album, error) {
	album, err := getAlbumById(c.Param("id"))
	if err != nil {
		return models.Album{}, err
	}

	return album, nil
}

func UpdateAlbumById(c *gin.Context) (models.Album, error, int) {
	album, err := getAlbumById(c.Param("id"))
	if err != nil {
		return models.Album{}, err, http.StatusNotFound
	}

	var form AlbumRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return models.Album{}, err, http.StatusBadRequest
	}
	album.Name = form.Name
	database.DB.Save(&album)

	return album, nil, http.StatusOK
}

func DeleteAlbumById(c *gin.Context) error {
	album, err := getAlbumById(c.Param("id"))
	if err != nil {
		return err
	}

	database.DB.Delete(album)
	return nil
}
