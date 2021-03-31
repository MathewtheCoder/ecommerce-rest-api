package main

import (
	"net/http"

	"ecommercestorebackend/controllers"
	"ecommercestorebackend/middlewares"
	"ecommercestorebackend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Allow all origins
	r.Use(cors.Default())
	r.Use(middlewares.UserLoaderMiddleware())
	// DB connection "sqlite"
	database := models.ConnectDataBase()
	defer database.Close()
	// Api routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	// user registration
	r.POST("/user/create", controllers.CreateUser)
	// Get list of users
	r.GET("/user/list", controllers.GetUsers)
	r.POST("/user/login", controllers.UsersLogin)

	// Get list of items
	r.GET("/item/list", controllers.GetItems)
	// Create Item
	r.POST("/item/create", controllers.CreateItem)
	// Get list of carts
	r.GET("/cart/list", controllers.GetCarts)
	// Add item to cart
	r.POST(
		"/cart/add",
		middlewares.EnsureLoggedIn(),
		controllers.AddToCart,
	)
	// changed the order api route due to below issue
	// https://github.com/gin-gonic/gin/issues/1681
	r.POST(
		"/cart/complete/:cartId",
		middlewares.EnsureLoggedIn(),
		controllers.BuyCart,
	)
	// Get list of carts
	r.GET("/order/list", controllers.GetOrders)
	// Start Server
	r.Run()
}
