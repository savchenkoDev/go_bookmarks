package main

import (
	"log/slog"
	"os"
	"time"

	"bookmarks/internal/config"
	"bookmarks/internal/logger"
	"bookmarks/internal/server"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func main() {
	log := logger.New(os.Getenv("LOG_LEVEL"))
	slog.SetDefault(log)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      os.Getenv("APP_ENV"),
		Release:          os.Getenv("APP_VERSION"),
		TracesSampleRate: 0.2,
	})
	if err != nil {
		slog.Error("sentry init failed", "error", err)
	}
	defer sentry.Flush(2 * time.Second)

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
