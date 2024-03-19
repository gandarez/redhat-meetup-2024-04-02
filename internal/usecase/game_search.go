package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/client/igdb"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/model"
)

type (
	// GameSearch is a use case for searching a game.
	GameSearch struct {
		igdbClient IgdbGameFinder
	}

	// IgdbGameFinder is an interface for finding games from IGDB.
	IgdbGameFinder interface {
		Games(ctx context.Context, name string) ([]igdb.Game, error)
		Genres(ctx context.Context, ids []int) ([]igdb.Genre, error)
		Platforms(ctx context.Context, ids []int) ([]igdb.Platform, error)
	}
)

// NewGameSearch creates a new game search use case.
func NewGameSearch(igdbClient IgdbGameFinder) *GameSearch {
	return &GameSearch{
		igdbClient: igdbClient,
	}
}

// Search searches games by the given name.
func (c *GameSearch) Search(ctx context.Context, name string) ([]model.Game, error) {
	games, err := c.igdbClient.Games(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to search games: %w", err)
	}

	var result []model.Game

	for _, game := range games {
		genres, err := c.igdbClient.Genres(ctx, game.Genres)
		if err != nil {
			return nil, fmt.Errorf("failed to get genres: %w", err)
		}

		platforms, err := c.igdbClient.Platforms(ctx, game.Platforms)
		if err != nil {
			return nil, fmt.Errorf("failed to get platforms: %w", err)
		}

		result = append(result, parseGameIgdb(game, genres, platforms))
	}

	return result, nil
}

func parseGameIgdb(input igdb.Game, genres []igdb.Genre, platforms []igdb.Platform) model.Game {
	var genresOut []string
	for _, genre := range genres {
		genresOut = append(genresOut, genre.Name)
	}

	var platformsOut []string
	for _, platform := range platforms {
		platformsOut = append(platformsOut, platform.Name)
	}

	return model.Game{
		ID:          input.ID,
		Name:        input.Name,
		ReleaseDate: time.Unix(input.FirstReleaseDate, 0).UTC().Format(time.DateOnly),
		Genres:      genresOut,
		Platforms:   platformsOut,
	}
}
