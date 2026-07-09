package main

import (
	"os"
	"log/slog"

	"bookmarks/internal/config"
	"bookmarks/internal/server"
	"bookmarks/internal/logger"

	"github.com/joho/godotenv"
)

func main() {
	slog.SetDefault(logger.New(os.Getenv("LOG_LEVEL")))

	_ = godotenv.Load()
	gormDB, err := config.NewGormDB()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
	}

	port := os.Getenv("APP_PORT")
	if port != "" && port[0] != ':' {
		port = ":" + port
	}

	server := server.New(port, gormDB)
	if err := server.Run(); err != nil {
		slog.Error("Failed to run server", "error", err)
	}
}