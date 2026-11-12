package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/types/monitors"
	"jespersen.zip.statis/internal/util"
)

type MonitorController struct {
	DB *gorm.DB
}

func NewMonitorController(db *gorm.DB) *MonitorController {
	return &MonitorController{
		DB: db,
	}
}

// GET
func (c *MonitorController) Get(ctx *gin.Context) {
	var m []model.Monitor
	c.DB.Find(&m)

	dtos := make([]dto.MonitorDto[any], 0)
	for i := range m {
		monitor := m[i]

		var effects []dto.MonitorEffect
		for k, v := range monitor.EffectedBy {
			effects = append(effects, dto.MonitorEffect{
				Name: k, Link: v,
			})
		}

		dtos = append(dtos, dto.MonitorDto[any]{
			Identifier:  monitor.Identifier,
			IsParent:    monitor.IsParent,
			Group:       monitor.Group,
			Name:        monitor.Name,
			Description: monitor.Description,
			Type:        monitor.Type,
			ExternalUrl: monitor.ExternalUrl,
			TypeData:    monitor.TypeData,
			State:       monitor.State,
			EffectedBy:  effects,
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
		util.CheckNull(body.State) ||
		util.CheckNull(body.IsParent) ||
		util.CheckNull(body.Name) ||
		util.CheckNull(body.ScheduleTime) ||
		util.CheckNull(body.Type) ||
		util.CheckNull(body.TypeData) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dto not valid})"})
		return
	}
	if !monitors.ToType(body.Type).IsTypeDtoValid(body) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid monitor type"})
		return
	}

	var monitor model.Monitor
	c.DB.Where(&model.Monitor{}).Where("identifier = ?", body.Identifier).First(&monitor)
	if len(monitor.Identifier) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Monitor already exists!"})
		return
	}
	var effects = map[string]string{}
	for i := range body.EffectedBy {
		effects[body.EffectedBy[i].Name] = body.EffectedBy[i].Link
	}

	c.DB.Create(&model.Monitor{
		Identifier:  body.Identifier,
		IsParent:    body.IsParent,
		Group:       body.Group,
		Name:        body.Name,
		Description: body.Description,
		Type:        body.Type,
		ExternalUrl: body.ExternalUrl,
		TypeData:    body.TypeData,
		State:       body.State,
		EffectedBy:  effects,
	})

	ctx.JSON(http.StatusOK, gin.H{"monitor": body})
}
