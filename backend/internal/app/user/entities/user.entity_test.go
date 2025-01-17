package user

import (
	"testing"

	user_errors "github.com/luizrgf2/real-time-chat-go/internal/app/user/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidUser(t *testing.T) {
	_, err := CreateUser(1, "test@email.com.br", "luizrgfg", "testePass123")
	assert.Nil(t, err)
}

func TestCreateUserWithInvalidUsername(t *testing.T) {
	_, err := CreateUser(1, "test@email.com.br", "l", "testePass123")
	assert.ErrorIs(t, err, user_errors.ErrShortUserName)

	_, err = CreateWithoutId("test@email.com.br", "l", "testePass123")
	assert.ErrorIs(t, err, user_errors.ErrShortUserName)
}

func TestCreateValidUserWithoutId(t *testing.T) {
	_, err := CreateWithoutId("test@email.com.br", "luizrgfg", "testePass123")
	assert.Nil(t, err)
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	_, err := CreateUser(1, "invalid", "luizrgfg", "testPass123")
	assert.ErrorIs(t, err, user_errors.ErrInvalidEmail)

	_, err = CreateWithoutId("invalid", "luizrgfg", "testPass123")
	assert.ErrorIs(t, err, user_errors.ErrInvalidEmail)
}

func TestCreateUserWithEmptyEmail(t *testing.T) {
	_, err := CreateUser(1, "", "luizrgfg", "testPass123")
	assert.ErrorIs(t, err, user_errors.ErrEmptyEmail)

	_, err = CreateWithoutId("", "luizrgfg", "testPass123")
	assert.ErrorIs(t, err, user_errors.ErrEmptyEmail)
}

func TestCreateUserWithInvalidPassword(t *testing.T) {
	_, err := CreateUser(1, "test@email.com.br", "luizrgfg", "invalid")
	assert.ErrorIs(t, err, user_errors.ErrShortPassword)
	_, err = CreateWithoutId("test@email.com.br", "luizrgfg", "invalid")
	assert.ErrorIs(t, err, user_errors.ErrShortPassword)
}
