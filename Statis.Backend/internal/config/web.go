package config

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"jespersen.zip.statis/internal/util"
)

func SetupFrontend(engine *gin.Engine) {
	enabled, ok := util.Env("DISABLE_FRONTEND")
	if ok && enabled == "true" {
		engine.NoRoute(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusNotFound)
		})
		return
	}

	icon := os.Getenv("FRONTEND_BUILD" + "/client/favicon")
	if util.HasEnv("FAVICON") {
		icon = util.EnvString("FAVICON")
	}

	engine.Static("/assets", os.Getenv("FRONTEND_BUILD")+"/client/assets")
	engine.StaticFile("/favicon", icon)
	engine.StaticFile("/", os.Getenv("FRONTEND_BUILD")+"/client/index.html")
	engine.NoRoute(func(c *gin.Context) {
		c.File(os.Getenv("FRONTEND_BUILD") + "/client/index.html")
	})

	engine.GET("/config.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"favicon": "/favicon",
			"title":   util.EnvString("APP_TITLE"),
		})
	})
}
