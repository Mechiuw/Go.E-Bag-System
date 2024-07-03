package main

import (
	"net/http"

	"crud/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/bag", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Inventory)
	})

	r.POST("/bag", func(c *gin.Context) {
		var bag models.Bag

		if err := c.ShouldBindJSON(&bag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		models.Inventory = append(models.Inventory, bag)
		c.JSON(http.StatusCreated, bag)
	})

	r.PUT("/bag/:id", func(c *gin.Context) {
		id := c.Param("id")

		var bag models.Bag
		if err := c.ShouldBindJSON(&bag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		for idx, element := range models.Inventory {
			if element.Id == id {
				element.Name = bag.Name
				element.Quantity = bag.Quantity

				models.Inventory[idx] = element

				c.JSON(http.StatusOK, element)
				return
			}
		}

		// If no matching id found in Inventory
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Bag not found",
		})
	})

	r.Run()
}
