package middleware

import (
	"GinGonicGorm/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ini Lewat Headers
		token := c.GetHeader("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		res, err := jwt.ParseWithClaims(token, &utils.Claims{}, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !res.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": err.Error()})
			c.Abort()
			return
		}

		c.Next()

	}
}
