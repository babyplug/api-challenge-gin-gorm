package controllers

import (
	"github.com/babyplug/api-challenge-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context)       { services.FindAllUser(c) }
func FindUserById(c *gin.Context)   { services.FindUserById(c) }
func CreateUser(c *gin.Context)     { services.CreateUser(c) }
func UpdateUserById(c *gin.Context) { services.UpdateUserById(c) }
func DeleteUserById(c *gin.Context) { services.DeleteUserById(c) }
