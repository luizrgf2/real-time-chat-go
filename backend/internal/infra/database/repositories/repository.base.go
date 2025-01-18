package repositories

import (
	"context"

	"github.com/luizrgf2/real-time-chat-go/internal/infra/database"
)

type RepositoryBase[T any] struct {
	db *database.Database
}

func (r *RepositoryBase[T]) Create(ctx context.Context, entity *T) error {
	result := r.db.Db.WithContext(ctx).Create(entity)
	return result.Error
}

func (r *RepositoryBase[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	result := r.db.Db.WithContext(ctx).First(&entity, id)
	return &entity, result.Error
}

func (r *RepositoryBase[T]) Update(ctx context.Context, entity *T) error {
	result := r.db.Db.WithContext(ctx).Save(entity)
	return result.Error
}

func (r *RepositoryBase[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	result := r.db.Db.WithContext(ctx).Delete(&entity, id)
	return result.Error
}

func (r *RepositoryBase[T]) FindAll(ctx context.Context) ([]T, error) {
	var entities []T
	result := r.db.Db.WithContext(ctx).Find(&entities)
	return entities, result.Error
}
