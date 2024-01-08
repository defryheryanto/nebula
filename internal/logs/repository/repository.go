package repository

import (
	"context"
	"time"

	"github.com/defryheryanto/nebula/internal/logs"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *Repository) Insert(ctx context.Context, service string, data any) error {
	type logData struct {
		Timestamp time.Time
		Service   string
		Log       any
	}
	_, err := r.db.InsertOne(ctx, &logData{
		Timestamp: time.Now(),
		Service:   service,
		Log:       data,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Find(ctx context.Context) ([]*logs.Log, error) {
	cur, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var result []*logs.Log
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
