package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	dbmigration "github.com/gandarez/redhat-meetup-2024-04-02/db"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/config"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/database"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/handler"
	"github.com/gandarez/redhat-meetup-2024-04-02/internal/server"
)

func main() {
	ctx := context.Background()

	// Setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Load configuration from env vars.
	cfg, err := config.Load(ctx, "./.env")
	if err != nil {
		logger.Error("failed to load configuration", slog.Any("error", err))

		os.Exit(1)
	}

	// Setup database
	db := database.NewClient(database.Configuration{
		DbName:   cfg.Database.Name,
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Port:     cfg.Database.Port,
	})

	// Open database connection
	if err = db.Open(ctx); err != nil {
		logger.Error("failed to open database connection", slog.Any("error", err))

		os.Exit(1)
	}

	// Run database migrations
	if err = dbmigration.Run(db.ConnectionString, logger); err != nil {
		logger.Error(err.Error())

		os.Exit(1)
	}

	// setup server
	httpserver := server.New(cfg.Server.Port,
		server.WithRecover(logger),
		server.WithDecompress(),
		server.WithGzip(),
	)

	// Add default routes for health check
	httpserver.AddRoute(server.ReadinessRoute())
	httpserver.AddRoute(server.LivenessRoute())

	// add http routes
	httpserver.AddRoute(handler.SearchConsoleByID(ctx, logger, db))
	httpserver.AddRoute(handler.CreateConsole(ctx, logger, db))

	// start httpserver
	go func() {
		if err := httpserver.Start(); err != http.ErrServerClosed {
			logger.Error("failed to start server", slog.Any("error", err))
		}
	}()

	logger.Info("http server started", slog.Int("port", cfg.Server.Port))

	// wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// shutdown http server
	if err := httpserver.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown server", slog.Any("error", err))
	}

	logger.Info("server gracefully stopped")
}
