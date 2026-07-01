package user

import (
	"github.com/gin-gonic/gin"
	"bookmarks/internal/repository"
	"database/sql"
	"net/http"
)

func MeHandler(c *gin.Context, db *sql.DB) {
  userID := c.GetInt64("userID")
  u, err := repository.GetUserByID(db, userID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"user": u.ToResponse()})
}