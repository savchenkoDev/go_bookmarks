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
		bookmark.CreateHandler(c, s.db, s.cache)
	})

	protected.POST("/bookmarks/:id/tags/:tag_id", func(c *gin.Context) {
		bookmark.AttachTagHandler(c, s.db, s.cache)
	})

	protected.DELETE("/bookmarks/:id/bookmark_tags/:bookmark_tag_id", func(c *gin.Context) {
		bookmark.DetachTagHandler(c, s.db, s.cache)
	})

	protected.GET("/bookmarks/:id", func(c *gin.Context) {
		bookmark.ShowHandler(c, s.db)
	})

	protected.PUT("/bookmarks/:id/toggle_archive", func(c *gin.Context) {
		bookmark.ToggleArchiveHandler(c, s.db, s.cache)
	})

	protected.PUT("/bookmarks/:id", func(c *gin.Context) {
		bookmark.UpdateHandler(c, s.db, s.cache)
	})

	protected.DELETE("/bookmarks/:id", func(c *gin.Context) {
		bookmark.DeleteHandler(c, s.db, s.cache)
	})
}
