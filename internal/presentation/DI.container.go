package presentation

import (
	"my-story-time-api/internal/presentation/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(validator.New),
	http.DIContainer,
)
