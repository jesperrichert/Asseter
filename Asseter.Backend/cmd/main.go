package main

import "go.Asseter/internal/config"

func main() {
	app := config.NewGin()
	db := config.NewDatabase()
	config.NewWeb(app)

	config.Build(&config.Appconfig{
		App: app,
		DB:  db,
	})

	err := app.Run(":3000")
	if err != nil {
		return
	}
}
