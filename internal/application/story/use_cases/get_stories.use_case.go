package use_cases

import (
	story2 "my-story-time-api/internal/domain/story"
)

type GetStoriesUseCase interface {
	Execute(userId string, page, pageSize int) (*[]story2.Story, error)
}

type getStoriesUseCase struct {
	storyRepository story2.StoryRepository
}

func NewGetStoriesUseCase(storyRepository story2.StoryRepository) GetStoriesUseCase {
	return &getStoriesUseCase{storyRepository}
}

func (r *getStoriesUseCase) Execute(userId string, page, pageSize int) (*[]story2.Story, error) {
	return r.storyRepository.Get(userId, page, pageSize)
}
