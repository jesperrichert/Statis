package services

import (
	"net/http"

	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/util"
)

type EventService struct {
	DB *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{
		DB: db,
	}
}

func (e *EventService) Validate(dto *dto.EventDto) (int, string) {
	if util.CheckNull(dto.Identifier) ||
		util.CheckNull(dto.Description) ||
		util.CheckNull(dto.Date) ||
		util.CheckNull(dto.Title) {

		return http.StatusBadRequest, "invalid dto!"
	}

	var event model.EventModel
	e.DB.Where("identifier = ?", dto.Identifier).First(&event)
	if len(event.Identifier) > 0 {
		return http.StatusConflict, "event already exists"
	}

	return http.StatusOK, ""
}

// Success: http.StatusCreated
func (e *EventService) Add(dto *dto.EventDto) int {

	var monitors []*model.MonitorModel
	for i := range dto.Monitors {
		id := dto.Monitors[i]
		var monitor model.MonitorModel
		e.DB.Where("identifier = ?", id).First(&monitor)
		monitors = append(monitors, &monitor)
	}

	e.DB.Create(&model.EventModel{
		Identifier:  dto.Identifier,
		Title:       dto.Title,
		Description: dto.Description,
		Type:        dto.Type,
		Date:        dto.Date,
		Monitors:    monitors,
	})

	return http.StatusCreated
}

func (e *EventService) Update(identifier string, dto *dto.EventDto) int {
	var event model.EventModel
	e.DB.Where("identifier = ?", identifier).First(&event)
	if util.CheckNull(event.Identifier) {
		return http.StatusNotFound
	}

	if !util.CheckNull(dto.Type) {
		event.Type = dto.Type
	}
	if !util.CheckNull(dto.Title) {
		event.Title = dto.Title
	}
	if !util.CheckNull(dto.Description) {
		event.Description = dto.Description
	}
	if !util.CheckNull(dto.Date) {
		event.Date = dto.Date
	}
	if !util.CheckNullLen(dto.Monitors) {

		var monitors []*model.MonitorModel
		for i := range dto.Monitors {
			id := dto.Monitors[i]
			var monitor model.MonitorModel
			e.DB.Where("identifier = ?", id).First(&monitor)
			if !util.CheckNull(monitor.Identifier) {
				monitors = append(monitors, &monitor)
			}
		}
		event.Monitors = monitors
	} else {
		event.Monitors = []*model.MonitorModel{}
	}

	if event.IncidentID != 0 {
		var incident model.IncidentModel
		e.DB.Where("id = ?", event.IncidentID).First(&incident)
		if incident.EventID != event.ID {
			event.IncidentID = 0
		}
	}

	e.DB.Where("id = ?", event.ID).Updates(&event)

	return http.StatusOK
}
