package db

import (
	"context"
	configDomain "my-story-time-api/internal/domain/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type DB struct {
	myStoryTimeDB *mongo.Database
}

func NewMongoClient(lc fx.Lifecycle, config *configDomain.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Mongo.Uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return client.Disconnect(context.Background())
		},
	})

	return client, nil
}
