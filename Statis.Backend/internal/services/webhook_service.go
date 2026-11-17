package services

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
)

type WebhookService struct {
	DB *gorm.DB
}

func NewWebhookService(db *gorm.DB) *WebhookService {
	return &WebhookService{
		DB: db,
	}
}

func (e *WebhookService) Validate(dto *dto.WebhookDto) bool {

	return false
}

func (e *WebhookService) Add(dto *dto.WebhookDto) {
}

func (e *WebhookService) Update(dto *dto.WebhookDto) {
}
