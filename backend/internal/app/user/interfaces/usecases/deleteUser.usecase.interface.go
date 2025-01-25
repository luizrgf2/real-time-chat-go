package user_interfaces_usecases

type DeleteUserUseCaseInput struct {
	ID uint
}

type IDeleteUserUseCase interface {
	Exec(input DeleteUserUseCaseInput) error
}
