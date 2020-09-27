package controllers

import (
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindAuthor(c *gin.Context) {
	authors, err := services.FindAllAuthor(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": authors,
		},
	)
}

func CreateAuthor(c *gin.Context) {
	author, err := services.CreateAuthor(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": author,
		},
	)
}

func FindAuthorById(c *gin.Context) {
	author, err := services.FindAuthorById(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": author,
		},
	)
}

func UpdateAuthorById(c *gin.Context) {
	author, err, httpStatusCode := services.UpdateAuthorById(c)

	if err != nil {
		c.JSON(httpStatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": author,
		},
	)
}

func DeleteAuthorById(c *gin.Context) {
	err := services.DeleteAuthorById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{"status": "success"},
	)
}
