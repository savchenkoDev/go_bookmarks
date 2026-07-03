package auth

import (
	"net/http"

  "bookmarks/internal/models"
	"bookmarks/internal/repository"
  
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandler(c *gin.Context, db *gorm.DB) {
  repo := repository.NewUserRepository(db)
  ur := models.UserRequest{
    Email: c.PostForm("email"),
    Password: c.PostForm("password"),
  }
  user, err := repo.Create(ur)
  if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user.ToResponse()})
}