package presentation

import (
	"my-story-time-api/presentation/http"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	http.DIContainer,
)
