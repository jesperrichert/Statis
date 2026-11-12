package main

import "jespersen.zip.statis/internal/config"

func main() {
	app := config.NewGin()
	db := config.NewDatabase()

	config.SetupFrontend(app)
	config.Build(&config.StatisConfig{
		Api: app,
		DB:  db,
	})

	err := app.Run(":3000")
	if err != nil {
		return
	}
}
