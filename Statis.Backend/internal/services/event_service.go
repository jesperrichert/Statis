package services

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
)

type EventService struct {
	DB *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{
		DB: db,
	}
}

func (e *EventService) Validate(dto *dto.EventDto) bool {

	return false
}

func (e *EventService) Add(dto *dto.EventDto) {
}

func (e *EventService) Update(dto *dto.EventDto) {
}
