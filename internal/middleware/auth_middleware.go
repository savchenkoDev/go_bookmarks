package middleware

import (
	"strings"

	apperr "bookmarks/internal/errors"
	"bookmarks/internal/handler"
	"bookmarks/internal/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		parts := strings.Split(c.GetHeader("Authorization"), " ")
		if len(parts) != 2 || parts[1] == "" {
			handler.RespondError(c, apperr.UnauthorizedError())
			c.Abort()
			return
		}

		uid, err := jwt.VerifyToken(parts[1])
		if err != nil {
			handler.RespondError(c, apperr.UnauthorizedError())
			c.Abort()
			return
		}
		c.Set("userID", int64(uid))
		c.Next()
	}
}
