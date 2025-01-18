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

type CreateUserUseCaseImp struct {
	data               *user_interfaces_usecases.CreateUserUseCaseInput
	UserRepository     user_interfaces_repository.IUserRepository
	PassEncryptService user_interfaces_services.PassEncrypt
	userToSave         *user_entities.UserEntity
}

func (c *CreateUserUseCaseImp) checkIfEmailExists(ctx context.Context) error {
	_, err := c.UserRepository.FindByEmail(ctx, &c.data.Email)
	if err == nil {
		return user_errors.ErrEmailAlrearyExists
	}
	return nil
}

func (c *CreateUserUseCaseImp) checkIfExistsUserName(ctx context.Context) error {
	_, err := c.UserRepository.FindByUserName(ctx, &c.data.Username)
	if err == nil {
		return user_errors.ErrUserNameAlrearyExists
	}
	return nil
}

func (c *CreateUserUseCaseImp) encryptPassword() (*string, error) {
	result, err := c.PassEncryptService.EncryptPassword(&c.data.Password)
	if err != nil {
		return nil, err
	}
	return result.PasswordEncrypted, nil
}

func (c *CreateUserUseCaseImp) saveUser(ctx context.Context) error {
	err := c.UserRepository.Create(ctx, c.userToSave)
	return err
}

func (c *CreateUserUseCaseImp) Exec(input user_interfaces_usecases.CreateUserUseCaseInput) (*user_interfaces_usecases.CreateUserUseCaseOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c.data = &input

	user, err := user_entities.CreateWithoutId(input.Email, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	isExistsEmail := c.checkIfEmailExists(ctx)
	if isExistsEmail != nil {
		return nil, isExistsEmail
	}

	isExistsUserName := c.checkIfExistsUserName(ctx)
	if isExistsUserName != nil {
		return nil, isExistsUserName
	}

	encryptedPass, err := c.encryptPassword()
	if err != nil {
		return nil, err
	}

	user.Password = encryptedPass
	c.userToSave = user

	err = c.saveUser(ctx)
	if err != nil {
		return nil, err
	}

	user.Password = nil
	return &user_interfaces_usecases.CreateUserUseCaseOutput{
		ID:        *user.ID,
		Email:     user.UserName,
		Username:  user.UserName,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}, nil
}
