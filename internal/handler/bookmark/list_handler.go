package bookmark

import (
	"net/http"
	"strconv"

	"bookmarks/internal/repository"
	"bookmarks/internal/models"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListHandler(c *gin.Context, db *gorm.DB) {
	params, _ := parseBookmarkListParams(c)
	repo := repository.NewBookmarkRepository(db)
	userID := c.GetInt64("userID")
	result, err := repo.ListByUserID(userID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, result)
}

func parseBookmarkListParams(c *gin.Context) (models.BookmarkListParams, error) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

    if page < 1 {
        page = 1
    }
    if perPage < 1 {
        perPage = 20
    }
    if perPage > 100 {
        perPage = 100
    }

    params := models.BookmarkListParams{
        Page:    page,
        PerPage: perPage,
        Sort:    c.DefaultQuery("sort", "created_at"),
        Order:   c.DefaultQuery("order", "desc"),
        Tag:     c.Query("tag"),
        Query:   c.Query("q"),
    }

    if v := c.Query("is_favorite"); v != "" {
        b := v == "true"
        params.IsFavorite = &b
    }
    if v := c.Query("is_archived"); v != "" {
        b := v == "true"
        params.IsArchived = &b
    }

    return params, nil
}