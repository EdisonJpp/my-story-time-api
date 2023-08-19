package presentation

import (
	"github.com/go-playground/validator/v10"
	"my-story-time-api/presentation/http"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(validator.New),
	http.DIContainer,
)
