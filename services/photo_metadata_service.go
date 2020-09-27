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

func FindAllPhotoMetadata(c *gin.Context) {
	var photoMetadata []models.PhotoMetadata
	database.DB.Find(&photoMetadata)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photoMetadata,
		},
	)
}

func CreatePhotoMetadata(c *gin.Context) {
	var form PhotoMetadataRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	photoMetadata := models.PhotoMetadata{
		Height:      form.Height,
		Width:       form.Width,
		Orientation: form.Orientation,
		Compressed:  form.Compressed,
		Comment:     form.Comment,
		PhotoId:     form.PhotoId,
	}
	database.DB.Create(&photoMetadata)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": photoMetadata,
		},
	)
}

func getPhotoMetadataById(id string) (models.PhotoMetadata, error) {
	var photoMetadata models.PhotoMetadata
	if err := database.DB.Where("id = ?", id).First(&photoMetadata).Error; err != nil {
		return photoMetadata, errors.New("PhotoMetadata not found!")
	}
	return photoMetadata, nil
}

func FindPhotoMetadataById(c *gin.Context) {
	photoMetadata, err := getPhotoMetadataById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photoMetadata,
		},
	)
}

func UpdatePhotoMetadataById(c *gin.Context) {
	photoMetadata, err := getPhotoMetadataById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var form PhotoMetadataRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	photoMetadata.Height = form.Height
	photoMetadata.Width = form.Width
	photoMetadata.Orientation = form.Orientation
	photoMetadata.Compressed = form.Compressed
	photoMetadata.Comment = form.Comment
	photoMetadata.PhotoId = form.PhotoId
	database.DB.Save(&photoMetadata)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photoMetadata,
		},
	)
}

func DeletePhotoMetadataById(c *gin.Context) {
	photoMetadata, err := getPhotoMetadataById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	database.DB.Delete(photoMetadata)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
		},
	)

}
