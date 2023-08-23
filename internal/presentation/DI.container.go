package presentation

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"my-story-time-api/internal/presentation/http"
)

var DIContainer = fx.Options(
	fx.Provide(validator.New),
	http.DIContainer,
)
