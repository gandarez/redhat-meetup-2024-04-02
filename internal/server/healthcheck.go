package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ReadinessRoute checks if the service and its dependencies are healthy.
func ReadinessRoute() Route {
	return Route{
		Method: "GET",
		Path:   "/readiness",
		Handler: func(c echo.Context) error {
			return c.String(http.StatusOK, "Ok")
		},
	}
}

// LivenessRoute checks if the service is alive.
func LivenessRoute() Route {
	return Route{
		Method: "GET",
		Path:   "/liveness",
		Handler: func(c echo.Context) error {
			return c.String(http.StatusOK, "Ok")
		},
	}
}
