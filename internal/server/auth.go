package server

import (
	"bookmarks/internal/handler/auth"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerAuthRoutes(api *gin.RouterGroup) {
	api.POST("/auth/login", func(c *gin.Context) {
		auth.LoginHandler(c, s.db)
	})
	api.POST("/auth/register", func(c *gin.Context) {
		auth.RegisterHandler(c, s.db)
	})
}