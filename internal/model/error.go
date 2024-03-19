package model

import (
	"encoding/json"
	"strings"
)

// ErrorResponse is the default error response.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Error implements the error interface.
func (e ErrorResponse) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return strings.Join(e.Errors, ", ")
	}

	return string(data)
}
