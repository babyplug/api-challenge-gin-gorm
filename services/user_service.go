package services

import (
	"errors"
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRequestform struct {
	Name      string  `json:"name" binding:"required"`
	Email     *string `json:"email" binding:"required"`
	Age       uint8   `json:"age" binding:"required"`
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}

func FindAllUser(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": users,
		},
	)
}

func CreateUser(c *gin.Context) {
	var form UserRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)

	// Create book
	user := models.User{
		Name:      form.Name,
		Email:     form.Email,
		Age:       form.Age,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Username:  form.Username,
		Password:  string(hashPassword),
	}
	database.DB.Create(&user)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": user,
		},
	)
}

func getUserById(id string) (models.User, error) {
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, errors.New("User not found!")
	}
	return user, nil
}

func findUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, errors.New("User not found!")
	}
	return user, nil
}

func FindUserById(c *gin.Context) {
	user, err := getUserById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": user,
		},
	)
}

func UpdateUserById(c *gin.Context) {
	user, err := getUserById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var form UserRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(form)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": user,
		},
	)
}

func DeleteUserById(c *gin.Context) {
	user, err := getUserById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	database.DB.Delete(user)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}
