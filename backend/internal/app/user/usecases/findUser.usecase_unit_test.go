package user_usecases

import (
	"errors"
	"testing"
	"time"

	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
	user_errors "github.com/luizrgf2/real-time-chat-go/internal/app/user/errors"
	user_interfaces_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/usecases"
	tests_repositories "github.com/luizrgf2/real-time-chat-go/tests/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var validUserToTestFindUserUseCase = user_entities.UserEntity{
	ID:        nil,
	UserName:  "testUser",
	Email:     "test@gmail.com",
	Password:  nil,
	CreatedAt: nil,
	UpdatedAt: nil,
}

func TestFindUserWithValidEmail(t *testing.T) {

	id := uint(1)
	now := time.Now()

	validUserToTestFindUserUseCase.ID = &id
	validUserToTestFindUserUseCase.CreatedAt = &now
	validUserToTestFindUserUseCase.UpdatedAt = &now

	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.On("FindByUserName", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)

	sut := FindUserUseCaseImp{
		UserRepository: userRepo,
	}

	res, err := sut.Exec(user_interfaces_usecases.FindUserUseCaseInput{
		Email: &validUserToTestFindUserUseCase.Email,
	})

	assert.ErrorIs(t, err, nil)
	assert.Equal(t, validUserToTestFindUserUseCase.UserName, res.User.UserName)
	assert.Equal(t, validUserToTestFindUserUseCase.Email, res.User.Email)

}

func TestFindUserWithValidUserName(t *testing.T) {

	id := uint(1)
	now := time.Now()

	validUserToTestFindUserUseCase.ID = &id
	validUserToTestFindUserUseCase.CreatedAt = &now
	validUserToTestFindUserUseCase.UpdatedAt = &now

	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.On("FindByUserName", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)

	sut := FindUserUseCaseImp{
		UserRepository: userRepo,
	}

	res, err := sut.Exec(user_interfaces_usecases.FindUserUseCaseInput{
		Username: &validUserToTestFindUserUseCase.UserName,
	})

	assert.ErrorIs(t, err, nil)
	assert.Equal(t, validUserToTestFindUserUseCase.UserName, res.User.UserName)
	assert.Equal(t, validUserToTestFindUserUseCase.Email, res.User.Email)

}

func TestFindUserWithValidId(t *testing.T) {

	id := uint(1)
	now := time.Now()

	validUserToTestFindUserUseCase.ID = &id
	validUserToTestFindUserUseCase.CreatedAt = &now
	validUserToTestFindUserUseCase.UpdatedAt = &now

	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.On("FindByUserName", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)

	sut := FindUserUseCaseImp{
		UserRepository: userRepo,
	}

	res, err := sut.Exec(user_interfaces_usecases.FindUserUseCaseInput{
		ID: validUserToTestFindUserUseCase.ID,
	})

	assert.ErrorIs(t, err, nil)
	assert.Equal(t, validUserToTestFindUserUseCase.ID, &id)
	assert.Equal(t, validUserToTestFindUserUseCase.UserName, res.User.UserName)
	assert.Equal(t, validUserToTestFindUserUseCase.Email, res.User.Email)

}

func TestFindUserWithEmptyFields(t *testing.T) {

	id := uint(1)
	now := time.Now()

	validUserToTestFindUserUseCase.ID = &id
	validUserToTestFindUserUseCase.CreatedAt = &now
	validUserToTestFindUserUseCase.UpdatedAt = &now

	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.On("FindByUserName", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)

	sut := FindUserUseCaseImp{
		UserRepository: userRepo,
	}

	_, err := sut.Exec(user_interfaces_usecases.FindUserUseCaseInput{})

	assert.ErrorIs(t, err, user_errors.ErrFieldsEmpties)
}

func TestFindUserWithEmailError(t *testing.T) {

	id := uint(1)
	now := time.Now()

	errExpected := errors.New("Erro para pegar usuário por email")

	validUserToTestFindUserUseCase.ID = &id
	validUserToTestFindUserUseCase.CreatedAt = &now
	validUserToTestFindUserUseCase.UpdatedAt = &now

	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(nil, errExpected)
	userRepo.On("FindByUserName", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)

	sut := FindUserUseCaseImp{
		UserRepository: userRepo,
	}

	_, err := sut.Exec(user_interfaces_usecases.FindUserUseCaseInput{
		Email: &validUserToTestFindUserUseCase.Email,
	})

	assert.ErrorIs(t, err, errExpected)
}

func TestFindUserWithUsernameError(t *testing.T) {

	id := uint(1)
	now := time.Now()

	errExpected := errors.New("Erro para pegar usuário por nome de usuário")

	validUserToTestFindUserUseCase.ID = &id
	validUserToTestFindUserUseCase.CreatedAt = &now
	validUserToTestFindUserUseCase.UpdatedAt = &now

	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.On("FindByUserName", mock.Anything, mock.Anything).Return(nil, errExpected)
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)

	sut := FindUserUseCaseImp{
		UserRepository: userRepo,
	}

	_, err := sut.Exec(user_interfaces_usecases.FindUserUseCaseInput{
		Username: &validUserToTestFindUserUseCase.UserName,
	})

	assert.ErrorIs(t, err, errExpected)
}

func TestFindUserWithIDError(t *testing.T) {

	id := uint(1)
	now := time.Now()

	errExpected := errors.New("Erro para pegar usuário por ID")

	validUserToTestFindUserUseCase.ID = &id
	validUserToTestFindUserUseCase.CreatedAt = &now
	validUserToTestFindUserUseCase.UpdatedAt = &now

	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.On("FindByUserName", mock.Anything, mock.Anything).Return(&validUserToTestFindUserUseCase, nil)
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, mock.Anything).Return(nil, errExpected)

	sut := FindUserUseCaseImp{
		UserRepository: userRepo,
	}

	_, err := sut.Exec(user_interfaces_usecases.FindUserUseCaseInput{
		ID: validUserToTestFindUserUseCase.ID,
	})

	assert.ErrorIs(t, err, errExpected)
}
