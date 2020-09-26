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

	// user
	r.GET("/user", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.FindUserById)
	r.PUT("/user/:id", controllers.UpdateUserById)
	r.DELETE("/user/:id", controllers.DeleteUserById)

	// Author
	r.GET("/author", controllers.FindAuthor)
	r.POST("/author", controllers.CreateAuthor)
	r.GET("/author/:id", controllers.FindAuthorById)
	r.PUT("/author/:id", controllers.UpdateAuthorById)
	r.DELETE("/author/:id", controllers.DeleteAuthorById)

	r.Run()
}
