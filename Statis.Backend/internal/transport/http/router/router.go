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

			notifications.GET("/webhooks/:id", middleware.DefaultMiddleware, c.NotificationsController.GetWebhook)
			notifications.POST("/webhooks/:id", middleware.DefaultMiddleware, c.NotificationsController.UpdateWebhook)
			notifications.DELETE("/webhooks/:id", middleware.DefaultMiddleware, c.NotificationsController.DeleteWebhook)
		}

		status := api.Group("/status")
		{
			monitors := status.Group("/monitors")
			{
				monitors.GET("/", c.MonitorController.Discover)
				monitors.POST("/", middleware.DefaultMiddleware, c.MonitorController.Post)

				monitors.GET("/:id", c.MonitorController.Get)
				monitors.POST("/:id", middleware.DefaultMiddleware, c.MonitorController.Update)
				monitors.DELETE("/:id", middleware.DefaultMiddleware, c.MonitorController.Delete)
			}

			events := status.Group("/events")
			{
				events.GET("/", c.EventsController.Discover)
				events.POST("/", middleware.DefaultMiddleware, c.EventsController.Post)

				events.GET("/:id", c.EventsController.Get)
				events.POST("/:id", middleware.DefaultMiddleware, c.EventsController.Update)
				events.DELETE("/:id", middleware.DefaultMiddleware, c.EventsController.Delete)
			}

			incidences := status.Group("/incidences")
			{
				incidences.POST("/", middleware.DefaultMiddleware, c.IncidencesController.Post)

				incidences.GET("/:id/", c.IncidencesController.Get)
				incidences.GET("/:id/raw", c.IncidencesController.RawContent)
				incidences.POST("/:id", middleware.DefaultMiddleware, c.IncidencesController.Update)
				incidences.DELETE("/:id", middleware.DefaultMiddleware, c.IncidencesController.Delete)
			}
		}
	}

}
