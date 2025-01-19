package user_usecases

import (
	"testing"
	"time"

	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
	user_errors "github.com/luizrgf2/real-time-chat-go/internal/app/user/errors"
	user_interfaces_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/usecases"
	tests_repositories "github.com/luizrgf2/real-time-chat-go/tests/repository"
	tests_services "github.com/luizrgf2/real-time-chat-go/tests/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var passToTest = "validPassword"
var tokenJWTToTest = "sfsdfsafasf2234"

var id = uint(1)
var now = time.Now()

var validUserToTestAuthUseCase = user_entities.UserEntity{
	ID:        &id,
	UserName:  "testeUsername",
	Email:     "teste@email",
	Password:  &passToTest,
	CreatedAt: &now,
	UpdatedAt: &now,
}

func TestAuthWithValidEmailAndPassword(t *testing.T) {
	bcryptService := new(tests_services.BcryptServiceMock)
	jwtService := new(tests_services.JWTServiceMock)
	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("*string")).Return(&validUserToTestAuthUseCase, nil)
	bcryptService.On("ValidatePassword", mock.AnythingOfType("*string"), mock.AnythingOfType("*string")).Return(true, nil)
	jwtService.On("Encode", mock.Anything).Return(&tokenJWTToTest, nil)

	sut := AuthUserUseCaseImp{
		UserRepository:     userRepo,
		PassEncryptService: bcryptService,
		JwtService:         jwtService,
	}

	res, err := sut.Exec(user_interfaces_usecases.AuthUserUseCaseInput{
		Email:    validUserToTest.Email,
		Password: *validUserToTest.Password,
	})

	assert.ErrorIs(t, err, nil)
	assert.Equal(t, tokenJWTToTest, res.Token)
}

func TestAuthWithEmailNotExists(t *testing.T) {
	bcryptService := new(tests_services.BcryptServiceMock)
	jwtService := new(tests_services.JWTServiceMock)
	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("*string")).Return(nil, user_errors.ErrAuthWrong)
	bcryptService.On("ValidatePassword", mock.AnythingOfType("*string"), mock.AnythingOfType("*string")).Return(true, nil)
	jwtService.On("Encode", mock.Anything).Return(&tokenJWTToTest, nil)

	sut := AuthUserUseCaseImp{
		UserRepository:     userRepo,
		PassEncryptService: bcryptService,
		JwtService:         jwtService,
	}

	_, err := sut.Exec(user_interfaces_usecases.AuthUserUseCaseInput{
		Email:    validUserToTest.Email,
		Password: *validUserToTest.Password,
	})

	assert.ErrorIs(t, err, user_errors.ErrAuthWrong)
}

func TestAuthWithPasswordWrong(t *testing.T) {
	bcryptService := new(tests_services.BcryptServiceMock)
	jwtService := new(tests_services.JWTServiceMock)
	userRepo := new(tests_repositories.UserRepositoryMock)

	userRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("*string")).Return(&validUserToTestAuthUseCase, nil)
	bcryptService.On("ValidatePassword", mock.AnythingOfType("*string"), mock.AnythingOfType("*string")).Return(false, nil)
	jwtService.On("Encode", mock.Anything).Return(&tokenJWTToTest, nil)

	sut := AuthUserUseCaseImp{
		UserRepository:     userRepo,
		PassEncryptService: bcryptService,
		JwtService:         jwtService,
	}

	_, err := sut.Exec(user_interfaces_usecases.AuthUserUseCaseInput{
		Email:    validUserToTest.Email,
		Password: *validUserToTest.Password,
	})

	assert.ErrorIs(t, err, user_errors.ErrAuthWrong)
}
