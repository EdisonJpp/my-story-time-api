package http

import (
	"go.uber.org/fx"
	"mytimes-api/presentation/http/controllers"
)

var DIContainer = fx.Options(
	fx.Provide(ProvideHttp),
	fx.Invoke(controllers.ProvideAuthController),
	fx.Invoke(controllers.ProvideStoryController),
)
