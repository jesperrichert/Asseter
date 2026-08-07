package config

import (
	"log"

	"go.Asseter/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
[TODO]
Replace with factory
*/

func NewDatabase() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("dev.db"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Example{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}
