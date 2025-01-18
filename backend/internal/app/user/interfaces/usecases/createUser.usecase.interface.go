package user_interfaces_usecases

import "time"

type CreateUserUseCaseInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserUseCaseOutput struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type ICreateUserUseCase interface {
	Exec(input CreateUserUseCaseInput) (*CreateUserUseCaseOutput, error)
}
