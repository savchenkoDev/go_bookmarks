package auth

import (
	"net/http"

	"bookmarks/internal/jwt"
	"bookmarks/internal/repository"
	"bookmarks/internal/models"
	
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const ERROR_MESSAGE = "Invalid email or password"

func LoginHandler(c *gin.Context, db *gorm.DB) {
	var ur models.UserRequest
	c.ShouldBindJSON(&ur)

	repo := repository.NewUserRepository(db)
	user, err := repo.GetUserByEmail(ur.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": ERROR_MESSAGE})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(ur.Password))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": ERROR_MESSAGE})
		return
	}
  token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": ERROR_MESSAGE})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}