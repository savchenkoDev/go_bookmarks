package tag

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
	"bookmarks/internal/repository"
	"gorm.io/gorm"
)

func DeleteHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewTagRepository(db)
	userID := c.GetInt64("userID")
  id := c.Param("id")
  idInt, err := strconv.ParseInt(id, 10, 64)
  if err != nil {
    error_message := "Invalid ID: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  err = repo.Delete(userID, idInt)
  if err != nil {
    error_message := "Failed to delete tag: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}