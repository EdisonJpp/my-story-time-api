package user

import (
	"errors"
	userDomain "my-story-time-api/internal/domain/user"
	"my-story-time-api/internal/infrastructure/config/db"
	"my-story-time-api/internal/infrastructure/shared/utils/extension"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(collections db.Collections) userDomain.UserRepository {
	userCollection := collections.Get("user")

	repo := userRepository{userCollection}

	return &repo
}

func (repo *userRepository) GetOneBy(filter *bson.M) (*userDomain.User, error) {
	item, err := extension.GetOne[userDomain.User](repo.userCollection, filter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, userDomain.ErrUserNotFound
		}

		return nil, err
	}

	return item, nil
}
