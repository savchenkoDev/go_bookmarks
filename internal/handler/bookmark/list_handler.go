package bookmark

import (
	"github.com/gin-gonic/gin"
	"bookmarks/internal/repository"
	"net/http"

	"bookmarks/internal/models"
	"gorm.io/gorm"
)

func ListHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewBookmarkRepository(db)
	userID := c.GetInt64("userID")
	bookmarks, err := repo.GetBookmarksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	b_resp := make([]models.BookmarkResponse, len(bookmarks))
	for i, b := range bookmarks {
		b_resp[i] = b.ToResponse()
	}
	c.JSON(http.StatusOK, gin.H{
		"count": len(b_resp),
		"data": b_resp,
	})
}