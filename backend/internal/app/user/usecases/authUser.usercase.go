package user_usecases

import (
	"context"
	"time"

	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
	user_errors "github.com/luizrgf2/real-time-chat-go/internal/app/user/errors"
	user_interfaces_repository "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/repositories"
	user_interfaces_services "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/services"
	user_interfaces_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/usecases"
)

type AuthUserUseCaseImp struct {
	data               *user_interfaces_usecases.AuthUserUseCaseInput
	UserRepository     user_interfaces_repository.IUserRepository
	PassEncryptService user_interfaces_services.PassEncrypt
	JwtService         user_interfaces_services.JWTService
}

func (a *AuthUserUseCaseImp) findUserByEmail(ctx context.Context) (*user_entities.UserEntity, error) {
	user, err := a.UserRepository.FindByEmail(ctx, &a.data.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AuthUserUseCaseImp) passwordIsRight(encryptedPassword *string) error {
	isRight, err := a.PassEncryptService.ValidatePassword(encryptedPassword, &a.data.Password)
	if err != nil {
		return err
	}
	if !isRight {
		return user_errors.ErrAuthWrong
	}
	return nil
}

func (a *AuthUserUseCaseImp) createJwtToken(userId uint) (*string, error) {
	token, err := a.JwtService.Encode(userId)
	if err != nil {
		return nil, err
	}
	return token, err
}

func (a *AuthUserUseCaseImp) Exec(input user_interfaces_usecases.AuthUserUseCaseInput) (*user_interfaces_usecases.AuthUserUseCaseOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	a.data = &input
	user, err := a.findUserByEmail(ctx)
	if err != nil {
		return nil, err
	}

	err = a.passwordIsRight(user.Password)
	if err != nil {
		return nil, err
	}
	token, err := a.createJwtToken(*user.ID)
	if err != nil {
		return nil, err
	}

	return &user_interfaces_usecases.AuthUserUseCaseOutput{
		Token:  *token,
		UserId: *user.ID,
	}, nil
}
