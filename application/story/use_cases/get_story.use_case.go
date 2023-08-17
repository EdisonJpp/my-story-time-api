package use_cases

import "my-story-time-api/domain/story"

type GetStoryUseCase interface {
	Execute(id string) (*story.Story, error)
}

type getStoryUseCase struct {
	storyRepository story.StoryRepository
}

func NewGetStoryUseCase(storyRepository story.StoryRepository) GetStoryUseCase {
	return &getStoryUseCase{storyRepository}
}

func (r *getStoryUseCase) Execute(id string) (*story.Story, error) {
	return r.storyRepository.GetById(id)
}
