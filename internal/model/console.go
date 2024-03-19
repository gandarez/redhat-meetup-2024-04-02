package model

import "time"

type (
	// Console represents the console model response.
	Console struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Manufacturer string `json:"manufacturer"`
		ReleaseDate  string `json:"release_date"`
	}

	// ConsoleInsert represents the console model request.
	ConsoleInsert struct {
		Name         string `json:"name"`
		Manufacturer string `json:"manufacturer"`
		ReleaseDate  string `json:"release_date"`
	}
)

// Validate validates the console insert request.
func (c ConsoleInsert) Validate() error {
	var errors ErrorResponse

	if c.Name == "" {
		errors.Errors = append(errors.Errors, "name is required")
	}

	if c.Name != "" && len(c.Name) > 50 {
		errors.Errors = append(errors.Errors, "name is too long (maximum is 50 characters)")
	}

	if c.Manufacturer == "" {
		errors.Errors = append(errors.Errors, "manufacturer is required")
	}

	if c.Manufacturer != "" && len(c.Manufacturer) > 50 {
		errors.Errors = append(errors.Errors, "manufacturer is too long (maximum is 50 characters)")
	}

	if c.ReleaseDate == "" {
		errors.Errors = append(errors.Errors, "release_date is required")
	}

	if c.ReleaseDate != "" {
		if _, err := time.Parse(time.DateOnly, c.ReleaseDate); err != nil {
			errors.Errors = append(errors.Errors, "release_date is invalid")
		}
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}
