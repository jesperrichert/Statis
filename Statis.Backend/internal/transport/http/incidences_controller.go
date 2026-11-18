package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
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
	incidentDto, ok := c.IncidentService.Get(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "incidence not found"})
		return
	}
	ctx.JSON(http.StatusOK, incidentDto)
}

// GET
func (c *IncidencesController) RawContent(ctx *gin.Context) {

	incidentDto, ok := c.IncidentService.Get(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "incidence not found"})
		return
	}

	ctx.Data(http.StatusOK, "text/plain", []byte(incidentDto.Content))
}

// POST
func (c *IncidencesController) Post(ctx *gin.Context) {

	var body dto.IncidentDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.JsonResult("failed to to parse dto", map[string]any{}))
		return
	}

	status, message := c.IncidentService.Validate(&body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}
	status = c.IncidentService.Add(&body)
	if status != http.StatusCreated {
		ctx.JSON(status, util.JsonResult("failed to create", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

// Post
func (c *IncidencesController) Update(ctx *gin.Context) {
	var body dto.IncidentDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	status, message := c.IncidentService.Validate(&body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}
	status = c.IncidentService.Update(ctx.Param("id"), &body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult("failed to update dto", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// GET
func (c *IncidencesController) Delete(ctx *gin.Context) {

	ok := c.IncidentService.Delete(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "incidence not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
