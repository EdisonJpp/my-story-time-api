package use_cases

import (
	"mime/multipart"
	"my-story-time-api/domain/storage"
	"my-story-time-api/domain/story"
	"my-story-time-api/infrastructure/shared/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateStoryUseCaseRequest struct {
	Caption string
	IsAudio string
	Text    string
	UserId  string
	File    *multipart.FileHeader
}

type CreateStoryUseCase interface {
	Execute(request *CreateStoryUseCaseRequest) (*story.Story, error)
}

type createStoryUseCase struct {
	storyRepository   story.StoryRepository
	storageRepository storage.StorageRepository
}

func NewCreateStoryUseCase(storyRepository story.StoryRepository, storageRepository storage.StorageRepository) CreateStoryUseCase {
	return &createStoryUseCase{storyRepository, storageRepository}
}

func (r *createStoryUseCase) Execute(request *CreateStoryUseCaseRequest) (*story.Story, error) {
	var fileUrl string
	var fileSize int

	date, _ := utils.Date("")

	if request.IsAudio == "1" {
		file, err := r.storageRepository.Upload(
			request.File,
			request.UserId+"/stories/"+date.String()+"__"+request.File.Filename,
			"my-times-bucket",
		)

		if err != nil {
			return nil, err
		}

		fileUrl = file.Name
		fileSize = file.ChunkSize
	}

	return r.storyRepository.Create(story.Story{
		ID:      primitive.NewObjectID(),
		Caption: request.Caption,
		File: story.File{
			Url:  fileUrl,
			Size: fileSize,
		},
		UserId:    request.UserId,
		CreatedAt: date,
	})
}
