package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/entity"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/repository"

	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConsoleFindByID(t *testing.T) {
	db, conn, teardown := setupTestDb(t)
	defer teardown()

	releaseDate, err := time.Parse(time.DateOnly, "2020-11-12")
	require.NoError(t, err)

	result := &entity.Console{
		ID:           uuid.New(),
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  releaseDate,
	}

	conn.
		ExpectQuery(`
			SELECT id, name, manufacturer, release_date
			FROM console 
			WHERE id = \$1`).
		WithArgs(result.ID.String()).
		WillReturnRows(
			pgxmock.NewRows([]string{"id", "name", "manufacturer", "release_date"}).
				AddRow(result.ID, result.Name, result.Manufacturer, result.ReleaseDate),
		).
		Times(1)

	repo := repository.NewConsole(db)

	console, err := repo.FindByID(context.Background(), result.ID.String())
	require.NoError(t, err)

	assert.Equal(t, result, console)
}

func TestConsoleSave(t *testing.T) {
	db, conn, teardown := setupTestDb(t)
	defer teardown()

	releaseDate, err := time.Parse(time.DateOnly, "2020-11-12")
	require.NoError(t, err)

	entity := &entity.Console{
		ID:           uuid.New(),
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  releaseDate,
	}

	conn.
		ExpectExec(`
			INSERT INTO console \(id, name, manufacturer, release_date\) 
			VALUES \(\$1, \$2, \$3, \$4\)`).
		WithArgs(
			entity.ID,
			entity.Name,
			entity.Manufacturer,
			entity.ReleaseDate,
		).
		WillReturnResult(pgxmock.NewResult("INSERT", 1)).
		Times(1)

	repo := repository.NewConsole(db)

	err = repo.Save(context.Background(), entity)
	assert.NoError(t, err)
}

func TestConsoleSave_NoneAffected(t *testing.T) {
	db, conn, teardown := setupTestDb(t)
	defer teardown()

	conn.
		ExpectExec(`
			INSERT INTO console \(id, name, manufacturer, release_date\) 
			VALUES \(\$1, \$2, \$3, \$4\)`).
		WithArgs(
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
		).
		WillReturnResult(pgxmock.NewResult("INSERT", 0)).
		Times(1)

	repo := repository.NewConsole(db)

	err := repo.Save(context.Background(), &entity.Console{})

	assert.EqualError(t, err, "expected to affect 1 row, affected 0")
}

func TestConsoleSave_MultipleAffected(t *testing.T) {
	db, conn, teardown := setupTestDb(t)
	defer teardown()

	conn.
		ExpectExec(`
			INSERT INTO console \(id, name, manufacturer, release_date\) 
			VALUES \(\$1, \$2, \$3, \$4\)`).
		WithArgs(
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
		).
		WillReturnResult(pgxmock.NewResult("INSERT", 2)).
		Times(1)

	repo := repository.NewConsole(db)

	err := repo.Save(context.Background(), &entity.Console{})

	assert.EqualError(t, err, "expected to affect 1 row, affected 2")
}
