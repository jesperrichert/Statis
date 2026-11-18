package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
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

	eventDto, ok := c.EventService.Get(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, eventDto)
}

// Delete
func (c *EventsController) Delete(ctx *gin.Context) {

	ok := c.EventService.Delete(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// GET
func (c *EventsController) Discover(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.EventService.Discover())
}

// Post
func (c *EventsController) Post(ctx *gin.Context) {
	var body dto.EventDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.JsonResult("failed to parse dto", map[string]any{}))
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

	status, message := c.EventService.Validate(&body)
	if status != http.StatusOK && status != http.StatusConflict {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}

	status = c.EventService.Update(ctx.Param("id"), &body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult("failed to update monitor!", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
