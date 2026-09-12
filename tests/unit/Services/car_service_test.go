package services_test

import (
	"context"
	"errors"
	"testing"

	domains "github.com/lucas-de-lima/car-shop-go/src/Domains"
	services "github.com/lucas-de-lima/car-shop-go/src/Services"
)

type mockCarRepo struct {
	createFunc  func(ctx context.Context, car domains.Car) (domains.Car, error)
	getAllFunc  func(ctx context.Context) ([]domains.Car, error)
	getByIDFunc func(ctx context.Context, id string) (domains.Car, error)
	updateFunc  func(ctx context.Context, id string, car domains.Car) (domains.Car, error)
	deleteFunc  func(ctx context.Context, id string) error
}

func (m *mockCarRepo) Create(ctx context.Context, car domains.Car) (domains.Car, error) {
	return m.createFunc(ctx, car)
}

func (m *mockCarRepo) GetAll(ctx context.Context) ([]domains.Car, error) {
	return m.getAllFunc(ctx)
}

func (m *mockCarRepo) GetByID(ctx context.Context, id string) (domains.Car, error) {
	return m.getByIDFunc(ctx, id)
}

func (m *mockCarRepo) Update(ctx context.Context, id string, car domains.Car) (domains.Car, error) {
	return m.updateFunc(ctx, id, car)
}

func (m *mockCarRepo) Delete(ctx context.Context, id string) error {
	return m.deleteFunc(ctx, id)
}

func TestCarService_Create(t *testing.T) {
	repo := &mockCarRepo{
		createFunc: func(ctx context.Context, car domains.Car) (domains.Car, error) {
			car.ID = "507f1f77bcf86cd799439011"
			return car, nil
		},
	}
	svc := services.NewCarService(repo)

	car := domains.NewCar("Marea", 2002, "Black", true, 15.99, 4, 5)
	result, err := svc.Create(context.Background(), car)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.GetModel() != "Marea" {
		t.Errorf("expected model Marea, got %s", result.GetModel())
	}
	if result.GetDoorsQty() != 4 {
		t.Errorf("expected doorsQty 4, got %d", result.GetDoorsQty())
	}
	if result.GetSeatsQty() != 5 {
		t.Errorf("expected seatsQty 5, got %d", result.GetSeatsQty())
	}
	if result.GetStatus() != true {
		t.Errorf("expected status true, got %v", result.GetStatus())
	}
}

func TestCarService_GetAll(t *testing.T) {
	expected := []domains.Car{
		domains.NewCar("Marea", 2002, "Black", true, 15.99, 4, 5),
		domains.NewCar("Tempra", 1995, "Black", false, 39.0, 2, 5),
	}

	repo := &mockCarRepo{
		getAllFunc: func(ctx context.Context) ([]domains.Car, error) {
			return expected, nil
		},
	}
	svc := services.NewCarService(repo)

	result, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 cars, got %d", len(result))
	}
}

func TestCarService_GetByID(t *testing.T) {
	car := domains.NewCar("Marea", 2002, "Black", true, 15.99, 4, 5)
	car.ID = "507f1f77bcf86cd799439011"

	repo := &mockCarRepo{
		getByIDFunc: func(ctx context.Context, id string) (domains.Car, error) {
			if id == car.ID {
				return car, nil
			}
			return domains.Car{}, errors.New("not found")
		},
	}
	svc := services.NewCarService(repo)

	result, err := svc.GetByID(context.Background(), car.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.GetModel() != "Marea" {
		t.Errorf("expected model Marea, got %s", result.GetModel())
	}
}

func TestCarService_GetByID_NotFound(t *testing.T) {
	repo := &mockCarRepo{
		getByIDFunc: func(ctx context.Context, id string) (domains.Car, error) {
			return domains.Car{}, errors.New("not found")
		},
	}
	svc := services.NewCarService(repo)

	_, err := svc.GetByID(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCarService_Delete(t *testing.T) {
	repo := &mockCarRepo{
		deleteFunc: func(ctx context.Context, id string) error {
			return nil
		},
	}
	svc := services.NewCarService(repo)

	err := svc.Delete(context.Background(), "507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCarService_Delete_NotFound(t *testing.T) {
	repo := &mockCarRepo{
		deleteFunc: func(ctx context.Context, id string) error {
			return errors.New("not found")
		},
	}
	svc := services.NewCarService(repo)

	err := svc.Delete(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCarService_Update(t *testing.T) {
	repo := &mockCarRepo{
		updateFunc: func(ctx context.Context, id string, car domains.Car) (domains.Car, error) {
			car.ID = id
			return car, nil
		},
	}
	svc := services.NewCarService(repo)

	car := domains.NewCar("Marea", 1992, "Red", true, 12.0, 2, 5)
	result, err := svc.Update(context.Background(), "507f1f77bcf86cd799439011", car)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.GetModel() != "Marea" {
		t.Errorf("expected model Marea, got %s", result.GetModel())
	}
	if result.GetYear() != 1992 {
		t.Errorf("expected year 1992, got %d", result.GetYear())
	}
}