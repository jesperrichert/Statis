package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
)

type IncidencesController struct {
	DB *gorm.DB
}

func NewIncidencesController(db *gorm.DB) *IncidencesController {
	return &IncidencesController{
		DB: db,
	}
}

// GET
func (c *IncidencesController) Get(ctx *gin.Context) {

	id := ctx.Param("id")

	var incident model.Incident
	c.DB.Where("identifier = ?", id).First(&incident)

	ctx.Data(http.StatusOK, "text/plain", []byte(incident.Content))
}

// POST
func (c *IncidencesController) Post(ctx *gin.Context) {

	var body dto.IncidentDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event model.Event
	c.DB.Where("identifier = ?", body.EventId).First(&event)

	if len(event.Identifier) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no event found"})
		return
	}

	c.DB.Create(&model.Incident{
		Identifier: body.Identifier,
		Content:    body.Content,
		EventID:    event.ID,
	})

	ctx.JSON(http.StatusCreated, gin.H{})
}
