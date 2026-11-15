package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/schedule"
	"jespersen.zip.statis/internal/types/events"
	"jespersen.zip.statis/internal/types/monitors"
	"jespersen.zip.statis/internal/util"
)

type MonitorController struct {
	DB               *gorm.DB
	MonitorScheduler *schedule.MonitorScheduler
}

func NewMonitorController(db *gorm.DB, scheduler *schedule.MonitorScheduler) *MonitorController {
	return &MonitorController{
		DB:               db,
		MonitorScheduler: scheduler,
	}
}

// GET
func (c *MonitorController) Get(ctx *gin.Context) {
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
	if util.CheckNull(body.Identifier) ||
		util.CheckNull(body.IsParent) ||
		util.CheckNull(body.Name) ||
		util.CheckNull(body.ScheduleTime) ||
		util.CheckNull(body.Type) ||
		util.CheckNull(body.TypeData) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dto not valid"})
		return
	}
	if !monitors.ToType(body.Type).IsTypeDtoValid(body) ||
		body.TypeData == nil ||
		util.IsEmpty(body.TypeData) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid monitor type"})
		return
	}

	var monitorModel model.MonitorModel
	c.DB.Where(&model.MonitorModel{}).Where("identifier = ?", body.Identifier).First(&monitorModel)
	if len(monitorModel.Identifier) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Monitor already exists!"})
		return
	}
	var effects = map[string]string{}
	for i := range body.AffectedBy {
		effects[body.AffectedBy[i].Name] = body.AffectedBy[i].Link
	}

	c.DB.Create(&model.MonitorModel{
		Identifier:         body.Identifier,
		IsParent:           body.IsParent,
		Group:              body.Group,
		Name:               body.Name,
		Description:        body.Description,
		Type:               strings.ToUpper(body.Type),
		ExternalUrl:        body.ExternalUrl,
		TypeData:           body.TypeData,
		State:              body.State,
		AffectedByExternal: effects,
	})
	monitor := model.Monitor[any](monitorModel)
	c.MonitorScheduler.ScheduleMonitor(&monitor)

	ctx.JSON(http.StatusOK, gin.H{"monitor": body})
}

// Post
func (c *MonitorController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var monitorModel model.MonitorModel
	c.DB.Where("identifier = ?", id).First(&monitorModel)
	if util.CheckNull(monitorModel.Identifier) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "monitor not found"})
		return
	}
	var body dto.MonitorDto[any]
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	if !util.CheckNull(body.Type) && util.CheckNull(body.TypeData) {
		if monitors.ToType(body.Type).IsTypeDtoValid(body) ||
			body.TypeData != nil ||
			!util.IsEmpty(body.TypeData) {
			monitorModel.Type = strings.ToUpper(body.Type)
			monitorModel.TypeData = body.TypeData
		}
	}
	if !util.CheckNull(body.Name) {
		monitorModel.Name = body.Name
	}
	if !util.CheckNull(body.Description) {
		monitorModel.Description = body.Description
	}
	if !util.CheckNull(body.AffectedBy) {
		effects := map[string]string{}
		for i := range body.AffectedBy {
			effect := body.AffectedBy[i]
			effects[effect.Name] = effect.Link
		}
		monitorModel.AffectedByExternal = effects
	}
	if !util.CheckNull(body.ScheduleTime) {
		monitorModel.ScheduleTime = body.ScheduleTime
	}
	if !util.CheckNull(body.ExternalUrl) {
		monitorModel.ExternalUrl = body.ExternalUrl
	}
	if !util.CheckNull(body.IsParent) {
		monitorModel.IsParent = body.IsParent
	}
	if !util.CheckNull(body.Group) {
		monitorModel.Group = body.Group
	}

	key := string(rune(monitorModel.ID)) + monitorModel.Identifier
	if c.MonitorScheduler.Jobs[key] != nil {
		_ = c.MonitorScheduler.Scheduler.RemoveJob(c.MonitorScheduler.Jobs[key].ID())
	}
	monitor := model.Monitor[any](monitorModel)
	c.MonitorScheduler.ScheduleMonitor(&monitor)

	c.DB.Where("id = ?", monitorModel.ID).Updates(&monitorModel)
}
