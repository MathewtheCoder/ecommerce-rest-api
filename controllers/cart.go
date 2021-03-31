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
