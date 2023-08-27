package user

import "go.uber.org/fx"

var DIContainer = fx.Options(
	fx.Provide(NewUserRepository),
)
