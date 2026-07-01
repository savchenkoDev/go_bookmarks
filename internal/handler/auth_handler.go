package handler

import (
	"database/sql"
	"net/http"

	"bookmarks/internal/jwt"
	"bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const ERROR_MESSAGE = "Invalid email or password"

func AuthHandler(c *gin.Context, db *sql.DB) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := repository.GetUserByEmail(db, email)
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