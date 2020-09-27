package controllers

import (
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindAlbum(c *gin.Context) {
	albums, err := services.FindAllAlbum(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": albums,
		},
	)
}

func CreateAlbum(c *gin.Context) {
	album, err := services.CreateAlbum(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": album,
		},
	)
}

func FindAlbumById(c *gin.Context) {
	album, err := services.FindAlbumById(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	album, err, httpStatusCode := services.UpdateAlbumById(c)

	if err != nil {
		c.JSON(httpStatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": album,
		},
	)
}

func DeleteAlbumById(c *gin.Context) {
	err := services.DeleteAlbumById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{"status": "success"},
	)
}
