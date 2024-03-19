package db

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // nolint:revive
	_ "github.com/golang-migrate/migrate/v4/source/file"       // nolint:revive
)

// Run runs the database migrations.
func Run(connString string, logger *slog.Logger) error {
	logger.Debug("starting database migrations")

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %s", err)
	}

	var path string

	dockerized := os.Getenv("DOCKER")
	if dockerized == "true" {
		path = filepath.Join(wd, "db/migrations")
	} else {
		path = filepath.Join(wd, "../../", "db/migrations")
	}

	m, err := migrate.New(
		"file://"+path,
		connString,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize database for migration: %s", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Debug("no migrations to run")

			return nil
		}

		return fmt.Errorf("failed to run migrations: %s", err)
	}

	logger.Debug("database migrations completed")

	return nil
}
