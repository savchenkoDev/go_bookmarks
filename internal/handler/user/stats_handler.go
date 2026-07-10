package user

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"bookmarks/internal/cache"
	"bookmarks/internal/handler"
	"bookmarks/internal/models"
	"bookmarks/internal/repository"
	"bookmarks/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatsHandler(c *gin.Context, db *gorm.DB, cache *cache.Cache) {
	userID := c.GetInt64("userID")

	redisKey := fmt.Sprintf("stats:user:%d", userID)
  var stats models.UserStats
	if ok, _ := cache.Get(c.Request.Context(), redisKey, &stats); ok {
		slog.Info("Stats found in cache", "userID", userID)
		c.JSON(http.StatusOK, gin.H{"data": stats})
		return
	}

	bookmarkRepo := repository.NewBookmarkRepository(db)
	tagsRepo := repository.NewTagRepository(db)
	statisticService := services.NewStatisticService(bookmarkRepo, tagsRepo)
	stats, err := statisticService.CalculateUserStats(userID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	slog.Info("Stats calculated", "userID", userID)
	_ = cache.Set(c.Request.Context(), redisKey, stats, 5*time.Minute)
	slog.Info("Stats cached", "userID", userID)
	c.JSON(http.StatusOK, gin.H{"data": stats})
}
