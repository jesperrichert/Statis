package config

import (
	"github.com/gin-gonic/gin"
	"jespersen.zip.statis/internal/repository"
	"jespersen.zip.statis/internal/services"
	"jespersen.zip.statis/internal/transport/http"
	"jespersen.zip.statis/internal/transport/http/router"
	"gorm.io/gorm"
)

type Appconfig struct {
	App *gin.Engine
	DB  *gorm.DB
}

func Build(config *Appconfig) {
	//Register Repositories
	exampleRepository := repository.NewExample()

	//Register Services
	exampleService := services.NewExampleService()

	//Register Controller
	exampleController := http.NewExampleController(config.DB, exampleService, exampleRepository) 

	routeConfig := router.RouterConfig{
		App: config.App, 
		ExampleController: exampleController, 
	} 

	routeConfig.Setup()
}
