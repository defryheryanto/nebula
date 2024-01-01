package main

import (
	"context"
	"database/sql"

	"github.com/defryheryanto/nebula/config"
	_ "github.com/lib/pq"
)

func setupDatabaseConnection(ctx context.Context) *sql.DB {
	conn, err := sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		panic(err)
	}

	return conn
}
