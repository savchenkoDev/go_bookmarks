package user

import (
	"github.com/gin-gonic/gin"
	"bookmarks/internal/user"
	"database/sql"
	"net/http"
)

func CreateHandler(c *gin.Context, db *sql.DB) {
  ur := user.UserRequest{
    Email: c.PostForm("email"),
    Password: c.PostForm("password"),
  }
  user, err := ur.Create(db)
  if err != nil {
    error_message := "Failed to create user: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": user,
	})
}