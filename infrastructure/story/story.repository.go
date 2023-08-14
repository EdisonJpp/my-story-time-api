package story

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mytimes-api/domain/story"
	"mytimes-api/infrastructure/config/db"
	"mytimes-api/infrastructure/shared/utils/extension"
)

type storyRepository struct {
	storyCollection *mongo.Collection
}

func ProvideStoryRepository(collections db.Collections) story.StoryRepository {
	storyCollection := collections.Get("story")

	repo := storyRepository{storyCollection}

	return &repo
}

func (repo *storyRepository) Get(userId string, page, pageSize int) (*[]story.Story, error) {
	filter := bson.D{{Key: "userId", Value: userId}}

	results, err := extension.Get[story.Story](
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

func (repo *storyRepository) GetById(id string) (*story.Story, error) {
	targetID, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": targetID}

	item, err := extension.GetOne[story.Story](repo.storyCollection, &filter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, story.ErrStoryNotFound
		}

		return nil, err
	}

	return item, nil
}

func (repo *storyRepository) Create(rawStory story.Story) (*story.Story, error) {
	inserted, err := extension.InsertOne[story.Story](repo.storyCollection, rawStory)

	if err != nil {
		return nil, err
	}

	return inserted, nil
}
