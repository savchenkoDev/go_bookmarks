package bookmark

import (
	"fmt"
	"net/http"
	"strconv"

	apperr "bookmarks/internal/errors"
	"bookmarks/internal/handler"
	"bookmarks/internal/repository"
	"bookmarks/internal/cache"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DetachTagHandler(c *gin.Context, db *gorm.DB, cache *cache.Cache) {
	bookmarkTagRepo := repository.NewBookmarkTagRepository(db)
	bookmarkTagID, err := strconv.ParseInt(c.Param("bookmark_tag_id"), 10, 64)
	userID := c.GetInt64("userID")
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}

	err = bookmarkTagRepo.DetachTagFromBookmark(bookmarkTagID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	_ = cache.Delete(c.Request.Context(), fmt.Sprintf("user:stats:%d", userID))
	c.JSON(http.StatusOK, gin.H{"message": "Tag detached from bookmark successfully"})
}
