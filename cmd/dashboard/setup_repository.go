package main

import (
	"database/sql"

	userrepository "github.com/defryheryanto/nebula/internal/user/repository"
)

type repositories struct {
	UserRepository *userrepository.Repository
}

func setupRepositories(db *sql.DB) *repositories {
	userRepository := setupUserRepository(db)

	return &repositories{
		UserRepository: userRepository,
	}
}

func setupUserRepository(db *sql.DB) *userrepository.Repository {
	return userrepository.New(db)
}
