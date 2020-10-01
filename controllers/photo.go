package controllers

import (
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindPhoto(c *gin.Context) {
	photos, err := services.FindAllPhoto(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photos,
		},
	)
}

func CreatePhoto(c *gin.Context) {
	photo, err := services.CreatePhoto(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": photo,
		},
	)
}

func FindPhotoById(c *gin.Context) {
	photo, err := services.FindPhotoById(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	photo, err, httpStatusCode := services.UpdatePhotoById(c)

	if err != nil {
		c.JSON(httpStatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photo,
		},
	)
}

func DeletePhotoById(c *gin.Context) {
	err := services.DeletePhotoById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{"status": "success"},
	)
}
