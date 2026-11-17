package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/services"
	"jespersen.zip.statis/internal/util"
)

type EventsController struct {
	DB           *gorm.DB
	EventService *services.EventService
}

func NewEventsController(db *gorm.DB, service *services.EventService) *EventsController {
	return &EventsController{
		DB:           db,
		EventService: service,
	}
}

// GET
func (c *EventsController) Get(ctx *gin.Context) {

	var events []model.EventModel
	c.DB.Preload("Monitors").Find(&events)

	eventDtos := make([]dto.EventDto, 0)
	for i := range events {
		event := events[i]

		var incident model.IncidentModel
		c.DB.Where("id = ?", event.IncidentID).First(&incident)

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

	ctx.JSON(http.StatusOK, eventDtos)
}

// Post
func (c *EventsController) Post(ctx *gin.Context) {
	var body dto.EventDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if util.CheckNull(body.Identifier) ||
		util.CheckNull(body.Description) ||
		util.CheckNull(body.Date) ||
		util.CheckNull(body.Title) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dto not valid"})
		return
	}

	var event model.EventModel
	c.DB.Where("identifier = ?", body.Identifier).First(&event)
	if len(event.Identifier) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "event already exists"})
		return
	}

	var monitors []*model.MonitorModel
	for i := range body.Monitors {
		id := body.Monitors[i]
		var monitor model.MonitorModel
		c.DB.Where("identifier = ?", id).First(&monitor)
		monitors = append(monitors, &monitor)
	}

	c.DB.Create(&model.EventModel{
		Identifier:  body.Identifier,
		Title:       body.Title,
		Description: body.Description,
		Type:        body.Type,
		Date:        body.Date,
		Monitors:    monitors,
	})

	ctx.JSON(http.StatusCreated, gin.H{})
}

// Post
func (c *EventsController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var event model.EventModel
	c.DB.Where("identifier = ?", id).First(&event)
	if util.CheckNull(event.Identifier) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	var body dto.EventDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	if !util.CheckNull(body.Type) {
		event.Type = body.Type
	}
	if !util.CheckNull(body.Title) {
		event.Title = body.Title
	}
	if !util.CheckNull(body.Description) {
		event.Description = body.Description
	}
	if !util.CheckNull(body.Date) {
		event.Date = body.Date
	}
	if len(body.Monitors) > 0 {

		var monitors []*model.MonitorModel
		for i := range body.Monitors {
			id := body.Monitors[i]
			var monitor model.MonitorModel
			c.DB.Where("identifier = ?", id).First(&monitor)
			if !util.CheckNull(monitor.Identifier) {
				monitors = append(monitors, &monitor)
			}
		}
		event.Monitors = monitors
	}
	// TODO: Fix "Old Incidences remove"
	c.DB.Where("id = ?", event.ID).Updates(&event)
}
