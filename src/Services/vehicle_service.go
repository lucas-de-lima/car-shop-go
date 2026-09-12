package services

import (
	"context"
	"fmt"

	domains "github.com/lucas-de-lima/car-shop-go/src/Domains"
	interfaces "github.com/lucas-de-lima/car-shop-go/src/Interfaces"
)

type VehicleService[T interfaces.IVehicle] struct {
	repo interfaces.IRepository[T]
}

func NewVehicleService[T interfaces.IVehicle](repo interfaces.IRepository[T]) *VehicleService[T] {
	return &VehicleService[T]{repo: repo}
}

func (s *VehicleService[T]) Create(ctx context.Context, entity T) (T, error) {
	if entity.GetModel() == "" {
		var zero T
		return zero, fmt.Errorf("model is required")
	}
	return s.repo.Create(ctx, entity)
}

func (s *VehicleService[T]) GetAll(ctx context.Context) ([]T, error) {
	return s.repo.GetAll(ctx)
}

func (s *VehicleService[T]) GetByID(ctx context.Context, id string) (T, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *VehicleService[T]) Update(ctx context.Context, id string, entity T) (T, error) {
	return s.repo.Update(ctx, id, entity)
}

func (s *VehicleService[T]) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

type CarService struct {
	*VehicleService[domains.Car]
}

func NewCarService(repo interfaces.IRepository[domains.Car]) *CarService {
	return &CarService{
		VehicleService: NewVehicleService[domains.Car](repo),
	}
}

type MotorcycleService struct {
	*VehicleService[domains.Motorcycle]
}

func NewMotorcycleService(repo interfaces.IRepository[domains.Motorcycle]) *MotorcycleService {
	return &MotorcycleService{
		VehicleService: NewVehicleService[domains.Motorcycle](repo),
	}
}