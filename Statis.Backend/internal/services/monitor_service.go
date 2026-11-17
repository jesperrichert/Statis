package services

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
)

type MonitorService struct {
	DB *gorm.DB
}

func NewMonitorService(db *gorm.DB) *MonitorService {
	return &MonitorService{
		DB: db,
	}
}

func (e *MonitorService) Validate(dto *dto.MonitorDto[any]) bool {

	return false
}

func (e *MonitorService) Add(dto *dto.MonitorDto[any]) {
}

func (e *MonitorService) Update(dto *dto.MonitorDto[any]) {
}
