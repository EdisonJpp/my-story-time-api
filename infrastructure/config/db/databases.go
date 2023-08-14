package db

import "go.mongodb.org/mongo-driver/mongo"

type Databases struct {
	myTimesDB *mongo.Database
}

func ProvideDatabases(client *mongo.Client) *Databases {
	myTimesDB := client.Database("my_times")

	return &Databases{
		myTimesDB,
	}
}
