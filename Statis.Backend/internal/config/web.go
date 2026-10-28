package config

import (
	"github.com/gin-gonic/gin"
	"os"
)

func NewWeb(engine *gin.Engine) {
	engine.Static("/assets", os.Getenv("FRONTEND_BUILD")+"/client/assets")
	engine.StaticFile("/statis.svg", os.Getenv("FRONTEND_BUILD")+"/client/statis.svg")
	engine.StaticFile("/", os.Getenv("FRONTEND_BUILD")+"/client/index.html")
	engine.NoRoute(func(c *gin.Context) {
		c.File(os.Getenv("FRONTEND_BUILD") + "/client/index.html")
	})
}