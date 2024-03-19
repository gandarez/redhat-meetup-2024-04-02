package repository

import (
	"context"
	"fmt"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/entity"

	"github.com/jackc/pgx/v5"
)

type (
	// DatabaseQueryExecutor is the interface for executing database queries.
	DatabaseQueryExecutor interface {
		Exec(ctx context.Context, sql string, args ...any) (int64, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	}

	// Console is the repository for console.
	Console struct {
		db DatabaseQueryExecutor
	}
)

// NewConsole creates a new console repository.
func NewConsole(db DatabaseQueryExecutor) *Console {
	return &Console{
		db: db,
	}
}

// FindByID returns a console by id.
func (c *Console) FindByID(ctx context.Context, id string) (*entity.Console, error) {
	var console entity.Console

	err := c.db.QueryRow(
		ctx,
		"SELECT id, name, manufacturer, release_date FROM console WHERE id = $1", id).
		Scan(&console.ID, &console.Name, &console.Manufacturer, &console.ReleaseDate)
	if err != nil {
		return nil, handleFindError(err, "console", id)
	}

	return &console, nil
}

// Save saves a console.
func (c *Console) Save(ctx context.Context, console *entity.Console) error {
	affected, err := c.db.Exec(
		ctx,
		"INSERT INTO console (id, name, manufacturer, release_date) VALUES ($1, $2, $3, $4)",
		console.ID,
		console.Name,
		console.Manufacturer,
		console.ReleaseDate,
	)
	if err != nil {
		return handleSaveError(err, "console")
	}

	if affected != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", affected)
	}

	return nil
}
