package tag

import (
	"net/http"
  "log"

	"bookmarks/internal/models"
	"bookmarks/internal/repository"
  
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

func CreateHandler(c *gin.Context, db *gorm.DB) {
  tr := models.TagRequest{}
  if err := c.ShouldBindJSON(&tr); err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
    return
  }
  tr.UserID = c.GetInt64("userID")
  log.Println(tr)
  repo := repository.NewTagRepository(db)
  t, err := repo.Create(tr)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusCreated, t)
}