package controllers

import (
	"net/http"

	"ecommercestorebackend/models"

	"github.com/gin-gonic/gin"
)

// GET /order/list
// List all orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	models.DB.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}
