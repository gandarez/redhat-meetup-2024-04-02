package model_test

import (
	"testing"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConsoleInsertValidate(t *testing.T) {
	insert := model.ConsoleInsert{
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  "2020-11-12",
	}

	err := insert.Validate()
	assert.NoError(t, err)
}

func TestConsoleInsertValidate_NameRequired(t *testing.T) {
	insert := model.ConsoleInsert{
		Manufacturer: "Sony",
		ReleaseDate:  "2020-11-12",
	}

	err := insert.Validate()

	errResponse, ok := err.(model.ErrorResponse)
	require.True(t, ok)

	require.Len(t, errResponse.Errors, 1)
	assert.Equal(t, "name is required", errResponse.Errors[0])
}

func TestConsoleInsertValidate_NameTooLong(t *testing.T) {
	insert := model.ConsoleInsert{
		Name:         "name too long name too long name too long name too long",
		Manufacturer: "Sony",
		ReleaseDate:  "2020-11-12",
	}

	err := insert.Validate()

	errResponse, ok := err.(model.ErrorResponse)
	require.True(t, ok)

	require.Len(t, errResponse.Errors, 1)
	assert.Equal(t, "name is too long (maximum is 50 characters)", errResponse.Errors[0])
}

func TestConsoleInsertValidate_ManufacturerRequired(t *testing.T) {
	insert := model.ConsoleInsert{
		Name:        "PlayStation 5",
		ReleaseDate: "2020-11-12",
	}

	err := insert.Validate()

	errResponse, ok := err.(model.ErrorResponse)
	require.True(t, ok)

	require.Len(t, errResponse.Errors, 1)
	assert.Equal(t, "manufacturer is required", errResponse.Errors[0])
}

func TestConsoleInsertValidate_ManufacturerTooLong(t *testing.T) {
	insert := model.ConsoleInsert{
		Name:         "PlayStation 5",
		Manufacturer: "manufacturer too long manufacturer too long manufacturer",
		ReleaseDate:  "2020-11-12",
	}

	err := insert.Validate()

	errResponse, ok := err.(model.ErrorResponse)
	require.True(t, ok)

	require.Len(t, errResponse.Errors, 1)
	assert.Equal(t, "manufacturer is too long (maximum is 50 characters)", errResponse.Errors[0])
}

func TestConsoleInsertValidate_ReleaseDateRequired(t *testing.T) {
	insert := model.ConsoleInsert{
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
	}

	err := insert.Validate()

	errResponse, ok := err.(model.ErrorResponse)
	require.True(t, ok)

	require.Len(t, errResponse.Errors, 1)
	assert.Equal(t, "release_date is required", errResponse.Errors[0])
}

func TestConsoleInsertValidate_ReleaseDateInvalid(t *testing.T) {
	insert := model.ConsoleInsert{
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  "invalid date",
	}

	err := insert.Validate()

	errResponse, ok := err.(model.ErrorResponse)
	require.True(t, ok)

	require.Len(t, errResponse.Errors, 1)
	assert.Equal(t, "release_date is invalid", errResponse.Errors[0])
}
