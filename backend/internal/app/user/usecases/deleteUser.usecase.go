package user_usecases

import (
	"context"
	"time"

	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
	user_interfaces_repository "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/repositories"
	user_interfaces_usecases "github.com/luizrgf2/real-time-chat-go/internal/app/user/interfaces/usecases"
)

type DeleteUserUseCaseImp struct {
	UserRepository user_interfaces_repository.IUserRepository
}

func (d *DeleteUserUseCaseImp) checkIfUserExists(ctx context.Context, id uint) (*user_entities.UserEntity, error) {
	user, err := d.UserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DeleteUserUseCaseImp) deleteUserById(ctx context.Context, id uint) error {
	err := d.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeleteUserUseCaseImp) Exec(input user_interfaces_usecases.DeleteUserUseCaseInput) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := d.checkIfUserExists(ctx, input.ID)
	if err != nil {
		return err
	}

	err = d.deleteUserById(ctx, input.ID)
	if err != nil {
		return err
	}

	return err
}
