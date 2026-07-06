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

func AttachTagHandler(c *gin.Context, db *gorm.DB) {
	userID := c.GetInt64("userID")
	bookmarkRepo := repository.NewBookmarkRepository(db)
	bookmarkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}
	bookmark, err := bookmarkRepo.GetBookmarkByIDAndUserID(bookmarkID, userID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}

	tagRepo := repository.NewTagRepository(db)
	tagID, err := strconv.ParseInt(c.Param("tag_id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}
	tag, err := tagRepo.GetTagByIDAndUserID(tagID, userID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}

	if bookmark.UserID != userID || userID != tag.UserID {
		handler.RespondError(c, apperr.ForbiddenError())
		return
	}

	bt, err := bookmarkRepo.AttachTagToBookmark(userID, bookmark.ID, tag.ID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "Tag attached to bookmark successfully",
		"bookmark_tag": bt,
	})
}
