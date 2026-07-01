package bookmark

import (
	"github.com/gin-gonic/gin"
	"bookmarks/internal/bookmark"
	"database/sql"
	"net/http"
)

func CreateHandler(c *gin.Context, db *sql.DB) {
	userID := c.GetInt64("userID")
  br := bookmark.BookmarkRequest{
    Title: c.PostForm("title"),
    URL: c.PostForm("url"),
    Description: c.PostForm("description"),
    IsFavorite: c.PostForm("is_favorite") == "true",
    IsArchived: c.PostForm("is_archived") == "true",
  }
  b, err := br.Create(db, userID)
  if err != nil {
    error_message := "Failed to create bookmark: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  c.JSON(http.StatusCreated, gin.H{
		"message": "Bookmark created successfully",
		"bookmark": b,
	})
}