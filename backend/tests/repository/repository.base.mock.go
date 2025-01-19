package tests_repositories

import (
	"context"
	"reflect"

	"github.com/stretchr/testify/mock"
)

type RepositoryBaseMock[T any] struct {
	mock.Mock
}

func (r *RepositoryBaseMock[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*T), args.Error(1)
}

func (r *RepositoryBaseMock[T]) Update(ctx context.Context, entity *T) error {
	args := r.Called(ctx, entity)
	return args.Error(0)
}

func (r *RepositoryBaseMock[T]) Delete(ctx context.Context, id uint) error {
	args := r.Called(ctx, id)
	return args.Error(0)
}

func (r *RepositoryBaseMock[T]) FindAll(ctx context.Context) ([]T, error) {
	args := r.Called(ctx)
	return args.Get(0).([]T), args.Error(1)
}

func (r *RepositoryBaseMock[T]) Create(ctx context.Context, entity *T) error {
	args := r.Called(ctx, entity)
	ID := uint(1)
	if entity != nil {
		idField := reflect.ValueOf(entity).Elem().FieldByName("ID")
		if idField.IsValid() && idField.Kind() == reflect.Ptr {
			idField.Set(reflect.ValueOf(&ID))
		}
	}
	return args.Error(0)
}
