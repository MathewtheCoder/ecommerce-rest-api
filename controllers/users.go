package controllers

import (
	"errors"
	"net/http"

	"ecommercestorebackend/interfaces"
	"ecommercestorebackend/models"
	"ecommercestorebackend/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GET /user/list
// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /user/create
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input interfaces.CreateUserInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Name: input.Name, UserName: input.UserName, Password: string(password)}
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully"})
}

func UsersLogin(c *gin.Context) {

	var input interfaces.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.FindOneUser(&models.User{UserName: input.UserName})

	if err != nil {
		c.JSON(http.StatusForbidden, interfaces.CreateDetailedErrorDto("login_error", err))
		return
	}

	if user.IsValidPassword(input.Password) != nil {
		c.JSON(http.StatusForbidden, interfaces.CreateDetailedErrorDto("login", errors.New("invalid credentials")))
		return
	}
	// Generate JWT Token
	token := user.GenerateJwtToken()
	// user.Token = token
	models.DB.Model(&user).Update("Token", token)
	c.JSON(http.StatusOK, interfaces.CreateLoginSuccessful(&user, token))

}
