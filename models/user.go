package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token" gorm:"null"`
	CartId   uint   `json:"cart_id" gorm:"default:null"`
}

func (u *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// Generate JWT token associated to this user
func (user *User) GenerateJwtToken() string {
	jwt_token := jwt.New(jwt.SigningMethodHS512)
	jwt_token.Claims = jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.UserName,
		"exp":      time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}
