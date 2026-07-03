package user

import (
	"net/http"
	"bookmarks/internal/services"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatsHandler(c *gin.Context, db *gorm.DB) {
  userID := c.GetInt64("userID")
  bookmarkRepo := repository.NewBookmarkRepository(db)
  tagsRepo := repository.NewTagRepository(db)
  statisticService := services.NewStatisticService(bookmarkRepo, tagsRepo)
  stats, err := statisticService.CalculateUserStats(userID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"data": stats})
}