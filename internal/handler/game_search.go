package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/server"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/usecase"

	"github.com/labstack/echo/v4"
)

// SearchGameByName returns a list of games that match the given name.
func SearchGameByName(ctx context.Context, logger *slog.Logger, igdbClient usecase.IgdbGameFinder) server.Route {
	return server.Route{
		Method: "GET",
		Path:   "/games/:name",
		Handler: func(c echo.Context) error {
			logger.Info("search game by name")

			name := c.Param("name")

			uc := usecase.NewGameSearch(igdbClient)

			games, err := uc.Search(ctx, name)
			if err != nil {
				return c.JSON(errorHandler(logger, err))
			}

			logger.Info("games found", slog.Int("count", len(games)))

			return c.JSON(http.StatusOK, games)
		},
	}
}
