package middleware

import (
	"net/http"
	"strings"

	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/facades"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak disediakan"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"]

		var user models.User
		if err := facades.DB.Where("id = ? AND jwt_token = ?", userId, tokenString).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid di database"})
			c.Abort()
			return
		}

		c.Next()
	}
}
