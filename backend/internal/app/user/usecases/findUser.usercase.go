package user_usecases

import (
	"context"
	"time"

	user_errors "github.com/luizrgf2/real-time-chat-go/internal/app/user/errors"
	user_interfaces_repository "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/repositories"
	user_interfaces_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/usecases"
)

type FindUserUseCaseImp struct {
	data           *user_interfaces_usecases.FindUserUseCaseInput
	UserRepository user_interfaces_repository.IUserRepository
}

func (f *FindUserUseCaseImp) validateFields() error {
	if f.data.Email == nil && f.data.ID == nil && f.data.Username == nil {
		return user_errors.ErrFieldsEmpties
	}
	return nil
}

func (f *FindUserUseCaseImp) Exec(input user_interfaces_usecases.FindUserUseCaseInput) (*user_interfaces_usecases.FindUserUseCaseOutput, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	f.data = &input

	err := f.validateFields()
	if err != nil {
		return nil, err
	}

	if f.data.Email != nil {
		user, err := f.UserRepository.FindByEmail(ctx, f.data.Email)
		if err != nil {
			return nil, err
		}
		return &user_interfaces_usecases.FindUserUseCaseOutput{User: user}, nil

	} else if f.data.Username != nil {
		user, err := f.UserRepository.FindByUserName(ctx, f.data.Username)
		if err != nil {
			return nil, err
		}
		return &user_interfaces_usecases.FindUserUseCaseOutput{User: user}, nil

	} else if f.data.ID != nil {
		user, err := f.UserRepository.FindByID(ctx, *input.ID)
		if err != nil {
			return nil, err
		}
		return &user_interfaces_usecases.FindUserUseCaseOutput{User: user}, nil
	} else {
		return nil, nil
	}
}
