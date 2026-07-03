package server

import (
	"bookmarks/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerUserRoutes(protected *gin.RouterGroup) {
	protected.GET("/me", func(c *gin.Context) {
		user.MeHandler(c, s.db)
	})

	protected.GET("/stats", func(c *gin.Context) {
		user.StatsHandler(c, s.db)
	})
}