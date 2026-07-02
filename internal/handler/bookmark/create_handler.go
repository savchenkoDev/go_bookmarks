package bookmark

import (
	"net/http"

	"bookmarks/internal/models"
	"bookmarks/internal/repository"
  
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

func CreateHandler(c *gin.Context, db *gorm.DB) {
  repo := repository.NewBookmarkRepository(db)
	userID := c.GetInt64("userID")
  br := models.BookmarkRequest{
    Title: c.PostForm("title"),
    URL: c.PostForm("url"),
    Description: c.PostForm("description"),
    IsFavorite: c.PostForm("is_favorite") == "true",
    IsArchived: c.PostForm("is_archived") == "true",
  }
  b, err := repo.Create(userID, br)
  if err != nil {
    error_message := "Failed to create bookmark: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  c.JSON(http.StatusCreated, gin.H{
		"message": "Bookmark created successfully",
		"bookmark": b.ToResponse(),
	})
}