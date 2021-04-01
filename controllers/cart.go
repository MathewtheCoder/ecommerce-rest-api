package controllers

import (
	"net/http"

	"ecommercestorebackend/models"
	"ecommercestorebackend/services"

	"github.com/gin-gonic/gin"
)

type AddCartReq struct {
	ItemId string `json:"item_id" binding:"required"`
}

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
	// Validate input
	var input AddCartReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := c.Get("currentUserId")
	// Check if user has active cart
	var cart models.Cart
	if err := models.DB.Where("user_id = ? AND is_purchased = ?", userId, false).First(&cart).Error; err != nil {
		// If user doesnt have active cart create one
		cartentry := models.Cart{UserId: userId.(uint), IsPurchased: false}
		models.DB.Create(&cartentry)
		// Add relations in cart item table
		cartitementry := models.CartItem{UserID: userId.(uint), CartID: cartentry.ID}
		models.DB.Create(&cartitementry)
		// Update cart id to corresponding user
		var user models.User
		models.DB.First(&user, userId)
		services.UpdateUser(user, &models.User{CartId: cartentry.ID})
	}
	// Add relations in cart item table
	cartitementry := models.CartItem{UserID: userId.(uint), CartID: cart.ID}
	models.DB.Create(&cartitementry)
	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart successfully"})
}

// POST /cart/complete/:cartId
// Add item to cart
func BuyCart(c *gin.Context) {
	userId, _ := c.Get("currentUserId")
	var cart models.Cart
	if err := models.DB.Where("id = ? AND is_purchased = ?", c.Param("cartId"), false).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found!"})
		return
	}
	// Update IsPurchased to true
	cart.IsPurchased = true
	models.DB.Save(&cart)
	// Insert into orders
	order := models.Order{CartID: cart.ID, UserID: userId.(uint)}
	models.DB.Create(&order)
	c.JSON(http.StatusCreated, gin.H{"message": "Cart items have been processed"})
}
