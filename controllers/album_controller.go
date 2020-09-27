package controllers

import (
	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindAlbum(c *gin.Context)       { services.FindAllAlbum(c) }
func FindAlbumById(c *gin.Context)   { services.FindAlbumById(c) }
func CreateAlbum(c *gin.Context)     { services.CreateAlbum(c) }
func UpdateAlbumById(c *gin.Context) { services.UpdateAlbumById(c) }
func DeleteAlbumById(c *gin.Context) { services.DeleteAlbumById(c) }
