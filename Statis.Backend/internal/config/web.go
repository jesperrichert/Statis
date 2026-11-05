package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func NewWeb(engine *gin.Engine) {
	engine.Static("/assets", os.Getenv("FRONTEND_BUILD")+"/client/assets")
	engine.StaticFile("/statis.svg", os.Getenv("FRONTEND_BUILD")+"/client/statis.svg")
	engine.StaticFile("/online.webp", os.Getenv("FRONTEND_BUILD")+"/client/online.webp")
	engine.StaticFile("/offline.webp", os.Getenv("FRONTEND_BUILD")+"/client/offline.webp")
	engine.StaticFile("/degraded.webp", os.Getenv("FRONTEND_BUILD")+"/client/degraded.webp")
	engine.StaticFile("/", os.Getenv("FRONTEND_BUILD")+"/client/index.html")
	engine.NoRoute(func(c *gin.Context) {
		c.File(os.Getenv("FRONTEND_BUILD") + "/client/index.html")
	})
}
