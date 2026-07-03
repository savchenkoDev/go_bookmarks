package bookmark

import (
	"net/http"
	"strconv"

  "bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DetachTagHandler(c *gin.Context, db *gorm.DB) {
	bookmarkTagRepo := repository.NewBookmarkTagRepository(db)
  bookmarkTagID, err := strconv.ParseInt(c.Param("bookmark_tag_id"), 10, 64)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid bookmark tag ID: " + err.Error()})
    return
  }
  
  err = bookmarkTagRepo.DetachTagFromBookmark(bookmarkTagID)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to detach tag from bookmark: " + err.Error()})
    return
  }
  
  c.JSON(http.StatusOK, gin.H{"message": "Tag detached from bookmark successfully"})
}
