package interfaces

import "ecommercestorebackend/models"

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

func CreateLoginSuccessful(user *models.User, token string) map[string]interface{} {

	return map[string]interface{}{
		"success": true,
		"token":   token,
		"user": map[string]interface{}{
			"username": user.UserName,
			"id":       user.ID,
		},
	}
}
