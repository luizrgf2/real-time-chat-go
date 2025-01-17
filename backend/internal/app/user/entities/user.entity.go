package user_entities

import (
	"regexp"
	"time"

	user_errors "github.com/luizrgf2/real-time-chat-go/internal/app/user/errors"
)

type UserEntity struct {
	ID        *uint
	UserName  string
	Email     string
	Password  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (u *UserEntity) ValidateID() error {
	if u.ID == nil || *u.ID <= 0 {
		return user_errors.ErrInvalidID
	}
	return nil
}

func (u *UserEntity) ValidateEmail() error {
	if u.Email == "" {
		return user_errors.ErrEmptyEmail
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return user_errors.ErrInvalidEmail
	}

	return nil
}

func (u *UserEntity) ValidateUserName() error {
	if len(u.UserName) < 4 {
		return user_errors.ErrShortUserName
	}
	return nil
}

func (u *UserEntity) ValidatePassword() error {
	if u.Password == nil || len(*u.Password) < 8 {
		return user_errors.ErrShortPassword
	}
	return nil
}

func (u *UserEntity) ValidateCreatedAt() error {
	if u.CreatedAt == nil || u.CreatedAt.IsZero() {
		return user_errors.ErrInvalidCreatedAt
	}
	return nil
}

func (u *UserEntity) ValidateUpdatedAt() error {
	if u.UpdatedAt == nil || u.UpdatedAt.IsZero() {
		return user_errors.ErrInvalidUpdatedAt
	}
	return nil
}

func (u *UserEntity) Validate() error {
	if err := u.ValidateID(); err != nil {
		return err
	}

	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if err := u.ValidateUserName(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	if err := u.ValidateCreatedAt(); err != nil {
		return err
	}

	if err := u.ValidateUpdatedAt(); err != nil {
		return err
	}

	return nil
}

func (u *UserEntity) ValidateWithoutId() error {
	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if err := u.ValidateUserName(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	if err := u.ValidateCreatedAt(); err != nil {
		return err
	}

	if err := u.ValidateUpdatedAt(); err != nil {
		return err
	}

	return nil
}

func CreateUser(ID uint, Email string, UserName string, Password string) (*UserEntity, error) {

	curretTime := time.Now()

	user := UserEntity{
		ID:        &ID,
		Email:     Email,
		UserName:  UserName,
		Password:  &Password,
		CreatedAt: &curretTime,
		UpdatedAt: &curretTime,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateWithoutId(Email string, UserName string, Password string) (*UserEntity, error) {

	curretTime := time.Now()

	user := UserEntity{
		ID:        nil,
		Email:     Email,
		UserName:  UserName,
		Password:  &Password,
		CreatedAt: &curretTime,
		UpdatedAt: &curretTime,
	}

	err := user.ValidateWithoutId()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
