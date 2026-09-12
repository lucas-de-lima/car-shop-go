package models

import (
	"context"

	domains "github.com/lucas-de-lima/car-shop-go/src/Domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarODM struct {
	*AbstractODM[domains.Car]
}

func NewCarODM(db *mongo.Database) *CarODM {
	return &CarODM{
		AbstractODM: NewAbstractODM[domains.Car](db, "cars"),
	}
}

func (o *CarODM) Create(ctx context.Context, car domains.Car) (domains.Car, error) {
	car.Vehicle.ID = primitive.NewObjectID().Hex()
	result, err := o.AbstractODM.Create(ctx, car)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (o *CarODM) GetByID(ctx context.Context, id string) (domains.Car, error) {
	return o.AbstractODM.GetByID(ctx, id)
}

func (o *CarODM) GetAll(ctx context.Context) ([]domains.Car, error) {
	return o.AbstractODM.GetAll(ctx)
}

func (o *CarODM) Update(ctx context.Context, id string, car domains.Car) (domains.Car, error) {
	return o.AbstractODM.Update(ctx, id, car)
}

func (o *CarODM) Delete(ctx context.Context, id string) error {
	return o.AbstractODM.Delete(ctx, id)
}