package main

import (
	"net/http"

	"crud/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Books)
	})

	r.POST("/books", func(c *gin.Context) {
		var book models.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		models.Books = append(models.Books, book)
		c.JSON(http.StatusCreated, book)
	})

	r.Run()
}
