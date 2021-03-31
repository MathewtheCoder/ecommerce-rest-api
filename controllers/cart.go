package controllers

import (
	"net/http"

	"ecommercestorebackend/models"

	"github.com/gin-gonic/gin"
)

// GET /cart/list
// List all carts
func GetCarts(c *gin.Context) {
	var carts []models.Cart
	models.DB.Find(&carts)

	c.JSON(http.StatusOK, gin.H{"data": carts})
}

// POST /cart/add
// Add item to cart
func AddToCart(c *gin.Context) {
	// // Validate input
	// var input CreateItemReq
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Create book
	// item := models.Item{Name: input.Name}
	// models.DB.Create(&item)

	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart successfully"})
}

// POST /cart/complete/:cartId
// Add item to cart
func BuyCart(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"message": "Cart items have been processed"})
}
