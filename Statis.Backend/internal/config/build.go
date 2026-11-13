package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/manager"
	"jespersen.zip.statis/internal/manager/monitor/checks"
	"jespersen.zip.statis/internal/schedule"
	"jespersen.zip.statis/internal/transport/http"
	"jespersen.zip.statis/internal/transport/http/router"
	"jespersen.zip.statis/internal/types/monitors"
)

type StatisConfig struct {
	Api *gin.Engine
	DB  *gorm.DB
}

func Build(config *StatisConfig) {

	monitorManager := manager.NewMonitorManager(map[monitors.MonitorType]checks.ICheckBase{
		monitors.Group: checks.NewCheckAdapter(checks.NewDummyCheck()),
		monitors.Http:  checks.NewCheckAdapter(checks.NewHttpCheck(config.DB)),
		monitors.Ping:  checks.NewCheckAdapter(checks.NewPingCheck(config.DB)),
		monitors.Tcp:   checks.NewCheckAdapter(checks.NewTCPCheck(config.DB)),
	})

	// Scheduler
	monitorScheduler := schedule.NewMonitorScheduler(config.DB, monitorManager)
	monitorScheduler.Setup()
	webhookScheduler := schedule.NewWebhookScheduler(config.DB)
	webhookScheduler.Setup()

	// Register Controller
	monitorController := http.NewMonitorController(config.DB)
	eventsController := http.NewEventsController(config.DB)
	incidencesController := http.NewIncidencesController(config.DB)
	notificationsController := http.NewNotificationsController(config.DB, webhookScheduler)

	routerConfig := router.RouterConfig{
		App:                     config.Api,
		MonitorController:       monitorController,
		EventsController:        eventsController,
		IncidencesController:    incidencesController,
		NotificationsController: notificationsController,
	}

	routerConfig.Setup()
}
