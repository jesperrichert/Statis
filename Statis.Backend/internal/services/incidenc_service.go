package services

import (
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
)

type IncidentService struct {
	DB *gorm.DB
}

func NewIncidentService(db *gorm.DB) *IncidentService {
	return &IncidentService{
		DB: db,
	}
}

func (e *IncidentService) Validate(dto *dto.IncidentDto) bool {

	return false
}

func (e *IncidentService) Add(dto *dto.IncidentDto) {
}

func (e *IncidentService) Update(dto *dto.IncidentDto) {
}
