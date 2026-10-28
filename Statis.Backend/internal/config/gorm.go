package config

import (
	"log"

	"jespersen.zip.statis/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("/app/db.sqlite"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Example{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}
