package use_cases

import "mytimes-api/domain/story"

type GetStoryUseCase interface {
	Execute(id string) (*story.Story, error)
}

type getStoryUseCase struct {
	storyRepository story.StoryRepository
}

func ProvideGetStoryUseCase(storyRepository story.StoryRepository) GetStoryUseCase {
	return &getStoryUseCase{storyRepository}
}

func (r *getStoryUseCase) Execute(id string) (*story.Story, error) {
	return r.storyRepository.GetById(id)
}
