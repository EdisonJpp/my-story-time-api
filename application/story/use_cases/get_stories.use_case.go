package use_cases

import "my-story-time-api/domain/story"

type GetStoriesUseCase interface {
	Execute(userId string, page, pageSize int) (*[]story.Story, error)
}

type getStoriesUseCase struct {
	storyRepository story.StoryRepository
}

func NewGetStoriesUseCase(storyRepository story.StoryRepository) GetStoriesUseCase {
	return &getStoriesUseCase{storyRepository}
}

func (r *getStoriesUseCase) Execute(userId string, page, pageSize int) (*[]story.Story, error) {
	return r.storyRepository.Get(userId, page, pageSize)
}
