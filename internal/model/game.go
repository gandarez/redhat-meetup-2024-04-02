package model

// Game represents a game.
type Game struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Genres      []string `json:"genres"`
	Platforms   []string `json:"platforms"`
	ReleaseDate string   `json:"release_date"`
}
