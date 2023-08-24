package config

import (
	db2 "my-story-time-api/internal/infrastructure/config/db"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(
		NewConfig,
		db2.NewMongoClient,
		db2.NewDatabases,
		db2.NewCollections,
	),
)
