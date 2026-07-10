package bookmark

import (
	"fmt"
	"net/http"
	"strconv"

	"bookmarks/internal/cache"
	apperr "bookmarks/internal/errors"
	"bookmarks/internal/handler"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ToggleArchiveHandler(c *gin.Context, db *gorm.DB, cache *cache.Cache) {
	repo := repository.NewBookmarkRepository(db)
	idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}
	b, err := repo.ToggleArchive(idInt)
	if err != nil {
		handler.RespondError(c, err)
		return
	}

	message := "Bookmark unarchived successfully"
	if b.IsArchived {
		message = "Bookmark archived successfully"
	}
	_ = cache.Delete(c.Request.Context(), fmt.Sprintf("user:stats:%d", b.UserID))
	c.JSON(http.StatusOK, gin.H{"message": message, "bookmark": b.ToResponse()})
}
