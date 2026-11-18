package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/manager"
	"jespersen.zip.statis/internal/manager/monitor/checks"
	"jespersen.zip.statis/internal/schedule"
	auto_config "jespersen.zip.statis/internal/schedule/config"
	"jespersen.zip.statis/internal/services"
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
	monitorScheduler := schedule.NewMonitorScheduler(config.DB, monitorManager).Setup()
	webhookScheduler := schedule.NewWebhookScheduler(config.DB).Setup()

	// Service
	webhookService := services.NewWebhookService(config.DB, webhookScheduler)
	incidentService := services.NewIncidentService(config.DB)
	monitorService := services.NewMonitorService(config.DB, monitorScheduler)
	eventService := services.NewEventService(config.DB)

	auto_config.NewAutoConfigureScheduler(
		config.DB,
		webhookService,
		monitorService,
		incidentService,
		eventService,
	).Setup()

	// Register Controller
	monitorController := http.NewMonitorController(config.DB, monitorService)
	eventsController := http.NewEventsController(config.DB, eventService)
	incidencesController := http.NewIncidencesController(config.DB, incidentService)
	notificationsController := http.NewNotificationsController(config.DB, webhookService)

	routerConfig := router.RouterConfig{
		App:                     config.Api,
		MonitorController:       monitorController,
		EventsController:        eventsController,
		IncidencesController:    incidencesController,
		NotificationsController: notificationsController,
	}

	routerConfig.Setup()
}
