package bookmark

import (
	"github.com/gin-gonic/gin"
	"bookmarks/internal/repository"
	"database/sql"
	"net/http"

	"bookmarks/internal/bookmark"
)

func ListHandler(c *gin.Context, db *sql.DB) {
	userID := c.GetInt64("userID")
	bookmarks, err := repository.GetBookmarksByUserID(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	b_resp := make([]bookmark.BookmarkResponse, len(bookmarks))
	for i, b := range bookmarks {
		b_resp[i] = b.ToResponse()
	}
	c.JSON(http.StatusOK, gin.H{
		"count": len(b_resp),
		"data": b_resp,
	})
}