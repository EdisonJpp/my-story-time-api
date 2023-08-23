package story

import (
	"errors"
	story2 "my-story-time-api/internal/domain/story"
	"my-story-time-api/internal/infrastructure/config/db"
	"my-story-time-api/internal/infrastructure/shared/utils/extension"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type storyRepository struct {
	storyCollection *mongo.Collection
}

func NewStoryRepository(collections db.Collections) story2.StoryRepository {
	storyCollection := collections.Get("story")

	repo := storyRepository{storyCollection}

	return &repo
}

func (repo *storyRepository) Get(userId string, page, pageSize int) (*[]story2.Story, error) {
	filter := bson.D{{Key: "userId", Value: userId}}

	results, err := extension.Get[story2.Story](
		repo.storyCollection,
		func(findOptions *options.FindOptions) bson.D {
			findOptions.SetLimit(int64(pageSize))
			findOptions.SetSkip(int64((page - 1) * pageSize))

			return filter
		})

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *storyRepository) GetById(id string) (*story2.Story, error) {
	targetID, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": targetID}

	item, err := extension.GetOne[story2.Story](repo.storyCollection, &filter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, story2.ErrStoryNotFound
		}

		return nil, err
	}

	return item, nil
}

func (repo *storyRepository) Create(rawStory story2.Story) (*story2.Story, error) {
	inserted, err := extension.InsertOne[story2.Story](repo.storyCollection, rawStory)

	if err != nil {
		return nil, err
	}

	return inserted, nil
}
