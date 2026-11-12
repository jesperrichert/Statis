package router

import (
	"github.com/gin-gonic/gin"
	"jespersen.zip.statis/internal/transport/http"
	"jespersen.zip.statis/internal/transport/http/middleware"
)

type RouterConfig struct {
	App                     *gin.Engine
	MonitorController       *http.MonitorController
	EventsController        *http.EventsController
	IncidencesController    *http.IncidencesController
	NotificationsController *http.NotificationsController
}

func (c *RouterConfig) Setup() {
	if c.App == nil {
		c.App = gin.Default()
	}

	api := c.App.Group("/api")
	{
		notifications := api.Group("/notifications")
		{
			notifications.POST("/webhooks", middleware.DefaultMiddleware, c.NotificationsController.WebhookPost)
		}

		status := api.Group("/status")
		{
			monitors := status.Group("/monitors")
			{
				monitors.GET("/", c.MonitorController.Get)
				monitors.POST("/", middleware.DefaultMiddleware, c.MonitorController.Post)
			}

			events := status.Group("/events")
			{
				events.GET("/", c.EventsController.Get)
				events.POST("/", middleware.DefaultMiddleware, c.EventsController.Get)
			}

			incidences := status.Group("/incidences")
			{
				incidences.GET("/:id", c.IncidencesController.Get)
				incidences.POST("/", middleware.DefaultMiddleware, c.IncidencesController.Post)
			}
		}
	}

}

func (c *RouterConfig) Post() {
	if c.App == nil {
		c.App = gin.Default()
	}

	api := c.App.Group("/api")
	{
		status := api.Group("/status")
		{
			monitors := status.Group("/monitors")
			{
				monitors.GET("/", c.MonitorController.Get)
				monitors.POST("/", middleware.DefaultMiddleware, c.MonitorController.Post)
			}

			events := status.Group("/events")
			{
				events.GET("/", c.EventsController.Get)
				events.POST("/", middleware.DefaultMiddleware, c.EventsController.Get)
			}

			incidences := status.Group("/incidences")
			{
				incidences.GET("/:id", c.IncidencesController.Get)
				incidences.POST("/", middleware.DefaultMiddleware, c.IncidencesController.Post)
			}
		}
	}

}
