package services

import (
	"net/http"

	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/events"
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

func (e *EventService) Get(identifier string) (dto.EventDto, bool) {
	var event model.EventModel
	e.DB.Preload("Monitors").Preload("Incident").Where("identifier = ?", identifier).First(&event)

	if util.CheckNull(event.Identifier) {
		return dto.EventDto{}, false
	}

	var monitors []string
	for i := range event.Monitors {
		monitors = append(monitors, event.Monitors[i].Identifier)
	}

	return dto.EventDto{
		Identifier:         event.Identifier,
		Title:              event.Title,
		Description:        event.Description,
		Type:               event.Type,
		Date:               event.Date,
		IncidentIdentifier: event.Incident.Identifier,
		Monitors:           monitors,
	}, true
}

func (e *EventService) Delete(identifier string) bool {
	var event model.EventModel
	e.DB.Preload("Monitors").Preload("Incident").Where("identifier = ?", identifier).First(&event)
	if util.CheckNull(event.Identifier) {
		return false
	}
	e.DB.Delete(&event)
	return true
}

func (e *EventService) Discover() []dto.EventDto {
	var events []model.EventModel
	e.DB.Preload("Monitors").Find(&events)

	eventDtos := make([]dto.EventDto, 0)
	for i := range events {
		event := events[i]

		var incident model.IncidentModel
		e.DB.Where("id = ?", event.IncidentID).First(&incident)

		var monitorDtos = make([]string, 0)
		for i := range event.Monitors {
			monitor := event.Monitors[i]
			monitorDtos = append(monitorDtos, monitor.Identifier)
		}

		eventDtos = append(eventDtos, dto.EventDto{
			Identifier:         event.Identifier,
			Title:              event.Title,
			Description:        event.Description,
			Type:               event.Type,
			Date:               event.Date,
			IncidentIdentifier: incident.Identifier,
			Monitors:           monitorDtos,
		})
	}

	return eventDtos
}

func (e *EventService) Validate(dto *dto.EventDto) (int, string) {
	if util.CheckNull(dto.Identifier) ||
		util.CheckNull(dto.Description) ||
		util.CheckNull(dto.Date) ||
		util.CheckNull(dto.Title) ||
		util.CheckNull(dto.Type) ||
		events.ToType(dto.Type) == events.Unkown {
		return http.StatusBadRequest, "invalid dto!"
	}

	var event model.EventModel
	e.DB.Where("identifier = ?", dto.Identifier).First(&event)
	if !util.CheckNull(event.Identifier) {
		return http.StatusConflict, "event already exists"
	}

	return http.StatusOK, ""
}

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
