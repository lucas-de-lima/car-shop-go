package services_test

import (
	"context"
	"errors"
	"testing"

	domains "github.com/lucas-de-lima/car-shop-go/src/Domains"
	services "github.com/lucas-de-lima/car-shop-go/src/Services"
)

type mockMotorcycleRepo struct {
	createFunc  func(ctx context.Context, m domains.Motorcycle) (domains.Motorcycle, error)
	getAllFunc  func(ctx context.Context) ([]domains.Motorcycle, error)
	getByIDFunc func(ctx context.Context, id string) (domains.Motorcycle, error)
	updateFunc  func(ctx context.Context, id string, m domains.Motorcycle) (domains.Motorcycle, error)
	deleteFunc  func(ctx context.Context, id string) error
}

func (m *mockMotorcycleRepo) Create(ctx context.Context, motorcycle domains.Motorcycle) (domains.Motorcycle, error) {
	return m.createFunc(ctx, motorcycle)
}

func (m *mockMotorcycleRepo) GetAll(ctx context.Context) ([]domains.Motorcycle, error) {
	return m.getAllFunc(ctx)
}

func (m *mockMotorcycleRepo) GetByID(ctx context.Context, id string) (domains.Motorcycle, error) {
	return m.getByIDFunc(ctx, id)
}

func (m *mockMotorcycleRepo) Update(ctx context.Context, id string, motorcycle domains.Motorcycle) (domains.Motorcycle, error) {
	return m.updateFunc(ctx, id, motorcycle)
}

func (m *mockMotorcycleRepo) Delete(ctx context.Context, id string) error {
	return m.deleteFunc(ctx, id)
}

func TestMotorcycleService_Create(t *testing.T) {
	repo := &mockMotorcycleRepo{
		createFunc: func(ctx context.Context, m domains.Motorcycle) (domains.Motorcycle, error) {
			m.ID = "507f1f77bcf86cd799439011"
			return m, nil
		},
	}
	svc := services.NewMotorcycleService(repo)

	m := domains.NewMotorcycle("Honda Cb 600f Hornet", 2005, "Yellow", true, 30.0, "Street", 600)
	result, err := svc.Create(context.Background(), m)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.GetModel() != "Honda Cb 600f Hornet" {
		t.Errorf("expected model 'Honda Cb 600f Hornet', got %s", result.GetModel())
	}
	if result.GetCategory() != "Street" {
		t.Errorf("expected category 'Street', got %s", result.GetCategory())
	}
	if result.GetEngineCapacity() != 600 {
		t.Errorf("expected engineCapacity 600, got %d", result.GetEngineCapacity())
	}
}

func TestMotorcycleService_GetAll(t *testing.T) {
	expected := []domains.Motorcycle{
		domains.NewMotorcycle("Honda Cb 600f Hornet", 2005, "Yellow", true, 30.0, "Street", 600),
		domains.NewMotorcycle("Honda Cbr 1000rr", 2011, "Orange", true, 59.9, "Street", 1000),
	}

	repo := &mockMotorcycleRepo{
		getAllFunc: func(ctx context.Context) ([]domains.Motorcycle, error) {
			return expected, nil
		},
	}
	svc := services.NewMotorcycleService(repo)

	result, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 motorcycles, got %d", len(result))
	}
}

func TestMotorcycleService_GetByID(t *testing.T) {
	m := domains.NewMotorcycle("Honda Cb 600f Hornet", 2005, "Yellow", true, 30.0, "Street", 600)
	m.ID = "507f1f77bcf86cd799439011"

	repo := &mockMotorcycleRepo{
		getByIDFunc: func(ctx context.Context, id string) (domains.Motorcycle, error) {
			if id == m.ID {
				return m, nil
			}
			return domains.Motorcycle{}, errors.New("not found")
		},
	}
	svc := services.NewMotorcycleService(repo)

	result, err := svc.GetByID(context.Background(), m.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.GetModel() != "Honda Cb 600f Hornet" {
		t.Errorf("expected model 'Honda Cb 600f Hornet', got %s", result.GetModel())
	}
}

func TestMotorcycleService_GetByID_NotFound(t *testing.T) {
	repo := &mockMotorcycleRepo{
		getByIDFunc: func(ctx context.Context, id string) (domains.Motorcycle, error) {
			return domains.Motorcycle{}, errors.New("not found")
		},
	}
	svc := services.NewMotorcycleService(repo)

	_, err := svc.GetByID(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestMotorcycleService_Update(t *testing.T) {
	repo := &mockMotorcycleRepo{
		updateFunc: func(ctx context.Context, id string, m domains.Motorcycle) (domains.Motorcycle, error) {
			m.ID = id
			return m, nil
		},
	}
	svc := services.NewMotorcycleService(repo)

	m := domains.NewMotorcycle("Honda Cb 600f Hornet", 2014, "Red", true, 45.0, "Street", 600)
	result, err := svc.Update(context.Background(), "507f1f77bcf86cd799439011", m)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.GetYear() != 2014 {
		t.Errorf("expected year 2014, got %d", result.GetYear())
	}
	if result.GetColor() != "Red" {
		t.Errorf("expected color Red, got %s", result.GetColor())
	}
}

func TestMotorcycleService_Delete(t *testing.T) {
	repo := &mockMotorcycleRepo{
		deleteFunc: func(ctx context.Context, id string) error {
			return nil
		},
	}
	svc := services.NewMotorcycleService(repo)

	err := svc.Delete(context.Background(), "507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMotorcycleService_Delete_NotFound(t *testing.T) {
	repo := &mockMotorcycleRepo{
		deleteFunc: func(ctx context.Context, id string) error {
			return errors.New("not found")
		},
	}
	svc := services.NewMotorcycleService(repo)

	err := svc.Delete(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}