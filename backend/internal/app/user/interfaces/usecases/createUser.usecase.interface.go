package interfaces_usecases

type CreateUserUseCaseInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserUseCaseOutput struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type ICreateUserUseCase interface {
	Exec(input CreateUserUseCaseInput) (*CreateUserUseCaseOutput, error)
}
