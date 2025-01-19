package tests_repositories

import (
	"context"

	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
	RepositoryBaseMock[user_entities.UserEntity]
}

func (u *UserRepositoryMock) FindByEmail(ctx context.Context, email *string) (*user_entities.UserEntity, error) {
	args := u.Called(ctx, email)
	if userEntity, ok := args.Get(0).(*user_entities.UserEntity); ok {
		return userEntity, args.Error(1)
	}
	// Se o tipo não for o esperado, retorne um erro.
	return nil, args.Error(1)
}

func (u *UserRepositoryMock) FindByUserName(ctx context.Context, userName *string) (*user_entities.UserEntity, error) {
	args := u.Called(ctx, userName)
	if userEntity, ok := args.Get(0).(*user_entities.UserEntity); ok {
		return userEntity, args.Error(1)
	}
	return nil, args.Error(1)
}
