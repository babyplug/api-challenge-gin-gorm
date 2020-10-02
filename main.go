package main

import (
	"fmt"

	"github.com/babyplug/api-challenge-gin-gorm/controllers"
	"github.com/babyplug/api-challenge-gin-gorm/database"
	_ "github.com/babyplug/api-challenge-gin-gorm/docs"
	"github.com/babyplug/api-challenge-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setupViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if viper.GetString("app.mode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	setupViper()
	database.ConnectDatabase()

	router := gin.Default()
	router.StaticFile("/swagger-ui.json", "./swagger.json")

	apiRouter := router.Group("/api")
	apiRouter.POST("/login", controllers.Login)
	{
		// create	auth routes
		authRouter := apiRouter.Group("/auth", middleware.TokenAuthMiddleware())

		authRouter.GET("/user", controllers.FindUser)
		authRouter.POST("/user", controllers.CreateUser)
		authRouter.GET("/user/:id", controllers.FindUserById)
		authRouter.PUT("/user/:id", controllers.UpdateUserById)
		authRouter.DELETE("/user/:id", controllers.DeleteUserById)

		// Author
		authRouter.GET("/author", controllers.FindAuthor)
		authRouter.POST("/author", controllers.CreateAuthor)
		authRouter.GET("/author/:id", controllers.FindAuthorById)
		authRouter.PUT("/author/:id", controllers.UpdateAuthorById)
		authRouter.DELETE("/author/:id", controllers.DeleteAuthorById)

		// Photo
		authRouter.GET("/photo", controllers.FindPhoto)
		authRouter.POST("/photo", controllers.CreatePhoto)
		authRouter.GET("/photo/:id", controllers.FindPhotoById)
		authRouter.PUT("/photo/:id", controllers.UpdatePhotoById)
		authRouter.DELETE("/photo/:id", controllers.DeletePhotoById)

		// PhotoMetadata
		authRouter.GET("/photo-metadata", controllers.FindPhotoMetadata)
		authRouter.POST("/photo-metadata", controllers.CreatePhotoMetadata)
		authRouter.GET("/photo-metadata/:id", controllers.FindPhotoMetadataById)
		authRouter.PUT("/photo-metadata/:id", controllers.UpdatePhotoMetadataById)
		authRouter.DELETE("/photo-metadata/:id", controllers.DeletePhotoMetadataById)

		// PhotoMetadata
		authRouter.GET("/album", controllers.FindAlbum)
		authRouter.POST("/album", controllers.CreateAlbum)
		authRouter.GET("/album/:id", controllers.FindAlbumById)
		authRouter.PUT("/album/:id", controllers.UpdateAlbumById)
		authRouter.DELETE("/album/:id", controllers.DeleteAlbumById)
	}

	router.Run()
}
