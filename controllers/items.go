package controllers

import (
	"net/http"

	"ecommercestorebackend/models"

	"github.com/gin-gonic/gin"
)

type CreateItemReq struct {
	Name string `json:"name" binding:"required"`
}

// GET /item/list
// List all items
func GetItems(c *gin.Context) {
	var items []models.Item
	models.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{"data": items})
}

// POST /item/create
// Create new item
func CreateItem(c *gin.Context) {
	// Validate input
	var input CreateItemReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	item := models.Item{Name: input.Name}
	models.DB.Create(&item)

	c.JSON(http.StatusCreated, gin.H{"message": "Item successfully created", "data": item})
}
