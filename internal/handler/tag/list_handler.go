package tag

import (
	"net/http"

	"bookmarks/internal/handler"
	"bookmarks/internal/models"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewTagRepository(db)
	userID := c.GetInt64("userID")
	tags, err := repo.GetTagsByUserID(userID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}

	tResp := make([]models.TagResponse, len(tags))
	for i, t := range tags {
		tResp[i] = t.ToResponse()
	}
	c.JSON(http.StatusOK, gin.H{
		"count": len(tResp),
		"data":  tResp,
	})
}
