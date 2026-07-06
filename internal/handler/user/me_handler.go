package user

import (
	"net/http"

	"bookmarks/internal/handler"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MeHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	userID := c.GetInt64("userID")
	u, err := repo.GetUserByID(userID)
	if err != nil {
		handler.RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": u.ToResponse()})
}
