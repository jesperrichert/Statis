package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
)

type EventsController struct {
	DB *gorm.DB
}

func NewEventsController(db *gorm.DB) *EventsController {
	return &EventsController{
		DB: db,
	}
}

// GET
func (c *EventsController) Get(ctx *gin.Context) {

	var events []model.Event
	c.DB.Find(&events)

	eventDtos := make([]dto.EventDto, 0)
	for i := range events {
		event := events[i]
		eventDtos = append(eventDtos, dto.EventDto{
			Identifier:  event.Identifier,
			Title:       event.Title,
			Description: event.Description,
			Type:        event.Type,
			Date:        event.Date,
		})
	}

	ctx.JSON(http.StatusOK, eventDtos)
}

// Post
func (c *EventsController) Post(ctx *gin.Context) {

}
