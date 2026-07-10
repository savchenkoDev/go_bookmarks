package server

import (
	"bookmarks/internal/handler/tag"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerTagProtectedRoutes(protected *gin.RouterGroup) {
	protected.GET("/tags", func(c *gin.Context) {
		tag.ListHandler(c, s.db)
	})

	protected.POST("/tags", func(c *gin.Context) {
		tag.CreateHandler(c, s.db, s.cache)
	})

	protected.GET("/tags/:id", func(c *gin.Context) {
		tag.ShowHandler(c, s.db)
	})

	protected.PUT("/tags/:id", func(c *gin.Context) {
		tag.UpdateHandler(c, s.db, s.cache)
	})

	protected.PATCH("/tags/:id", func(c *gin.Context) {
		tag.UpdateHandler(c, s.db, s.cache)
	})

	protected.DELETE("/tags/:id", func(c *gin.Context) {
		tag.DeleteHandler(c, s.db, s.cache)
	})
}
