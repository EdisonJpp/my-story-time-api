package main

import (
	"go.uber.org/fx"
	"my-story-time-api/internal/application"
	"my-story-time-api/internal/infrastructure"
	"my-story-time-api/internal/presentation"
)

func main() {
	app := fx.New(
		application.DIContainer,
		presentation.DIContainer,
		infrastructure.DIContainer,
	)

	app.Run()
}
