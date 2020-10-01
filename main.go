package main

import (
	"github.com/babyplug/api-challenge-gin-gorm/controllers"
	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	// router := gin.Default()
	database.ConnectDatabase()

	// set all routes with prefix "/api"
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
