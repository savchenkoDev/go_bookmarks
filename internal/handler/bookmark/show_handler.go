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

func ShowHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewBookmarkRepository(db)
	userID := c.GetInt64("userID")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}
	bookmark, err := repo.GetBookmarkByIDAndUserID(id, userID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bookmark retrieved successfully", "bookmark": bookmark.ToResponse()})
}
