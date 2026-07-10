package tag

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

func DeleteHandler(c *gin.Context, db *gorm.DB, cache *cache.Cache) {
	repo := repository.NewTagRepository(db)
	userID := c.GetInt64("userID")
	idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}
	err = repo.Delete(userID, idInt)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	_ = cache.Delete(c.Request.Context(), fmt.Sprintf("user:tags:%d", userID))
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
