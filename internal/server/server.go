package server

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	// DefaultIdleTimeout is the default idle timeout used for http server.
	DefaultIdleTimeout = 30 * time.Second
	// DefaultReadTimeout is the default read timeout used for http server.
	DefaultReadTimeout = 30 * time.Second
	// DefaultWriteTimeout is the default write timeout used for http server.
	DefaultWriteTimeout = 30 * time.Second
	// DefaultMaxHeaderBytes is the default max header used for http server.
	DefaultMaxHeaderBytes = 8190
)

type (
	// Configuration contains the necessary configuration to setup http server.
	Configuration struct {
		Port           int
		IdleTimeout    time.Duration
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
	}

	// Route contains the information for matching against requests.
	Route struct {
		Method      string
		Path        string
		Handler     echo.HandlerFunc
		Middlewares []echo.MiddlewareFunc
	}

	// Server serves http server.
	Server struct {
		addr     string
		provider *echo.Echo
	}
)

// NewWithConfig creates a new server with given configuration.
func NewWithConfig(c Configuration, opts ...Option) Server {
	e := echo.New()

	e.Server.IdleTimeout = c.IdleTimeout
	e.Server.ReadTimeout = c.ReadTimeout
	e.Server.WriteTimeout = c.WriteTimeout
	e.Server.MaxHeaderBytes = c.MaxHeaderBytes

	s := Server{
		addr:     fmt.Sprintf(":%d", c.Port),
		provider: e,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// New creates a new server with default configuration.
func New(port int, opts ...Option) Server {
	return NewWithConfig(Configuration{
		Port:           port,
		IdleTimeout:    DefaultIdleTimeout,
		ReadTimeout:    DefaultReadTimeout,
		WriteTimeout:   DefaultWriteTimeout,
		MaxHeaderBytes: DefaultMaxHeaderBytes,
	}, opts...)
}

// Start starts an HTTP server.
func (s Server) Start() error {
	return s.provider.Start(s.addr)
}

// Shutdown stops the server gracefully.
func (s Server) Shutdown(ctx context.Context) error {
	return s.provider.Shutdown(ctx)
}

// AddRoute registers a new route for an http method and path with matching handler
// in the router with optional route-level middleware.
func (s Server) AddRoute(route Route) {
	s.provider.Add(route.Method, route.Path, route.Handler, route.Middlewares...)
}
