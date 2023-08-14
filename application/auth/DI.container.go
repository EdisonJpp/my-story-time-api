package auth

import (
	"go.uber.org/fx"
	useCases "mytimes-api/application/auth/use_cases"
)

var DIContainer = fx.Options(
	fx.Provide(useCases.ProvideSignInUseCase),
)
