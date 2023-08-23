package http

import (
	"go.uber.org/fx"
	controllers2 "my-story-time-api/internal/presentation/http/controllers"
)

var DIContainer = fx.Options(
	fx.Provide(NewHttp),
	fx.Invoke(controllers2.NewAuthController),
	fx.Invoke(controllers2.NewStoryController),
)
