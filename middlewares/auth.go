package middlewares

import (
	"ecommercestorebackend/models"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		uid := user.(models.User).ID
		if exists && uid != 0 {
			return
		} else {
			_ = c.AbortWithError(http.StatusUnauthorized, errors.New("Not Authorized"))
			return
		}
	}
}

func UserLoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer != "" {
			jwtParts := strings.Split(bearer, " ")
			if len(jwtParts) == 2 {
				jwtEncoded := jwtParts[1]

				token, err := jwt.Parse(jwtEncoded, func(token *jwt.Token) (interface{}, error) {
					// Theorically we have also to validate the algorithm
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
					}
					secret := []byte(os.Getenv("JWT_SECRET"))
					return secret, nil
				})

				if err != nil {
					println(err.Error())
					return
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					userId := uint(claims["user_id"].(float64))
					fmt.Printf("[+] Authenticated request, authenticated user id is %d\n", userId)

					var user models.User
					if userId != 0 {
						database := models.GetDb()
						// We always need the Roles to be loaded to make authorization decisions based on Roles
						database.First(&user, userId)
					}

					c.Set("currentUser", user)
					c.Set("currentUserId", user.ID)
				} else {

				}

			}
		}
	}
}
