package controllers

import (
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/dto"
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

func bindJsonToForm(c *gin.Context, form *dto.AlbumRequestForm) error {
	if err := c.ShouldBindJSON(form); err != nil {
		return err
	}
	return nil
}

func CreateAlbum(c *gin.Context) {
	var form dto.AlbumRequestForm
	err := bindJsonToForm(c, &form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, _ := services.CreateAlbum(&form)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": album,
		},
	)
}

func FindAlbumById(c *gin.Context) {
	album, err := services.FindAlbumById(c.Param("id"))
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
	var form dto.AlbumRequestForm
	err := bindJsonToForm(c, &form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := services.UpdateAlbumById(c.Param("id"), &form)
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

func DeleteAlbumById(c *gin.Context) {
	err := services.DeleteAlbumById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{"status": "success"},
	)
}
