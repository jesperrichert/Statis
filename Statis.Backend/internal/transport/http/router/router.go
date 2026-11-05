package router

import (
	"github.com/gin-gonic/gin"
	"jespersen.zip.statis/internal/transport/http"
)

type RouterConfig struct {
	App                  *gin.Engine
	MonitorController    *http.MonitorController
	EventsController     *http.EventsController
	IncidencesController *http.IncidencesController
}

func (c *RouterConfig) Setup() {
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
			}

			events := status.Group("/events")
			{
				events.GET("/", c.EventsController.Get)
			}

			incidences := status.Group("incidences")
			{
				incidences.GET("/:id", c.IncidencesController.Get)
			}
		}
	}

}
