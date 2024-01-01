package main

import (
	"context"
	"database/sql"

	"github.com/defryheryanto/nebula/config"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDatabaseConnection(ctx context.Context) *sql.DB {
	conn, err := sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		panic(err)
	}

	return conn
}

func setupMongoClient(ctx context.Context) *mongo.Client {
	opt := options.Client().ApplyURI(config.MongoDBConnectionString)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}

	return client
}
