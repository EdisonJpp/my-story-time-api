package presentation

import (
	"go.uber.org/fx"
	"mytimes-api/presentation/http"
)

var DIContainer = fx.Options(
	http.DIContainer,
)
