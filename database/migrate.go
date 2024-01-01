package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/defryheryanto/nebula/config"
	_ "github.com/golang-migrate/migrate/source"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

	downFlag := flag.Bool("down", false, "database migration down")
	flag.Parse()

	dsn := config.DBConnectionString
	if dsn == "" {
		slog.Error("database connection string is empty")
		return
	}

	slog.Info(fmt.Sprintf("Opening database connection to %s", dsn))
	db, err := sql.Open("postgres", dsn)
	slog.Info("database connected")
	if err != nil {
		slog.Error("error opening migration database", "error", err)
		return
	}

	slog.Info("generating postgres instance")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error("error generating postgres instance", "error", err)
		return
	}
	slog.Info("postgres instance generated")

	slog.Info("opening migration files")
	fsrc, err := (&file.File{}).Open("file://database/migrations")
	if err != nil {
		slog.Error("error opening migration files", "error", err)
		return
	}
	slog.Info("migration files opened")

	slog.Info("creating migration instance")
	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		slog.Error("error generating migrate instance", "error", err)
		return
	}
	slog.Info("migration instance created")

	if *downFlag {
		slog.Info("rolling back migration")
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			slog.Error("error rollback migrations", "error", err)
			return
		}
		version, _, _ := m.Version()
		slog.Info(fmt.Sprintf("rolled back to version %d.\n", version))
	} else {
		slog.Info("migrating migration")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			slog.Error("error migrating migrations", "error", err)
			return
		}
		slog.Info("migrate complete")
	}
}
