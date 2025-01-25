package user_interfaces_usecases

import (
	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
)

type FindUserUseCaseInput struct {
	ID       *uint   `json:"id"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
}

type FindUserUseCaseOutput struct {
	User *user_entities.UserEntity `json:"user"`
}

type IFindUserUseCase interface {
	Exec(input FindUserUseCaseInput) (*FindUserUseCaseOutput, error)
}
