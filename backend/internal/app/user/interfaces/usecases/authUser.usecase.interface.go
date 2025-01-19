package user_interfaces_usecases

type AuthUserUseCaseInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserUseCaseOutput struct {
	Token  string `json:"token"`
	UserId uint   `json:"userId"`
}

type IAuthUserUseCase interface {
	Exec(input AuthUserUseCaseInput) (*AuthUserUseCaseOutput, error)
}
