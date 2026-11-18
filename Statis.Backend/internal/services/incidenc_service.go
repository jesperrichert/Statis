package services

import (
	"io"
	"net/http"
	"strings"

	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/util"
)

type IncidentService struct {
	DB *gorm.DB
}

func NewIncidentService(db *gorm.DB) *IncidentService {
	return &IncidentService{
		DB: db,
	}
}

func (e *IncidentService) Get(identifier string) (dto.IncidentDto, bool) {
	var incidentModel model.IncidentModel
	e.DB.Preload("Event").Where("identifier = ?", identifier).First(&incidentModel)

	if util.CheckNull(incidentModel.Identifier) {
		return dto.IncidentDto{}, false
	}

	return dto.IncidentDto{
		Identifier:      incidentModel.Identifier,
		EventIdentifier: incidentModel.Event.Identifier,
		Content:         incidentModel.Content,
	}, true
}

func (e *IncidentService) Delete(identifier string) bool {
	var incidentModel model.IncidentModel
	e.DB.Preload("Event").Where("identifier = ?", identifier).First(&incidentModel)
	if util.CheckNull(incidentModel.Identifier) {
		return false
	}

	e.DB.Delete(&incidentModel)
	return true
}

func (e *IncidentService) Validate(dto *dto.IncidentDto) (int, string) {

	if util.CheckNull(dto.Identifier) ||
		util.CheckNull(dto.EventIdentifier) ||
		util.CheckNull(dto.Content) {
		return http.StatusBadRequest, "invalid dto"
	}

	var incident model.IncidentModel
	e.DB.Where("identifier = ?", dto.Identifier).First(&incident)
	if !util.CheckNull(incident.Identifier) {
		return http.StatusConflict, "incident already exists!"
	}

	var event model.EventModel
	e.DB.Where("identifier = ?", dto.EventIdentifier).First(&event)
	if util.CheckNull(event.Identifier) {
		return http.StatusNoContent, "no event found"
	}

	if strings.HasPrefix(dto.Content, "http") &&
		strings.HasSuffix(dto.Content, ".mdx") {
		req, err := http.Get(dto.Content)
		if err != nil {
			return http.StatusBadRequest, "faild to fetch content"
		}
		defer req.Body.Close()
		bytes, _ := io.ReadAll(req.Body)

		dto.Content = string(bytes)
	}

	return http.StatusOK, ""
}

func (e *IncidentService) Add(dto *dto.IncidentDto) int {

	var event model.EventModel
	e.DB.Where("identifier = ?", dto.EventIdentifier).First(&event)

	var incident model.IncidentModel
	e.DB.Create(&model.IncidentModel{
		Identifier: dto.Identifier,
		Content:    dto.Content,
		EventID:    event.ID,
	}).Where("identifier = ?", dto.Identifier).First(&incident)

	if util.CheckNull(event.Identifier) {
		return http.StatusNotFound
	}

	event.IncidentID = incident.ID
	e.DB.Where("id = ?", event.ID).Updates(&event)

	return http.StatusCreated
}

func (e *IncidentService) Update(identifier string, dto *dto.IncidentDto) int {
	var incident model.IncidentModel
	e.DB.Where("identifier = ?", identifier).First(&incident)
	if util.CheckNull(incident.Identifier) {
		return http.StatusNotFound
	}

	if !util.CheckNull(dto.EventIdentifier) {
		var oldEvent model.EventModel
		e.DB.Where("id = ?", incident.EventID).First(&oldEvent)

		var event model.EventModel
		e.DB.Where("identifier = ?", dto.EventIdentifier).First(&event)
		if !util.CheckNull(event.Identifier) {
			oldEvent.IncidentID = 0
			oldEvent.Incident = model.IncidentModel{ID: 0}
			e.DB.Save(&oldEvent)

			incident.EventID = event.ID
			event.IncidentID = incident.ID
			e.DB.Where("id = ?", event.ID).Updates(&event)
		}
	} else {
		var oldEvent model.EventModel
		e.DB.Where("id = ?", incident.EventID).First(&oldEvent)
		oldEvent.IncidentID = 0
		oldEvent.Incident = model.IncidentModel{ID: 0}
		e.DB.Save(&oldEvent)
	}

	if !util.CheckNull(dto.Content) {
		incident.Content = dto.Content
	}

	e.DB.Where("id = ?", incident.ID).Updates(&incident)
	return http.StatusOK
}
