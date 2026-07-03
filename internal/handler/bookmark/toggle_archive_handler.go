package bookmark

import (
  "strconv"
	"net/http"
	
  "bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ToggleArchiveHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewBookmarkRepository(db)
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}
  b, err := repo.ToggleArchive(idInt)
  if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to toggle archive: " + err.Error()})
		return
	}
	message := ""
	if b.IsArchived {
		message = "Bookmark archived successfully"
	} else {
		message = "Bookmark unarchived successfully"
	}
	c.JSON(http.StatusOK, gin.H{"message": message, "bookmark": b.ToResponse()})
}