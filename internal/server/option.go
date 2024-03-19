package server

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Option is a functional option for Server.
type Option func(*Server)

// WithRecover recovers from panic.
func WithRecover(logger *slog.Logger) Option {
	return func(s *Server) {
		s.provider.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			Skipper:           middleware.DefaultSkipper,
			StackSize:         4 << 10, // 4 KB
			DisableStackAll:   false,
			DisablePrintStack: false,
			LogLevel:          4, // Error
			LogErrorFunc: func(_ echo.Context, err error, stack []byte) error {
				logger.Error("[panic recover]",
					slog.Any("error", err),
					slog.String("stack", string(stack)),
				)

				return nil
			},
		}))
	}
}

// WithGzip compresses HTTP response using gzip compression scheme.
func WithGzip() Option {
	return func(s *Server) {
		s.provider.Use(middleware.Gzip())
	}
}

// WithDecompress decompresses request body based if content encoding type is set to "gzip" with default config.
func WithDecompress() Option {
	return func(s *Server) {
		s.provider.Use(middleware.Decompress())
	}
}
