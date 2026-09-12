package interfaces

import "context"

type IRepository[T any] interface {
	Create(ctx context.Context, entity T) (T, error)
	GetAll(ctx context.Context) ([]T, error)
	GetByID(ctx context.Context, id string) (T, error)
	Update(ctx context.Context, id string, entity T) (T, error)
	Delete(ctx context.Context, id string) error
}