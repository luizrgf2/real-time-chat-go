package tests_services

import (
	user_interfaces_services "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/services"
	"github.com/stretchr/testify/mock"
)

type PassEncryptServiceMock struct {
	mock.Mock
}

func (p *PassEncryptServiceMock) EncryptPassword(passwordDecrypt *string) (*user_interfaces_services.PassEncryptOutput, error) {
	args := p.Called(passwordDecrypt)
	value := args.Get(0).(*user_interfaces_services.PassEncryptOutput)
	err := args.Error(1)
	return value, err
}
