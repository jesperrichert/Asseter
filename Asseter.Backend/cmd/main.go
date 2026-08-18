package main

import (
	"go.Asseter/internal/config"
	"go.Asseter/internal/env"
)

func main() {
	app := config.NewGin()
	db := config.NewDatabase()
	config.NewWeb(app)
	env := env.NewConfig()
	config.Build(&config.Appconfig{
		App:    app,
		DB:     db,
		Config: env,
	})

	err := app.Run(":3000")
	if err != nil {
		return
	}
}
