package services

import (
	"errors"
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

type PhotoRequestform struct {
	Description string `json:"description" binding:"required"`
	FileName    string `json:"fileName" binding:"required"`
	IsPublished bool   `json:"isPublished" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Views       int64  `json:"views" binding:"required"`
	AuthorId    uint   `json:"authorId" binding:"required"`
}

func FindAllPhoto(c *gin.Context) {
	var photos []models.Photo
	database.DB.Find(&photos)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photos,
		},
	)
}

func CreatePhoto(c *gin.Context) {
	var form PhotoRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": photo,
		},
	)
}

func getPhotoById(id string) (models.Photo, error) {
	var photo models.Photo
	if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		return photo, errors.New("Photo not found!")
	}
	return photo, nil
}

func FindPhotoById(c *gin.Context) {
	photo, err := getPhotoById(c.Param("id"))
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
			"data": photo,
		},
	)
}

func UpdatePhotoById(c *gin.Context) {
	photo, err := getPhotoById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var form PhotoRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	photo.Name = form.Name
	database.DB.Save(&photo)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photo,
		},
	)
}

func DeletePhotoById(c *gin.Context) {
	photo, err := getPhotoById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	database.DB.Delete(photo)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
		},
	)

}
