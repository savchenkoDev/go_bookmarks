package user

import (
	"bookmarks/internal/repository"
  
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func MeHandler(c *gin.Context, db *gorm.DB) {
  repo := repository.NewUserRepository(db)
  userID := c.GetInt64("userID")
  u, err := repo.GetUserByID(userID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"user": u.ToResponse()})
}