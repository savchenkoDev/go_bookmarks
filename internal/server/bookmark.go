package server

import (
	"bookmarks/internal/handler/bookmark"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerBookmarkProtectedRoutes(protected *gin.RouterGroup) {
	protected.GET("/bookmarks", func(c *gin.Context) {
		bookmark.ListHandler(c, s.db)
	})

	protected.POST("/bookmarks", func(c *gin.Context) {
		bookmark.CreateHandler(c, s.db)
	})

	protected.GET("/bookmarks/:id", func(c *gin.Context) {
		bookmark.ShowHandler(c, s.db)
	})

	protected.PUT("/bookmarks/:id", func(c *gin.Context) {
		bookmark.UpdateHandler(c, s.db)
	})

	protected.DELETE("/bookmarks/:id", func(c *gin.Context) {
		bookmark.DeleteHandler(c, s.db)
	})
}