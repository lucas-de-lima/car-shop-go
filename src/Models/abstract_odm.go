package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AbstractODM[T any] struct {
	collection *mongo.Collection
}

func NewAbstractODM[T any](db *mongo.Database, collectionName string) *AbstractODM[T] {
	return &AbstractODM[T]{collection: db.Collection(collectionName)}
}

func (o *AbstractODM[T]) Create(ctx context.Context, entity T) (T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := o.collection.InsertOne(ctx, entity)
	if err != nil {
		return entity, fmt.Errorf("failed to insert: %w", err)
	}

	_ = result.InsertedID

	return entity, nil
}

func (o *AbstractODM[T]) GetAll(ctx context.Context) ([]T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := o.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find: %w", err)
	}
	defer cursor.Close(ctx)

	var results []T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	return results, nil
}

func (o *AbstractODM[T]) GetByID(ctx context.Context, id string) (T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("invalid mongo id")
	}

	var result T
	err = o.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		var zero T
		if err == mongo.ErrNoDocuments {
			return zero, fmt.Errorf("not found")
		}
		return zero, fmt.Errorf("failed to find: %w", err)
	}

	return result, nil
}

func (o *AbstractODM[T]) Update(ctx context.Context, id string, entity T) (T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("invalid mongo id")
	}

	result, err := o.collection.ReplaceOne(ctx, bson.M{"_id": objectID}, entity)
	if err != nil {
		return entity, fmt.Errorf("failed to update: %w", err)
	}

	if result.MatchedCount == 0 {
		var zero T
		return zero, fmt.Errorf("not found")
	}

	return entity, nil
}

func (o *AbstractODM[T]) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid mongo id")
	}

	result, err := o.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}

var _options = options.Client().ApplyURI("mongodb://localhost:27017")

func Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, _options)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	return client, nil
}

func GetDatabase(client *mongo.Client) *mongo.Database {
	return client.Database("CarShop")
}