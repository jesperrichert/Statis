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
		// Badges
		// Polling

		notifications := api.Group("/notifications")
		{
			notifications.POST("/webhooks", middleware.DefaultMiddleware, c.NotificationsController.WebhookPost)
			notifications.POST("/:id", middleware.DefaultMiddleware, c.NotificationsController.UpdateWebhook)
		}

		status := api.Group("/status")
		{
			monitors := status.Group("/monitors")
			{
				monitors.GET("/", c.MonitorController.Get)
				monitors.POST("/", middleware.DefaultMiddleware, c.MonitorController.Post)
				monitors.POST("/:id", middleware.DefaultMiddleware, c.MonitorController.Update)
			}

			events := status.Group("/events")
			{
				events.GET("/", c.EventsController.Get)
				events.POST("/", middleware.DefaultMiddleware, c.EventsController.Post)
				events.POST("/:id", middleware.DefaultMiddleware, c.EventsController.Update)
			}

			incidences := status.Group("/incidences")
			{
				incidences.GET("/:id", c.IncidencesController.Get)
				incidences.POST("/", middleware.DefaultMiddleware, c.IncidencesController.Post)
				incidences.POST("/:id", middleware.DefaultMiddleware, c.IncidencesController.Update)
			}
		}
	}

}
