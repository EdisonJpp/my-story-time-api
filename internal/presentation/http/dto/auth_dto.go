package dto

type SignInRequestDto struct {
	UserName string `validate:"required,min=5,max=70"`
	Password string `validate:"required,min=5,max=170"`
	IsEmail  bool   `validate:"omitempty"`
}
