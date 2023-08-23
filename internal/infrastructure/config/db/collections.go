package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Collections interface {
	Get(collectionName string) *mongo.Collection
}

type collections struct {
	story *mongo.Collection
	user  *mongo.Collection
}

func NewCollections(databases *Databases) Collections {
	story := newStoryCollection(databases)
	user := newUserCollection(databases)

	return &collections{story, user}
}

func (c *collections) Get(name string) *mongo.Collection {
	switch name {
	case "story":
		return c.story
	case "user":
		return c.user
	default:
		panic("This collection does not exist: " + name)
	}
}

func newStoryCollection(databases *Databases) *mongo.Collection {
	storiesCollection := databases.myStoryTimeDB.Collection("c_stories")

	opt := options.Index()

	storiesIndexes := mongo.IndexModel{Keys: bson.M{"userId": 1}, Options: opt}

	_, err := storiesCollection.Indexes().CreateOne(context.Background(), storiesIndexes)

	if err != nil {
		log.Println("Could not create index:", err)
	}

	return storiesCollection
}

func newUserCollection(databases *Databases) *mongo.Collection {
	storiesCollection := databases.myStoryTimeDB.Collection("c_users")

	opt := options.Index()
	opt.SetUnique(true)

	storiesIndexes := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	_, err := storiesCollection.Indexes().CreateOne(context.Background(), storiesIndexes)

	if err != nil {
		log.Println("Could not create index:", err)
	}

	return storiesCollection
}
