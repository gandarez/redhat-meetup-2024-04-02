package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/entity"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/model"

	"github.com/google/uuid"
)

//go:generate mockery --name ConsoleRepositoryFinder --structname MockConsoleRepositoryFinder --inpackage --case snake
type (
	// ConsoleRepositoryFinder is an interface for finding a console from the repository.
	ConsoleRepositoryFinder interface {
		FindByID(ctx context.Context, id string) (*entity.Console, error)
	}

	// ConsoleSearch is a use case for searching a console.
	ConsoleSearch struct {
		repo ConsoleRepositoryFinder
	}
)

// NewConsoleSearch creates a new console search use case.
func NewConsoleSearch(repo ConsoleRepositoryFinder) *ConsoleSearch {
	return &ConsoleSearch{
		repo: repo,
	}
}

// Search searches a console.
func (c *ConsoleSearch) Search(ctx context.Context, id string) (*model.Console, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %s", id)
	}

	console, err := c.repo.FindByID(ctx, uid.String())
	if err != nil {
		return nil, err
	}

	return parseConsoleEntity(console), nil
}

func parseConsoleEntity(input *entity.Console) *model.Console {
	return &model.Console{
		ID:           input.ID.String(),
		Name:         input.Name,
		Manufacturer: input.Manufacturer,
		ReleaseDate:  input.ReleaseDate.Format(time.DateOnly),
	}
}
