package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/transport/http"
	"jespersen.zip.statis/internal/transport/http/router"
)

type Appconfig struct {
	App *gin.Engine
	DB  *gorm.DB
}

func Build(config *Appconfig) {

	//Register Controller
	monitorController := http.NewMonitorController(config.DB)

	routeConfig := router.RouterConfig{
		App:               config.App,
		MonitorController: monitorController,
	}

	routeConfig.Setup()
}
