package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/util"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dto not valid})"})
		return
	}

	var event model.Event
	c.DB.Where("identifier = ?", body.Identifier).First(&event)

	if len(event.Identifier) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no event found"})
		return
	}

	c.DB.Create(&model.Event{
		Identifier:  body.Identifier,
		Title:       body.Title,
		Description: body.Description,
		Type:        body.Type,
		Date:        body.Date,
	})

	ctx.JSON(http.StatusCreated, gin.H{})
}
