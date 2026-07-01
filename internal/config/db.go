package config

import (
	"os"
	"log"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping database:", err)
		return nil, err
	}
	log.Println("Database pinged successfully")
	return db, nil
}