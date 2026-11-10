package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
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

	dtos := make([]dto.MonitorDto, 0)
	for i := range m {
		monitor := m[i]

		var effects []dto.MonitorEffect
		for k, v := range monitor.EffectedBy {
			effects = append(effects, dto.MonitorEffect{
				Name: k, Link: v,
			})
		}

		dtos = append(dtos, dto.MonitorDto{
			Identifier:  monitor.Identifier,
			IsParent:    monitor.IsParent,
			Group:       monitor.Group,
			Name:        monitor.Name,
			Description: monitor.Description,
			Type:        monitor.Type,
			ExternalUrl: monitor.ExternalUrl,
			CheckData:   monitor.CheckData,
			State:       monitor.State,
			EffectedBy:  effects,
		})
	}

	ctx.JSON(http.StatusOK, dtos)
}

// Post
func (c *MonitorController) Post(ctx *gin.Context) {

	var m dto.MonitorDto
	err := ctx.BindJSON(&m)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var monitor model.Monitor
	c.DB.Where(&model.Monitor{}).Where("identifier = ?", m.Identifier).First(&monitor)
	if len(monitor.Identifier) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Monitor already exists!"})
	}
	var effects map[string]string
	for i := range m.EffectedBy {
		effects[m.EffectedBy[i].Name] = m.EffectedBy[i].Link
	}

	c.DB.Create(&model.Monitor{
		Identifier:  m.Identifier,
		IsParent:    m.IsParent,
		Group:       m.Group,
		Name:        m.Name,
		Description: m.Description,
		Type:        m.Type,
		ExternalUrl: m.ExternalUrl,
		CheckData:   m.CheckData,
		State:       m.State,
		EffectedBy:  effects,
	})

	ctx.JSON(http.StatusOK, gin.H{"monitor": m})
}
