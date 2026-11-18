package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/services"
	"jespersen.zip.statis/internal/types/events"
	"jespersen.zip.statis/internal/types/monitors"
	"jespersen.zip.statis/internal/util"
)

type MonitorController struct {
	DB             *gorm.DB
	MonitorService *services.MonitorService
}

func NewMonitorController(db *gorm.DB, service *services.MonitorService) *MonitorController {
	return &MonitorController{
		DB:             db,
		MonitorService: service,
	}
}

// GET
func (c *MonitorController) Discover(ctx *gin.Context) {
	var m []model.MonitorModel
	c.DB.Preload("Events").Find(&m)

	dtos := make([]dto.MonitorDto[any], 0)
	for i := range m {
		monitor := m[i]

		var effects = make([]dto.MonitorEffect, 0)
		for k, v := range monitor.AffectedByExternal {
			effects = append(effects, dto.MonitorEffect{
				Name: k, Link: v,
			})
		}
		for i := range monitor.Events {
			event := monitor.Events[i]
			if event.Type == events.Resolved.ToString() {
				continue
			}
			url := "#"
			if event.IncidentID != 0 {
				var incident model.IncidentModel
				c.DB.Where("id = ?", event.IncidentID).First(&incident)
				url = "/incidences/" + incident.Identifier
			}
			effects = append(effects, dto.MonitorEffect{
				Name: event.Title,
				Link: url,
			})
		}

		state := monitor.State
		if len(effects) > 0 {
			state = monitors.Affected.ToString()
		}
		dtos = append(dtos, dto.MonitorDto[any]{
			Identifier:  monitor.Identifier,
			IsParent:    monitor.IsParent,
			Group:       monitor.Group,
			Name:        monitor.Name,
			Description: monitor.Description,
			Type:        strings.ToUpper(monitor.Type),
			ExternalUrl: monitor.ExternalUrl,
			State:       state,
			AffectedBy:  effects,
		})
	}

	ctx.JSON(http.StatusOK, dtos)
}

// Post
func (c *MonitorController) Post(ctx *gin.Context) {

	var body dto.MonitorDto[any]
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, message := c.MonitorService.Validate(&body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}
	status = c.MonitorService.Add(&body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult("failed to create", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// Get
func (c *MonitorController) Get(ctx *gin.Context) {
	monitorDto, ok := c.MonitorService.Get(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, util.JsonResult("monitor not found", map[string]any{}))
		return
	}
	ctx.JSON(http.StatusOK, monitorDto)
}

// Post
func (c *MonitorController) Update(ctx *gin.Context) {

	var body dto.MonitorDto[any]
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	status, message := c.MonitorService.Validate(&body)
	if status != http.StatusOK && status != http.StatusConflict {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}
	status = c.MonitorService.Update(ctx.Param("id"), &body)
	if status != http.StatusCreated {
		ctx.JSON(status, util.JsonResult("failed to update", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// Delete
func (c *MonitorController) Delete(ctx *gin.Context) {
	ok := c.MonitorService.Delete(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, util.JsonResult("monitor not found", map[string]any{}))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
