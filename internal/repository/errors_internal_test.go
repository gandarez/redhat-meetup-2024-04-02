package repository

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
)

func TestHandleFindError_NoRows(t *testing.T) {
	err := handleFindError(pgx.ErrNoRows, "console", "123")

	assert.EqualError(t, err, "console with id 123 not found")
}

func TestHandleFindError_GenericError(t *testing.T) {
	err := handleFindError(errors.New("some error"), "console", "123")

	assert.EqualError(t, err, "failed to find console: some error")
}

func TestHandleSaveError_AlreadyExists(t *testing.T) {
	err := handleSaveError(&pgconn.PgError{
		Code: "23505",
	}, "console")

	assert.EqualError(t, err, "console already exists")
}

func TestHandleSaveError(t *testing.T) {
	err := handleSaveError(errors.New("some error"), "console")

	assert.EqualError(t, err, "failed to save console: some error")
}
