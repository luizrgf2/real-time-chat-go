package repositories

import (
	user_entities "github.com/luizrgf2/real-time-chat-go/internal/app/user/entities"
)

type UserRepository struct {
	RepositoryBase[user_entities.UserEntity]
}
