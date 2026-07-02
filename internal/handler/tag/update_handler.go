package tag

import (
  "strconv"
	"log"
	"net/http"
	
	"bookmarks/internal/models"
  "bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateHandler(c *gin.Context, db *gorm.DB) {
	repo := repository.NewTagRepository(db)
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}
	userID := c.GetInt64("userID")
  var tr models.TagUpdateRequest
  if err := c.ShouldBindJSON(&tr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
  log.Println("handler tr:", tr)
  t, err := repo.Update(userID, idInt, tr)
  if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to update tag: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tag updated successfully", "tag": t.ToResponse()})
}