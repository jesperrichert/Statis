package config

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupFrontend(engine *gin.Engine) {
	enabled, ok := os.LookupEnv("DISABLE_FRONTEND")
	if ok && enabled == "true" {
		engine.NoRoute(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusNotFound)
		})
		return
	}
	engine.Static("/assets", os.Getenv("FRONTEND_BUILD")+"/client/assets")
	engine.StaticFile("/statis.svg", os.Getenv("FRONTEND_BUILD")+"/client/statis.svg")
	engine.StaticFile("/", os.Getenv("FRONTEND_BUILD")+"/client/index.html")
	engine.NoRoute(func(c *gin.Context) {
		c.File(os.Getenv("FRONTEND_BUILD") + "/client/index.html")
	})
}
