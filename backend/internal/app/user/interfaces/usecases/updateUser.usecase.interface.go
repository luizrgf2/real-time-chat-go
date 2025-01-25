package user_interfaces_usecases

import user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"

type UpdateUserUseCaseInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserUseCaseOutput struct {
	User *user_entities.UserEntity `json:"user"`
}

type IUpdateUserUseCase interface {
	Exec(input UpdateUserUseCaseInput) (*UpdateUserUseCaseOutput, error)
}
