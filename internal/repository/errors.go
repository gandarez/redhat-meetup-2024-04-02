package repository

import (
	"fmt"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// handleFindError handles the error when finding a record.
func handleFindError(err error, params ...any) error {
	if err == pgx.ErrNoRows {
		return database.ErrNotFound(fmt.Sprintf("%s with id %s not found", params...))
	}

	msg := fmt.Sprintf("failed to find %s", params[0])

	return fmt.Errorf("%s: %w", msg, err)
}

// handleSaveError handles the error when saving a record.
func handleSaveError(err error, params ...any) error {
	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
		return database.ErrConflict(fmt.Sprintf("%s already exists", params...))
	}

	msg := fmt.Sprintf("failed to save %s", params...)

	return fmt.Errorf("%s: %w", msg, err)
}
