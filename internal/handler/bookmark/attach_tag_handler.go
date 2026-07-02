package bookmark

import (
	"net/http"
	"strconv"

  "bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AttachTagHandler(c *gin.Context, db *gorm.DB) {
	userID := c.GetInt64("userID")
	bookmark_repo := repository.NewBookmarkRepository(db)
  bookmarkID, err := strconv.ParseInt(c.Param("bookmark_id"), 10, 64)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid bookmark ID: " + err.Error()})
    return
  }
	bookmark, err := bookmark_repo.GetBookmarkByIDAndUserID(bookmarkID, userID)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to get bookmark: " + err.Error()})
    return
  }
  
	tag_repo := repository.NewTagRepository(db)
	tagID, err := strconv.ParseInt(c.Param("tag_id"), 10, 64)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid tag ID: " + err.Error()})
    return
  }
	tag, err := tag_repo.GetTagByIDAndUserID(tagID, userID)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to get tag: " + err.Error()})
    return
  }

	if bookmark.UserID != userID || userID != tag.UserID {
    c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
    return
  }
  
  bt, err := bookmark_repo.AttachTagToBookmark(userID, bookmark.ID, tag.ID)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to attach tag to bookmark: " + err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{
		"message": "Tag attached to bookmark successfully",
		"bookmark_tag": bt,
  })
}