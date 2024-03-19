package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/repository"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/server"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/usecase"

	"github.com/labstack/echo/v4"
)

// SearchConsoleByID returns a console by id.
func SearchConsoleByID(ctx context.Context, logger *slog.Logger, db repository.DatabaseQueryExecutor) server.Route {
	return server.Route{
		Method: "GET",
		Path:   "/consoles/:id",
		Handler: func(c echo.Context) error {
			logger.Info("search console by id")

			id := c.Param("id")

			uc := usecase.NewConsoleSearch(
				repository.NewConsole(db),
			)

			console, err := uc.Search(ctx, id)
			if err != nil {
				return c.JSON(errorHandler(logger, err))
			}

			logger.Info("console found",
				slog.String("id", console.ID),
				slog.String("name", console.Name),
			)

			return c.JSON(http.StatusOK, console)
		},
	}
}
