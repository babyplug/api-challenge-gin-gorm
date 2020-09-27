package controllers

import (
	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindPhotoMetadata(c *gin.Context)       { services.FindAllPhotoMetadata(c) }
func FindPhotoMetadataById(c *gin.Context)   { services.FindPhotoMetadataById(c) }
func CreatePhotoMetadata(c *gin.Context)     { services.CreatePhotoMetadata(c) }
func UpdatePhotoMetadataById(c *gin.Context) { services.UpdatePhotoMetadataById(c) }
func DeletePhotoMetadataById(c *gin.Context) { services.DeletePhotoMetadataById(c) }
