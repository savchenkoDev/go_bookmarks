package tag

import (
	"net/http"
	
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	t_resp := make([]models.TagResponse, len(tags))
	for i, t := range tags {
		t_resp[i] = t.ToResponse()
	}
	c.JSON(http.StatusOK, gin.H{
		"count": len(t_resp),
		"data": t_resp,
	})
}