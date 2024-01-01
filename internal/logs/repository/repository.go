package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func New(db *mongo.Collection) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Insert(ctx context.Context, data any) error {
	type logData struct {
		Timestamp time.Time
		Log       any
	}
	_, err := r.db.InsertOne(ctx, &logData{
		Timestamp: time.Now(),
		Log:       data,
	})
	if err != nil {
		return err
	}

	return nil
}
