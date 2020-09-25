package main

import (
	"github.com/gin-gonic/gin"
	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	_ "fmt"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		var users []models.User
		database.DB.Find(&users)
		c.JSON(200, gin.H{
			"message": "pong",
			"data": users,
		})
	})

	r.Run()
}
