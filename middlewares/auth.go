package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hepiska/todo-go/models/service"
	"github.com/hepiska/todo-go/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authentication midleware and pass user
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authheader := c.Request.Header.Get("Authorization")
		if len(authheader) == 0 {
			c.JSON(400, gin.H{"error": "no token provided"})
			return
		}
		tokenstr := strings.TrimSpace(authheader)
		fmt.Println(tokenstr)
		token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
			secretKey := utils.EnvVar("TOKEN_KEY")
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			userservice := service.Userservice{}
			user, err := userservice.FindbyEmail(email)
			if err != nil {
				c.JSON(400, gin.H{"error": "user unauthorized"})
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.JSON(400, gin.H{"error": "token invalid"})
		}
	}
}

//ErrorHandler is for global error
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}
