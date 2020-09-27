package services

import (
	"errors"
	"net/http"

	"github.com/babyplug/api-challenge-gin-gorm/database"
	"github.com/babyplug/api-challenge-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

type AuthorRequestform struct {
	Name   string `json:"name" binding:"required"`
	Photos []uint `json:"photoId,omitempty"`
}

func FindAllAuthor(c *gin.Context) ([]models.Author, error) {
	var authors []models.Author
	database.DB.Find(&authors)
	return authors, nil
}

func CreateAuthor(c *gin.Context) (models.Author, error) {
	var form AuthorRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return models.Author{}, err
	}

	author := models.Author{
		Name: form.Name,
	}
	database.DB.Create(&author)
	return author, nil
}

func getAuthorById(id string) (models.Author, error) {
	var author models.Author
	if err := database.DB.Where("id = ?", id).First(&author).Error; err != nil {
		return author, errors.New("Author not found!")
	}
	return author, nil
}

func FindAuthorById(c *gin.Context) (models.Author, error) {
	author, err := getAuthorById(c.Param("id"))
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}

func UpdateAuthorById(c *gin.Context) (models.Author, error, int) {
	author, err := getAuthorById(c.Param("id"))
	if err != nil {
		return author, err, http.StatusNotFound
	}

	var form AuthorRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		return author, err, http.StatusBadRequest
	}
	author.Name = form.Name
	database.DB.Save(&author)
	return author, nil, http.StatusOK
}

func DeleteAuthorById(c *gin.Context) error {
	author, err := getAuthorById(c.Param("id"))
	if err != nil {
		return err
	}

	database.DB.Delete(author)
	return nil
}
