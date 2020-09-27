package controllers

import (
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindPhotoMetadata(c *gin.Context) {
	photoMetadata, err := services.FindAllPhotoMetadata(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photoMetadata,
		},
	)
}

func CreatePhotoMetadata(c *gin.Context) {
	photoMetadata, err := services.CreatePhotoMetadata(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": photoMetadata,
		},
	)
}

func FindPhotoMetadataById(c *gin.Context) {
	photoMetadata, err := services.FindPhotoMetadataById(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	photoMetadata, err, httpStatusCode := services.UpdatePhotoMetadataById(c)

	if err != nil {
		c.JSON(httpStatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": photoMetadata,
		},
	)
}

func DeletePhotoMetadataById(c *gin.Context) {
	err := services.DeletePhotoMetadataById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{"status": "success"},
	)
}
