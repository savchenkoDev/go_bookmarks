package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the bookmarks API!",
		"version": "1.0.0",
		"status": "ok",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}