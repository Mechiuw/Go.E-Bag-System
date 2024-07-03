package main

import (
	"net/http"

	"crud/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/bag", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Bag)
	})

	r.POST("/bag", func(c *gin.Context) {
		var items models.Item

		if err := c.ShouldBindJSON(&items); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		models.Bag = append(models.Bag, items)
		c.JSON(http.StatusCreated, items)
	})

	r.PUT("/bag/:id", func(c *gin.Context) {
		id := c.Param("id")

		var items models.Item

		//exception || error catch
		if err := c.ShouldBindJSON(&items); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// success
		for idx, element := range models.Bag {
			if element.Id == id {
				element.Name = items.Name
				element.Quantity = items.Quantity

				models.Bag[idx] = element

				c.JSON(http.StatusOK, element)
				return
			}
		}

		// If no matching id found in Inventory
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Bag not found",
		})
	})

	r.DELETE("/bag/:id", func(c *gin.Context) {
		id := c.Param("id")

		var removedItem models.Item

		for idx, element := range models.Bag {
			if element.Id == id {

				removedItem = models.Bag[idx]
				models.Bag = append(models.Bag[:idx], models.Bag[idx+1:]...)
				break
			}
		}

		if removedItem.Id == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"removed": removedItem,
			})
		}
	})

	r.Run()
}
