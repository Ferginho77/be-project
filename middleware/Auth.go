package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization token diperlukan",
			})
			c.Abort()
			return
		}

		// Validasi format Bearer token
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Format Authorization harus Bearer token",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		jwtSecret := os.Getenv("JWT_TOKEN")

		if jwtSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "JWT configuration error",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token tidak valid atau sudah expired",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}