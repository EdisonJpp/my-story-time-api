package auth

import (
	useCases "my-story-time-api/application/auth/use_cases"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(useCases.NewSignInUseCase),
)
