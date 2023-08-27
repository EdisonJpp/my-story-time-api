package use_cases

import (
	"go.mongodb.org/mongo-driver/bson"
	"my-story-time-api/internal/domain/user"
)

type SignInUseCaseResponse struct {
	User        *user.User `json:"user"`
	AccessToken *string    `json:"accessToken"`
}

type SignInUseCase interface {
	Execute(isWithEmail bool, name string, password string) (*SignInUseCaseResponse, error)
}

type signInUseCase struct {
	userRepository user.UserRepository
}

func NewSignInUseCase(userRepository user.UserRepository) SignInUseCase {
	return &signInUseCase{userRepository}
}

func (signInUseCase *signInUseCase) Execute(
	isEmail bool,
	userName string,
	password string,
) (*SignInUseCaseResponse, error) {
	filter := bson.M{}

	if isEmail {
		filter["email"] = userName
	} else {
		filter["phoneNumber"] = userName
	}

	userValue, err := signInUseCase.userRepository.GetOneBy(&filter)

	if err != nil {
		return nil, err
	}

	if userValue.Password != password {
		return nil, user.ErrUserIncorrectInformation
	}

	accessToken := "accessToken"

	return &SignInUseCaseResponse{
		userValue,
		&accessToken,
	}, nil
}
