package story

type StoryRepository interface {
	Get(userId string, page, pageSize int) (*[]Story, error)
	GetById(id string) (*Story, error)
	Create(story2 Story) (*Story, error)
}
