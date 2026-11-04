package http

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
)

type MonitorController struct {
	db *gorm.DB
}

func NewMonitorController(db *gorm.DB) *MonitorController {
	return &MonitorController{
		db: db,
	}
}

// GET
func (e *MonitorController) Get(ctx *gin.Context) {

	data, ok := os.LookupEnv("MONITORS")
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

	var body []dto.MonitorDto
	json.NewDecoder(res.Body).Decode(&body)

	ctx.JSON(http.StatusOK, body)
}
