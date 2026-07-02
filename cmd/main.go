package main

import (
	"log"
	"os"

	"bookmarks/internal/config"
	"bookmarks/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	gormDB, err := config.NewGormDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port != "" && port[0] != ':' {
		port = ":" + port
	}

	server := server.New(port, gormDB)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}