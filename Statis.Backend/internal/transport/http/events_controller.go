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
	status, message := c.EventService.Validate(&body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}

	status = c.EventService.Add(&body)
	if status != http.StatusCreated {
		ctx.JSON(status, util.JsonResult("failed to add event", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// Post
func (c *EventsController) Update(ctx *gin.Context) {
	var body dto.EventDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	status := c.EventService.Update(ctx.Param("id"), &body)
	if status != http.StatusOK {
		ctx.JSON(status, gin.H{"error": "failed to update monitor!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
