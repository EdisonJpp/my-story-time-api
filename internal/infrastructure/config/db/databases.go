package db

import "go.mongodb.org/mongo-driver/mongo"

type Databases struct {
	myStoryTimeDB *mongo.Database
}

func NewDatabases(client *mongo.Client) *Databases {
	myStoryTimeDB := client.Database("my_times")

	return &Databases{
		myStoryTimeDB,
	}
}
