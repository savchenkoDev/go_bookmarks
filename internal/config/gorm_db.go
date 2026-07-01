package config

import (
	"os"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func NewGormDB() (*gorm.DB, error) {
    return gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
}