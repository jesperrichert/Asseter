package config

import (
	"log"
	"os"

	"go.Asseter/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.APIAccess{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}
