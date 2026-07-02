package tag

import (
	"net/http"
	"strconv"

	"bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewTagRepository(db)
	userID := c.GetInt64("userID")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}
	t, err := repo.GetTagByIDAndUserID(id, userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to get bookmark: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bookmark retrieved successfully", "bookmark": t.ToResponse()})
}