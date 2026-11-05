package http

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
)

type EventsController struct {
	db *gorm.DB
}

func NewEventsController(db *gorm.DB) *EventsController {
	return &EventsController{
		db: db,
	}
}

// GET
func (e *EventsController) Get(ctx *gin.Context) {

	data, ok := os.LookupEnv("EVENTS")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	res, err := http.Get(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer res.Body.Close()

	var body []dto.EventDto
	json.NewDecoder(res.Body).Decode(&body)

	ctx.JSON(http.StatusOK, body)
}
