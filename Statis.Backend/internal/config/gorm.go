package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/model"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("/app/db.sqlite"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Monitor{}, &model.Event{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}
