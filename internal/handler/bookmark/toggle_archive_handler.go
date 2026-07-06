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

func ToggleArchiveHandler(c *gin.Context, db *gorm.DB) {
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
	c.JSON(http.StatusOK, gin.H{"message": message, "bookmark": b.ToResponse()})
}
