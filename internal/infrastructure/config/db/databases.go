package db

import (
	"my-story-time-api/internal/domain/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Databases struct {
	myStoryTimeDB *mongo.Database
}

func NewDatabases(client *mongo.Client, config *config.Config) *Databases {
	myStoryTimeDB := client.Database(config.Mongo.Db)

	return &Databases{
		myStoryTimeDB,
	}
}
