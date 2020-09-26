package main

import (
	_ "fmt"

	"github.com/babyplug/api-challenge-gin-gorm/controllers"
	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/user", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.FindUserById)
	r.PUT("/user/:id", controllers.UpdateUserById)
	r.DELETE("/user/:id", controllers.DeleteUserById)

	r.Run()
}
