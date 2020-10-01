package controllers

import (
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/dto"

	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials dto.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(&credentials)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"prefix": "Bearer",
			"token":  token,
		},
	)
}
