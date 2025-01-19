package tests_services

import (
	user_interfaces_services "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/services"
	"github.com/stretchr/testify/mock"
)

type BcryptServiceMock struct {
	mock.Mock
}

func (b *BcryptServiceMock) EncryptPassword(passwordDecrypt *string) (*user_interfaces_services.PassEncryptOutput, error) {
	args := b.Called(passwordDecrypt)
	value := args.Get(0)
	err := args.Error(1)

	if value == nil {
		return nil, err
	}

	valueToReturn := value.(*user_interfaces_services.PassEncryptOutput)

	return valueToReturn, err
}

func (b *BcryptServiceMock) ValidatePassword(passEncrypted *string, passwordDecrypted *string) (bool, error) {
	args := b.Called(passEncrypted, passwordDecrypted)
	value := args.Bool(0)
	err := args.Error(1)
	return value, err
}
