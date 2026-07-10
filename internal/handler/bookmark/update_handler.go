package bookmark

import (
	"fmt"
	"net/http"
	"strconv"

	"bookmarks/internal/cache"
	apperr "bookmarks/internal/errors"
	"bookmarks/internal/handler"
	"bookmarks/internal/models"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateHandler(c *gin.Context, db *gorm.DB, cache *cache.Cache) {
	repo := repository.NewBookmarkRepository(db)
	idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}
	userID := c.GetInt64("userID")
	var br models.BookmarkUpdateRequest
	if err := c.ShouldBindJSON(&br); err != nil {
		handler.RespondError(c, apperr.RecordInvalidError())
		return
	}

	b, err := repo.Update(userID, idInt, br)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	_ = cache.Delete(c.Request.Context(), fmt.Sprintf("user:stats:%d", userID))
	c.JSON(http.StatusOK, gin.H{"message": "Bookmark updated successfully", "bookmark": b.ToResponse()})
}
