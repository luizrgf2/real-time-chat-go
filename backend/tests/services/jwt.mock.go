package tests_services

import "github.com/stretchr/testify/mock"

type JWTServiceMock struct {
	mock.Mock
}

func (j *JWTServiceMock) Encode(payload interface{}) (*string, error) {
	args := j.Called(payload)
	value := args.Get(0)
	err := args.Error(1)

	if value == nil {
		return nil, err
	}

	return value.(*string), err
}

func (j *JWTServiceMock) Decode(token *string) (*interface{}, error) {
	args := j.Called(token)
	value := args.Get(0)
	err := args.Error(1)
	return &value, err
}
