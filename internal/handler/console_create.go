package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/model"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/repository"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/server"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/usecase"

	"github.com/labstack/echo/v4"
)

// CreateConsole creates a new console.
func CreateConsole(ctx context.Context, logger *slog.Logger, db repository.DatabaseQueryExecutor) server.Route {
	return server.Route{
		Method: "POST",
		Path:   "/consoles",
		Handler: func(c echo.Context) error {
			logger.Info("create console")

			var body model.ConsoleInsert

			if err := c.Bind(&body); err != nil {
				return c.JSON(errorHandler(logger, err))
			}

			uc := usecase.NewConsoleCreate(
				repository.NewConsole(db),
			)

			console, err := uc.Create(ctx, body)
			if err != nil {
				return c.JSON(errorHandler(logger, err))
			}

			logger.Info("successfully created console",
				slog.String("id", console.ID),
				slog.String("name", console.Name),
			)

			return c.JSON(http.StatusCreated, console)
		},
	}
}
