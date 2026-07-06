package bookmark

import (
	"net/http"
	"strconv"

	apperr "bookmarks/internal/errors"
	"bookmarks/internal/handler"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DetachTagHandler(c *gin.Context, db *gorm.DB) {
	bookmarkTagRepo := repository.NewBookmarkTagRepository(db)
	bookmarkTagID, err := strconv.ParseInt(c.Param("bookmark_tag_id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}

	err = bookmarkTagRepo.DetachTagFromBookmark(bookmarkTagID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag detached from bookmark successfully"})
}
