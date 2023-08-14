package extension

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOne[T any](
	collection *mongo.Collection,
	query *bson.M,
) (*T, error) {
	var item T

	err := collection.FindOne(context.Background(), query).Decode(&item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func Get[T any](
	collection *mongo.Collection,
	callback func(findOptions *options.FindOptions) bson.D,
) (*[]T, error) {
	var results []T

	findOptions := options.Find()
	filters := callback(findOptions)

	cursor, err := collection.Find(context.Background(), filters, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var item T
		_ = cursor.Decode(&item)
		results = append(results, item)
	}

	if results == nil {
		return &[]T{}, nil
	}

	return &results, nil
}

func InsertOne[T any](
	collection *mongo.Collection,
	item T,
) (*T, error) {
	_, err := collection.InsertOne(context.Background(), item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}
