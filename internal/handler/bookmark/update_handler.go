package bookmark

import (
	"github.com/gin-gonic/gin"
	"bookmarks/internal/bookmark"
	"database/sql"
	"net/http"
  "strconv"
)

func UpdateHandler(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}
	userID := c.GetInt64("userID")
  var br bookmark.BookmarkUpdateRequest
  if err := c.ShouldBindJSON(&br); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

  b, err := br.Update(db, idInt, userID)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to update bookmark: " + err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{
		"message": "Bookmark updated successfully",
		"bookmark": b.ToResponse(),
	})
}