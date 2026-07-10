package tag

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
	repo := repository.NewTagRepository(db)
	idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handler.RespondError(c, apperr.InvalidIDError())
		return
	}
	userID := c.GetInt64("userID")
	var tr models.TagUpdateRequest
	if err := c.ShouldBindJSON(&tr); err != nil {
		handler.RespondError(c, apperr.RecordInvalidError())
		return
	}

	t, err := repo.Update(userID, idInt, tr)
	if err != nil {
		handler.RespondError(c, err)
		return
	}

	_ = cache.Delete(c.Request.Context(), fmt.Sprintf("user:tags:%d", userID))
	c.JSON(http.StatusOK, gin.H{"message": "Tag updated successfully", "tag": t.ToResponse()})
}
