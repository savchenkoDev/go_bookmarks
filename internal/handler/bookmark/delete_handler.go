package bookmark

import (
	"database/sql"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
	"bookmarks/internal/repository"
)

func DeleteHandler(c *gin.Context, db *sql.DB) {
	userID := c.GetInt64("userID")
  id := c.Param("id")
  idInt, err := strconv.ParseInt(id, 10, 64)
  if err != nil {
    error_message := "Invalid ID: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  b, err := repository.GetBookmarkByIDAndUserID(db, idInt, userID)
  if err != nil {
    error_message := "Bookmark not found: " + err.Error()
    c.JSON(http.StatusNotFound, gin.H{"error": error_message})
    return
  }
  err = b.Delete(db)
  if err != nil {
    error_message := "Failed to delete bookmark: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Bookmark deleted successfully"})
}