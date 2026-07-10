package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewGormDB() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
}
