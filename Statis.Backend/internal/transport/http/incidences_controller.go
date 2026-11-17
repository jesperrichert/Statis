package http

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/services"
	"jespersen.zip.statis/internal/util"
)

type IncidencesController struct {
	DB              *gorm.DB
	IncidentService *services.IncidentService
}

func NewIncidencesController(db *gorm.DB, service *services.IncidentService) *IncidencesController {
	return &IncidencesController{
		DB:              db,
		IncidentService: service,
	}
}

// GET
func (c *IncidencesController) Get(ctx *gin.Context) {

	id := ctx.Param("id")

	var incident model.IncidentModel
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
	if util.CheckNull(body.Identifier) ||
		util.CheckNull(body.EventIdentifier) ||
		util.CheckNull(body.Content) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dto not valid})"})
		return
	}

	if strings.HasPrefix(body.Content, "http") &&
		strings.HasSuffix(body.Content, ".mdx") {
		req, err := http.Get(body.Content)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		defer req.Body.Close()
		bytes, _ := io.ReadAll(req.Body)
		body.Content = string(bytes)
	}

	var incident model.IncidentModel
	c.DB.Where("identifier = ?", body.Identifier).First(&incident)
	if len(incident.Identifier) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "incident already exists!"})
		return
	}

	var event model.EventModel
	c.DB.Where("identifier = ?", body.EventIdentifier).First(&event)
	if len(event.Identifier) <= 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"error": "no event found"})
		return
	}

	c.DB.Create(&model.IncidentModel{
		Identifier: body.Identifier,
		Content:    body.Content,
		EventID:    event.ID,
	})
	c.DB.Where("identifier = ?", body.Identifier).First(&incident)
	event.IncidentID = incident.ID
	c.DB.Where("id = ?", event.ID).Updates(&event)

	ctx.JSON(http.StatusCreated, gin.H{})
}

// Post
func (c *IncidencesController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var incident model.IncidentModel
	c.DB.Where("identifier = ?", id).First(&incident)
	if util.CheckNull(incident.Identifier) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "incident not found"})
		return
	}

	var body dto.IncidentDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	if !util.CheckNull(body.Content) {
		if strings.HasPrefix(body.Content, "http") &&
			strings.HasSuffix(body.Content, ".mdx") {
			req, err := http.Get(body.Content)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{})
				return
			}
			defer req.Body.Close()
			bytes, _ := io.ReadAll(req.Body)
			body.Content = string(bytes)
		}
	}
	if !util.CheckNull(body.EventIdentifier) {
		var oldEvent model.EventModel
		c.DB.Where("identifier = ?", incident.EventID).First(&oldEvent)

		var event model.EventModel
		c.DB.Where("identifier = ?", body.EventIdentifier).First(&event)
		if len(event.Identifier) > 0 {
			incident.EventID = event.ID
			event.IncidentID = incident.ID
			c.DB.Where("id = ?", event.ID).Updates(&event)
			oldEvent.IncidentID = 0
			c.DB.Updates(&oldEvent)
		}
	} else {
		var oldEvent model.EventModel
		c.DB.Where("identifier = ?", incident.EventID).First(&oldEvent)
		oldEvent.IncidentID = 0
		c.DB.Updates(&oldEvent)
	}
	c.DB.Where("id = ?", incident.ID).Updates(&incident)
}
