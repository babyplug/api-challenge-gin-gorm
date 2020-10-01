package services

import (
	"errors"
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/dto"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func FindAllPhoto(c *gin.Context) ([]models.Photo, error) {
	var photos []models.Photo
	database.DB.Find(&photos)
	return photos, nil
}

func CreatePhoto(c *gin.Context) (models.Photo, error) {
	var form dto.PhotoRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return models.Photo{}, err
	}

	// Create book
	photo := models.Photo{
		Name:        form.Name,
		Description: form.Description,
		FileName:    form.FileName,
		IsPublished: form.IsPublished,
		Views:       form.Views,
		AuthorId:    form.AuthorId,
	}
	database.DB.Create(&photo)
	return photo, nil
}

func getPhotoById(id string) (models.Photo, error) {
	var photo models.Photo
	if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		return photo, errors.New("Photo not found!")
	}
	return photo, nil
}

func FindPhotoById(c *gin.Context) (models.Photo, error) {
	photo, err := getPhotoById(c.Param("id"))
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func UpdatePhotoById(c *gin.Context) (models.Photo, error, int) {
	photo, err := getPhotoById(c.Param("id"))
	if err != nil {
		return photo, err, http.StatusNotFound
	}

	var form dto.PhotoRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return photo, err, http.StatusBadRequest
	}

	photo.Name = form.Name
	photo.Description = form.Description
	photo.FileName = form.FileName
	photo.IsPublished = form.IsPublished
	photo.Views = form.Views
	photo.AuthorId = form.AuthorId

	database.DB.Save(&photo)
	return photo, nil, http.StatusOK
}

func DeletePhotoById(c *gin.Context) error {
	photo, err := getPhotoById(c.Param("id"))
	if err != nil {
		return err
	}
	database.DB.Delete(&photo)
	return nil
}
