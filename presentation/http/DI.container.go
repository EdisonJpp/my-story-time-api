package http

import (
	"my-story-time-api/presentation/http/controllers"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(NewHttp),
	fx.Invoke(controllers.NewAuthController),
	fx.Invoke(controllers.NewStoryController),
)
