package server

import (
	"bookmarks/internal/handler"
	"bookmarks/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerUserRoutes(api *gin.RouterGroup) {
	api.POST("/auth", func(c *gin.Context) {
		handler.AuthHandler(c, s.db)
	})
	api.POST("/users", func(c *gin.Context) {
		user.CreateHandler(c, s.db)
	})
}

func (s *Server) registerUserProtectedRoutes(protected *gin.RouterGroup) {
	protected.GET("/users/me", func(c *gin.Context) {
		user.MeHandler(c, s.db)
	})
}