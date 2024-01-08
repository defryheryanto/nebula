package main

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/defryheryanto/nebula/config"
	logsrepository "github.com/defryheryanto/nebula/internal/logs/repository"
	userrepository "github.com/defryheryanto/nebula/internal/user/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositories struct {
	UserRepository *userrepository.Repository
	LogRepository  *logsrepository.Repository
}

func setupRepositories(db *sql.DB, mongoClient *mongo.Client) *repositories {
	userRepository := setupUserRepository(db)
	logRepository := setupLogsRepository(mongoClient)

	return &repositories{
		UserRepository: userRepository,
		LogRepository:  logRepository,
	}
}

func setupUserRepository(db *sql.DB) *userrepository.Repository {
	return userrepository.New(db)
}

func setupLogsRepository(mongoClient *mongo.Client) *logsrepository.Repository {
	collection := mongoClient.Database(config.MongoDBName).Collection("logs")
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "service", Value: 1},
		},
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		slog.Error("failed create logs indexes", "error", err)
	}

	return logsrepository.New(collection)
}
