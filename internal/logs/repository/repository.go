package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/defryheryanto/nebula/internal/logs"
	"github.com/defryheryanto/nebula/pkg/mongoconverter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *Repository) Find(ctx context.Context, filter *logs.Filter) ([]*logs.Log, error) {
	opt := options.Find().SetSort(bson.M{"timestamp": -1})
	isPagination, page, pageSize := filter.GetPagination()
	if isPagination {
		skip := (page - 1) * pageSize
		opt = opt.SetLimit(int64(pageSize)).SetSkip(int64(skip))
	}

	queryFilter := bson.M{}
	if filter.ServiceName != "" {
		queryFilter["service"] = filter.ServiceName
	}

	cur, err := r.db.Find(ctx, queryFilter, opt)
	if err != nil {
		return nil, err
	}

	var result []*logs.Log
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	for _, res := range result {
		if bsonD, ok := res.Log.(bson.D); ok {
			res.Log = mongoconverter.BsonDToMap(bsonD)
		}
	}

	return result, nil
}

func (r *Repository) AvailableServices(ctx context.Context) ([]string, error) {
	values, err := r.db.Distinct(ctx, "service", bson.M{})
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(values))

	for _, val := range values {
		result = append(result, fmt.Sprintf("%v", val))
	}

	return result, nil
}
