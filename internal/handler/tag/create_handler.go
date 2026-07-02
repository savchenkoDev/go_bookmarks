package tag

import (
	"net/http"

	"bookmarks/internal/models"
	"bookmarks/internal/repository"
  
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

func CreateHandler(c *gin.Context, db *gorm.DB) {
  repo := repository.NewTagRepository(db)
	userID := c.GetInt64("userID")
  tr := models.TagRequest{
    UserID: userID,
    Name: c.PostForm("name"),
  }
  t, err := repo.Create(tr)
  if err != nil {
    error_message := "Failed to create tag: " + err.Error()
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error_message})
    return
  }
  c.JSON(http.StatusCreated, gin.H{
		"message": "Tag created successfully",
		"tag": t.ToResponse(),
	})
}