package user_usecases_test

import (
	"errors"
	"testing"

	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
	user_interfaces_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/usecases"
	user_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/usecases"
	tests_repositories "github.com/luizrgf2/real-time-chat-go/tests/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteUser_Success(t *testing.T) {
	// Mocks
	userRepo := new(tests_repositories.UserRepositoryMock)
	validUserID := uint(1)

	// Dados simulados
	userToDelete := user_entities.UserEntity{
		ID:        &validUserID,
		UserName:  "testuser",
		Email:     "test@email.com",
		Password:  nil,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, validUserID).Return(&userToDelete, nil).Once()
	userRepo.RepositoryBaseMock.On("Delete", mock.Anything, validUserID).Return(nil).Once()

	sut := user_usecases.DeleteUserUseCaseImp{
		UserRepository: userRepo,
	}

	err := sut.Exec(user_interfaces_usecases.DeleteUserUseCaseInput{
		ID: validUserID,
	})

	// Validações
	assert.NoError(t, err)
	userRepo.RepositoryBaseMock.AssertCalled(t, "FindByID", mock.Anything, validUserID)
	userRepo.RepositoryBaseMock.AssertCalled(t, "Delete", mock.Anything, validUserID)
}

func TestDeleteUser_NotFound(t *testing.T) {
	userRepo := new(tests_repositories.UserRepositoryMock)
	invalidUserID := uint(999)

	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, invalidUserID).Return(nil, errors.New("user not found")).Once()

	sut := user_usecases.DeleteUserUseCaseImp{
		UserRepository: userRepo,
	}

	err := sut.Exec(user_interfaces_usecases.DeleteUserUseCaseInput{
		ID: invalidUserID,
	})

	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
	userRepo.RepositoryBaseMock.AssertCalled(t, "FindByID", mock.Anything, invalidUserID)
	userRepo.RepositoryBaseMock.AssertNotCalled(t, "Delete", mock.Anything, invalidUserID)
}

func TestDeleteUser_RepositoryDeleteError(t *testing.T) {
	// Mocks
	userRepo := new(tests_repositories.UserRepositoryMock)
	validUserID := uint(1)

	// Dados simulados
	userToDelete := user_entities.UserEntity{
		ID:        &validUserID,
		UserName:  "testuser",
		Email:     "test@email.com",
		Password:  nil,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	// Mock do repositório
	userRepo.RepositoryBaseMock.On("FindByID", mock.Anything, validUserID).Return(&userToDelete, nil).Once()
	userRepo.RepositoryBaseMock.On("Delete", mock.Anything, validUserID).Return(errors.New("delete operation failed")).Once()

	// Instância do caso de uso
	sut := user_usecases.DeleteUserUseCaseImp{
		UserRepository: userRepo,
	}

	// Execução
	err := sut.Exec(user_interfaces_usecases.DeleteUserUseCaseInput{
		ID: validUserID,
	})

	// Validações
	assert.Error(t, err)
	assert.EqualError(t, err, "delete operation failed")
	userRepo.RepositoryBaseMock.AssertCalled(t, "FindByID", mock.Anything, validUserID)
	userRepo.RepositoryBaseMock.AssertCalled(t, "Delete", mock.Anything, validUserID)
}
