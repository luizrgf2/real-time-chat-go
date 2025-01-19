package user_usecases

import (
	"errors"
	"testing"
	"time"

	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
	user_errors "github.com/luizrgf2/real-time-chat-go/internal/app/user/errors"
	user_interfaces_services "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/services"
	user_interfaces_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/usecases"
	tests_repositories "github.com/luizrgf2/real-time-chat-go/tests/repository"
	tests_services "github.com/luizrgf2/real-time-chat-go/tests/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pass = "Tetes123"
var currentTime = time.Now()

var validUserToTest = user_entities.UserEntity{
	ID:        nil,
	UserName:  "luizrg2",
	Email:     "luiz@email.com",
	Password:  &pass,
	CreatedAt: &currentTime,
	UpdatedAt: &currentTime,
}

func TestCreateValidUser(t *testing.T) {
	userRepo := new(tests_repositories.UserRepositoryMock)
	PassEncryptService := new(tests_services.PassEncryptServiceMock)
	passEncrypted := "sdfdsfa1221313"

	userRepo.On("FindByEmail", mock.Anything, &validUserToTest.Email).Return(nil, errors.New("usuário não existe"))
	userRepo.On("FindByUserName", mock.Anything, &validUserToTest.UserName).Return(nil, errors.New("usuário não existe"))
	userRepo.RepositoryBaseMock.On("Create", mock.Anything, mock.AnythingOfType("*user_entities.UserEntity")).Return(nil)
	PassEncryptService.On("EncryptPassword", validUserToTest.Password).Return(&user_interfaces_services.PassEncryptOutput{PasswordEncrypted: &passEncrypted}, nil)

	sut := CreateUserUseCaseImp{
		UserRepository:     userRepo,
		PassEncryptService: PassEncryptService,
	}

	res, err := sut.Exec(user_interfaces_usecases.CreateUserUseCaseInput{
		Email:    validUserToTest.Email,
		Username: validUserToTest.UserName,
		Password: *validUserToTest.Password,
	})

	assert.ErrorIs(t, err, nil)
	assert.Equal(t, validUserToTest.UserName, res.Username)
	assert.Equal(t, validUserToTest.Email, res.Email)
}

func TestCreateUserWithInvalidEmail(t *testing.T) {

	userWithInvalidEmail := validUserToTest
	userWithInvalidEmail.Email = "invalidEmail"

	userRepo := new(tests_repositories.UserRepositoryMock)
	PassEncryptService := new(tests_services.PassEncryptServiceMock)

	sut := CreateUserUseCaseImp{
		UserRepository:     userRepo,
		PassEncryptService: PassEncryptService,
	}

	_, err := sut.Exec(user_interfaces_usecases.CreateUserUseCaseInput{
		Email:    userWithInvalidEmail.Email,
		Username: userWithInvalidEmail.UserName,
		Password: *userWithInvalidEmail.Password,
	})

	assert.ErrorIs(t, err, user_errors.ErrInvalidEmail)
}

func TestCreateUserWithEmailAlreadyExists(t *testing.T) {

	userWithInvalidEmail := validUserToTest
	userWithInvalidEmail.Email = "emailAlreadyExists@email.com"

	userRepo := new(tests_repositories.UserRepositoryMock)
	PassEncryptService := new(tests_services.PassEncryptServiceMock)
	passEncrypted := "sdfdsfa1221313"

	userRepo.On("FindByEmail", mock.Anything, &userWithInvalidEmail.Email).Return(user_entities.UserEntity{}, nil)
	userRepo.On("FindByUserName", mock.Anything, &userWithInvalidEmail.UserName).Return(nil, errors.New("usuário não existe"))
	userRepo.RepositoryBaseMock.On("Create", mock.Anything, mock.AnythingOfType("*user_entities.UserEntity")).Return(nil)
	PassEncryptService.On("EncryptPassword", userWithInvalidEmail.Password).Return(&user_interfaces_services.PassEncryptOutput{PasswordEncrypted: &passEncrypted}, nil)

	sut := CreateUserUseCaseImp{
		UserRepository:     userRepo,
		PassEncryptService: PassEncryptService,
	}

	_, err := sut.Exec(user_interfaces_usecases.CreateUserUseCaseInput{
		Email:    userWithInvalidEmail.Email,
		Username: userWithInvalidEmail.UserName,
		Password: *userWithInvalidEmail.Password,
	})

	assert.ErrorIs(t, err, user_errors.ErrEmailAlrearyExists)
}

func TestCreateUserWithUserNameAlreadyExists(t *testing.T) {

	userWithInvalidEmail := validUserToTest
	userWithInvalidEmail.Email = "emailAlreadyExists@email.com"

	userRepo := new(tests_repositories.UserRepositoryMock)
	PassEncryptService := new(tests_services.PassEncryptServiceMock)
	passEncrypted := "sdfdsfa1221313"

	userRepo.On("FindByEmail", mock.Anything, &userWithInvalidEmail.Email).Return(nil, errors.New("error"))
	userRepo.On("FindByUserName", mock.Anything, &userWithInvalidEmail.UserName).Return(user_entities.UserEntity{}, nil)
	userRepo.RepositoryBaseMock.On("Create", mock.Anything, mock.AnythingOfType("*user_entities.UserEntity")).Return(nil)
	PassEncryptService.On("EncryptPassword", userWithInvalidEmail.Password).Return(&user_interfaces_services.PassEncryptOutput{PasswordEncrypted: &passEncrypted}, nil)

	sut := CreateUserUseCaseImp{
		UserRepository:     userRepo,
		PassEncryptService: PassEncryptService,
	}

	_, err := sut.Exec(user_interfaces_usecases.CreateUserUseCaseInput{
		Email:    userWithInvalidEmail.Email,
		Username: userWithInvalidEmail.UserName,
		Password: *userWithInvalidEmail.Password,
	})

	assert.ErrorIs(t, err, user_errors.ErrUserNameAlrearyExists)
}
