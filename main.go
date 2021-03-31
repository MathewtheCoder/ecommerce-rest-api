package main

import (
	"net/http"

	"ecommercestorebackend/controllers"
	"ecommercestorebackend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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

	// Start Server
	r.Run()
}
