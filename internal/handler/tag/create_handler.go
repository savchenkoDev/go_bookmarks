package tag

import (
	"net/http"

	apperr "bookmarks/internal/errors"
	"bookmarks/internal/handler"
	"bookmarks/internal/models"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateHandler(c *gin.Context, db *gorm.DB) {
	tr := models.TagRequest{}
	if err := c.ShouldBindJSON(&tr); err != nil {
		handler.RespondError(c, apperr.RecordInvalidError())
		return
	}
	tr.UserID = c.GetInt64("userID")

	repo := repository.NewTagRepository(db)
	t, err := repo.Create(tr)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, t)
}
