package bookmark

import (
	"fmt"
	"net/http"

	"bookmarks/internal/cache"
	"bookmarks/internal/handler"
	"bookmarks/internal/models"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateHandler(c *gin.Context, db *gorm.DB, cache *cache.Cache) {
	repo := repository.NewBookmarkRepository(db)
	userID := c.GetInt64("userID")
	br := models.BookmarkRequest{
		Title:       c.PostForm("title"),
		URL:         c.PostForm("url"),
		Description: c.PostForm("description"),
		IsFavorite:  c.PostForm("is_favorite") == "true",
		IsArchived:  c.PostForm("is_archived") == "true",
	}
	b, err := repo.Create(userID, br)
	if err != nil {
		handler.RespondError(c, err)
		return
	}

	_ = cache.Delete(c.Request.Context(), fmt.Sprintf("user:stats:%d", userID))
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Bookmark created successfully",
		"bookmark": b.ToResponse(),
	})
}
