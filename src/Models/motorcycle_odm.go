package models

import (
	"context"

	domains "github.com/lucas-de-lima/car-shop-go/src/Domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MotorcycleODM struct {
	*AbstractODM[domains.Motorcycle]
}

func NewMotorcycleODM(db *mongo.Database) *MotorcycleODM {
	return &MotorcycleODM{
		AbstractODM: NewAbstractODM[domains.Motorcycle](db, "motorcycles"),
	}
}

func (o *MotorcycleODM) Create(ctx context.Context, m domains.Motorcycle) (domains.Motorcycle, error) {
	m.Vehicle.ID = primitive.NewObjectID().Hex()
	result, err := o.AbstractODM.Create(ctx, m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (o *MotorcycleODM) GetByID(ctx context.Context, id string) (domains.Motorcycle, error) {
	return o.AbstractODM.GetByID(ctx, id)
}

func (o *MotorcycleODM) GetAll(ctx context.Context) ([]domains.Motorcycle, error) {
	return o.AbstractODM.GetAll(ctx)
}

func (o *MotorcycleODM) Update(ctx context.Context, id string, m domains.Motorcycle) (domains.Motorcycle, error) {
	return o.AbstractODM.Update(ctx, id, m)
}

func (o *MotorcycleODM) Delete(ctx context.Context, id string) error {
	return o.AbstractODM.Delete(ctx, id)
}