package controllers

import (
	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindPhoto(c *gin.Context)       { services.FindAllPhoto(c) }
func FindPhotoById(c *gin.Context)   { services.FindPhotoById(c) }
func CreatePhoto(c *gin.Context)     { services.CreatePhoto(c) }
func UpdatePhotoById(c *gin.Context) { services.UpdatePhotoById(c) }
func DeletePhotoById(c *gin.Context) { services.DeletePhotoById(c) }
