package use_cases

type SignInUseCase interface {
	Execute(name string) string
}

type signInUseCase struct{}

func NewSignInUseCase() SignInUseCase {
	return &signInUseCase{}
}

func (signInUseCase *signInUseCase) Execute(name string) string {
	return "Sign in successful " + name + "!"
}
