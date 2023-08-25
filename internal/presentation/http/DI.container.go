package http

import (
	controllers2 "my-story-time-api/internal/presentation/http/controllers"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(NewHttp),
	fx.Invoke(controllers2.NewAuthController, controllers2.NewStoryController),
)
