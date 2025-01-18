package user_interfaces_repository

import (
	"context"

	app_shared_interfaces "github.com/luizrgf2/real-time-chat-go/internal/app/shared/interfaces"
	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
)

type IUserRepository interface {
	app_shared_interfaces.BaseRepository[user_entities.UserEntity]
	FindByEmail(ctx context.Context, email *string) (*user_entities.UserEntity, error)
	FindByUserName(ctx context.Context, userName *string) (*user_entities.UserEntity, error)
}
