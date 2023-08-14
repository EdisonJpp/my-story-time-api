package main

import (
	"mytimes-api/application"
	"mytimes-api/infrastructure"
	"mytimes-api/presentation"

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
