package services

import (
	"errors"
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

type PhotoMetadataRequestform struct {
	Height      int64  `json:"height" binding:"required"`
	Width       int64  `json:"width" binding:"required"`
	Orientation string `json:"orientation" binding:"required"`
	Compressed  int64  `json:"compressed" binding:"required"`
	Comment     string `json:"comment" binding:"required"`
	PhotoId     uint   `json:"photoId" binding:"required"`
}

func FindAllPhotoMetadata(c *gin.Context) ([]models.PhotoMetadata, error) {
	var photoMetadata []models.PhotoMetadata
	database.DB.Find(&photoMetadata)
	return photoMetadata, nil
}

func CreatePhotoMetadata(c *gin.Context) (models.PhotoMetadata, error) {
	var form PhotoMetadataRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return models.PhotoMetadata{}, nil
	}

	photoMetadata := models.PhotoMetadata{
		Height:      form.Height,
		Width:       form.Width,
		Orientation: form.Orientation,
		Compressed:  form.Compressed,
		Comment:     form.Comment,
		PhotoId:     form.PhotoId,
	}
	database.DB.Create(&photoMetadata)
	return photoMetadata, nil
}

func getPhotoMetadataById(id string) (models.PhotoMetadata, error) {
	var photoMetadata models.PhotoMetadata
	if err := database.DB.Where("id = ?", id).First(&photoMetadata).Error; err != nil {
		return photoMetadata, errors.New("PhotoMetadata not found!")
	}
	return photoMetadata, nil
}

func FindPhotoMetadataById(c *gin.Context) (models.PhotoMetadata, error) {
	photoMetadata, err := getPhotoMetadataById(c.Param("id"))
	if err != nil {
		return photoMetadata, err
	}
	return photoMetadata, nil
}

func UpdatePhotoMetadataById(c *gin.Context) (models.PhotoMetadata, error, int) {
	photoMetadata, err := getPhotoMetadataById(c.Param("id"))
	if err != nil {
		return photoMetadata, err, http.StatusNotFound
	}

	var form PhotoMetadataRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return photoMetadata, err, http.StatusBadRequest
	}
	photoMetadata.Height = form.Height
	photoMetadata.Width = form.Width
	photoMetadata.Orientation = form.Orientation
	photoMetadata.Compressed = form.Compressed
	photoMetadata.Comment = form.Comment
	photoMetadata.PhotoId = form.PhotoId
	database.DB.Save(&photoMetadata)
	return photoMetadata, nil, http.StatusOK
}

func DeletePhotoMetadataById(c *gin.Context) error {
	photoMetadata, err := getPhotoMetadataById(c.Param("id"))
	if err != nil {
		return err
	}
	database.DB.Delete(photoMetadata)
	return nil
}
