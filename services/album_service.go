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

func FindAllAlbum(c *gin.Context) {
	var albums []models.Album
	database.DB.Find(&albums)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": albums,
		},
	)
}

func CreateAlbum(c *gin.Context) {
	var form AlbumRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	album := models.Album{
		Name: form.Name,
	}
	database.DB.Create(&album)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": album,
		},
	)
}

func getAlbumById(id string) (models.Album, error) {
	var album models.Album
	if err := database.DB.Where("id = ?", id).First(&album).Error; err != nil {
		return album, errors.New("Album not found!")
	}
	return album, nil
}

func FindAlbumById(c *gin.Context) {
	album, err := getAlbumById(c.Param("id"))
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
			"data": album,
		},
	)
}

func UpdateAlbumById(c *gin.Context) {
	album, err := getAlbumById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var form AlbumRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album.Name = form.Name
	database.DB.Save(&album)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": album,
		},
	)
}

func DeleteAlbumById(c *gin.Context) {
	album, err := getAlbumById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	database.DB.Delete(album)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
		},
	)

}
