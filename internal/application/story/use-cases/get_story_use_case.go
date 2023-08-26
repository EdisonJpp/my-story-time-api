package use_cases

import (
	story2 "my-story-time-api/internal/domain/story"
)

type GetStoryUseCase interface {
	Execute(id string) (*story2.Story, error)
}

type getStoryUseCase struct {
	storyRepository story2.StoryRepository
}

func NewGetStoryUseCase(storyRepository story2.StoryRepository) GetStoryUseCase {
	return &getStoryUseCase{storyRepository}
}

func (r *getStoryUseCase) Execute(id string) (*story2.Story, error) {
	return r.storyRepository.GetById(id)
}
