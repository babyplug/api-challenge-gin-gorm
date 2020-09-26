package controllers

import (
	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindAuthor(c *gin.Context)       { services.FindAllAuthor(c) }
func FindAuthorById(c *gin.Context)   { services.FindAuthorById(c) }
func CreateAuthor(c *gin.Context)     { services.CreateAuthor(c) }
func UpdateAuthorById(c *gin.Context) { services.UpdateAuthorById(c) }
func DeleteAuthorById(c *gin.Context) { services.DeleteAuthorById(c) }
