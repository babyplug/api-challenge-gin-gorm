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
	Photos []uint `json:"Photos,omitempty"`
}

func FindAllAuthor(c *gin.Context) {
	var authors []models.Author
	database.DB.Find(&authors)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": authors,
		},
	)
}

func CreateAuthor(c *gin.Context) {
	var form AuthorRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	author := models.Author{
		Name: form.Name,
	}
	database.DB.Create(&author)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": author,
		},
	)
}

func getAuthorById(id string) (models.Author, error) {
	var author models.Author
	if err := database.DB.Where("id = ?", id).First(&author).Error; err != nil {
		return author, errors.New("Author not found!")
	}
	return author, nil
}

func FindAuthorById(c *gin.Context) {
	author, err := getAuthorById(c.Param("id"))
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
			"data": author,
		},
	)
}

func UpdateAuthorById(c *gin.Context) {
	author, err := getAuthorById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var form AuthorRequestform
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author.Name = form.Name
	database.DB.Save(&author)

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": author,
		},
	)
}

func DeleteAuthorById(c *gin.Context) {
	author, err := getAuthorById(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	database.DB.Delete(author)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
		},
	)

}
