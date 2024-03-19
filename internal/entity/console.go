package entity

import (
	"time"

	"github.com/google/uuid"
)

// Console represents a video game console.
type Console struct {
	ID           uuid.UUID
	Name         string
	Manufacturer string
	ReleaseDate  time.Time
}
