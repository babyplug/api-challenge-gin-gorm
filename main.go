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

	// Photo
	r.GET("/photo", controllers.FindPhoto)
	r.POST("/photo", controllers.CreatePhoto)
	r.GET("/photo/:id", controllers.FindPhotoById)
	r.PUT("/photo/:id", controllers.UpdatePhotoById)
	r.DELETE("/photo/:id", controllers.DeletePhotoById)

	// PhotoMetadata
	r.GET("/photo-metadata", controllers.FindPhotoMetadata)
	r.POST("/photo-metadata", controllers.CreatePhotoMetadata)
	r.GET("/photo-metadata/:id", controllers.FindPhotoMetadataById)
	r.PUT("/photo-metadata/:id", controllers.UpdatePhotoMetadataById)
	r.DELETE("/photo-metadata/:id", controllers.DeletePhotoMetadataById)

	r.Run()
}
