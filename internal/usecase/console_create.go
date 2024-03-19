package usecase

import (
	"context"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/entity"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/model"

	"github.com/google/uuid"
)

//go:generate mockery --name ConsoleRepositoryCreator --structname MockConsoleRepositoryCreator --inpackage --case snake
type (
	// ConsoleRepositoryCreator is an interface for creating a console in the repository.
	ConsoleRepositoryCreator interface {
		Save(ctx context.Context, console *entity.Console) error
	}

	// ConsoleCreate is a use case for creating a console.
	ConsoleCreate struct {
		repo ConsoleRepositoryCreator
	}
)

// NewConsoleCreate creates a new console create use case.
func NewConsoleCreate(repo ConsoleRepositoryCreator) *ConsoleCreate {
	return &ConsoleCreate{
		repo: repo,
	}
}

// Create creates a console.
func (c *ConsoleCreate) Create(ctx context.Context, input model.ConsoleInsert) (*model.Console, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	console := parseConsoleInsertModel(input)

	if err := c.repo.Save(ctx, console); err != nil {
		return nil, err
	}

	return parseConsoleEntity(console), nil
}

func parseConsoleInsertModel(input model.ConsoleInsert) *entity.Console {
	var releaseDate time.Time

	if input.ReleaseDate != "" {
		releaseDate, _ = time.Parse(time.DateOnly, input.ReleaseDate)
	}

	return &entity.Console{
		ID:           uuid.New(),
		Name:         input.Name,
		Manufacturer: input.Manufacturer,
		ReleaseDate:  releaseDate,
	}
}
