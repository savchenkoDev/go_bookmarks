package bookmark

import (
	"database/sql"
	"net/http"
	"strconv"

	"bookmarks/internal/repository"
	
	"github.com/gin-gonic/gin"
)

func ShowHandler(c *gin.Context, db *sql.DB) {
	userID := c.GetInt64("userID")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}
	bookmark, err := repository.GetBookmarkByIDAndUserID(db, id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookmark": bookmark.ToResponse()})
}