package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/model"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestConsoleCreate(t *testing.T) {
	repo, uc := setupConsoleCreateTest(t)

	repo.
		On("Save",
			context.Background(),
			mock.AnythingOfType("*entity.Console"),
		).
		Return(nil)

	console, err := uc.Create(context.Background(), model.ConsoleInsert{
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  "2020-11-12",
	})
	require.NoError(t, err)

	repo.AssertExpectations(t)

	assert.Equal(t, &model.Console{
		ID:           console.ID,
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  "2020-11-12",
	}, console)
}

func TestConsoleCreate_Repository_Err(t *testing.T) {
	repo, uc := setupConsoleCreateTest(t)

	repo.
		On("Save",
			context.Background(),
			mock.AnythingOfType("*entity.Console"),
		).
		Return(errors.New("some error"))

	console, err := uc.Create(context.Background(), model.ConsoleInsert{
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  "2020-11-12",
	})

	repo.AssertExpectations(t)

	assert.Nil(t, console)
	assert.EqualError(t, err, "some error")
}

func setupConsoleCreateTest(t *testing.T) (
	*usecase.MockConsoleRepositoryCreator,
	*usecase.ConsoleCreate) {
	repo := usecase.NewMockConsoleRepositoryCreator(t)

	return repo, usecase.NewConsoleCreate(repo)
}
