package config

import (
	db2 "my-story-time-api/internal/infrastructure/config/db"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(db2.NewMongoClient),
	fx.Provide(db2.NewDatabases),
	fx.Provide(db2.NewCollections),
)
