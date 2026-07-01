package middleware

import (
	"net/http"
	"bookmarks/internal/jwt"
	"github.com/gin-gonic/gin"
	"strings"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization"), " ")[1]

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		log.Println("Token:", token)
		uid, err := jwt.VerifyToken(token)
		log.Println("UID:", uid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("userID", int64(uid))
		c.Next()
	}
}