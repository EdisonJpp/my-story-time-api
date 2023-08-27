package dto

type GetStoryRequestDto struct {
	ID string `validate:"required,min=5,max=70"`
}
