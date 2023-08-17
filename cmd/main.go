package main

import (
	"my-story-time-api/application"
	"my-story-time-api/infrastructure"
	"my-story-time-api/presentation"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		application.DIContainer,
		presentation.DIContainer,
		infrastructure.DIContainer,
	)

	app.Run()
}
