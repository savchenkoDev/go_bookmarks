package handler

import (
	"net/http"

	"bookmarks/internal/jwt"
	"bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const ERROR_MESSAGE = "Invalid email or password"

func AuthHandler(c *gin.Context, db *gorm.DB) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	repo := repository.NewUserRepository(db)
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": ERROR_MESSAGE})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
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
	return
}