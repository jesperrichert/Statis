package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jespersen.zip.statis/internal/model"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("/app/db.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.MonitorModel{}, &model.EventModel{}, &model.IncidentModel{}, &model.WebhookModel{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}
