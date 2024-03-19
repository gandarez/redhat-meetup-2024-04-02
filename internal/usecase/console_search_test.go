package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/entity"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/model"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConsoleSearch(t *testing.T) {
	repo, uc := setupConsoleSearchTest(t)

	id := uuid.MustParse("946d31b1-bcf7-4ebe-9790-b152f0a511ca")
	releaseDate, err := time.Parse(time.DateOnly, "2020-11-12")
	require.NoError(t, err)

	repo.
		On("FindByID",
			context.Background(),
			id.String(),
		).
		Return(&entity.Console{
			ID:           id,
			Name:         "PlayStation 5",
			Manufacturer: "Sony",
			ReleaseDate:  releaseDate,
		}, nil)

	console, err := uc.Search(context.Background(), id.String())
	require.NoError(t, err)

	repo.AssertExpectations(t)

	assert.Equal(t, &model.Console{
		ID:           id.String(),
		Name:         "PlayStation 5",
		Manufacturer: "Sony",
		ReleaseDate:  "2020-11-12",
	}, console)
}

func TestConsoleSearch_InvalidID(t *testing.T) {
	repo, uc := setupConsoleSearchTest(t)

	console, err := uc.Search(context.Background(), "123")

	repo.AssertExpectations(t)

	assert.Nil(t, console)
	assert.EqualError(t, err, "invalid id: 123")
}

func TestConsoleSearch_Repository_Err(t *testing.T) {
	repo, uc := setupConsoleSearchTest(t)

	id := uuid.MustParse("946d31b1-bcf7-4ebe-9790-b152f0a511ca")

	repo.
		On("FindByID",
			context.Background(),
			id.String(),
		).
		Return(nil, errors.New("some error"))

	console, err := uc.Search(context.Background(), id.String())

	repo.AssertExpectations(t)

	assert.Nil(t, console)
	assert.EqualError(t, err, "some error")
}

func setupConsoleSearchTest(t *testing.T) (
	*usecase.MockConsoleRepositoryFinder,
	*usecase.ConsoleSearch) {
	repo := usecase.NewMockConsoleRepositoryFinder(t)

	return repo, usecase.NewConsoleSearch(repo)
}
