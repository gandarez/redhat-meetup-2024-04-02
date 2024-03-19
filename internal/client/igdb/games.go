package igdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const fieldsGamesEndpoint = "first_release_date,name,slug,summary,genres,platforms"

// Game represents a IGDB game.
type Game struct {
	ID               int    `json:"id"`
	FirstReleaseDate int64  `json:"first_release_date"`
	Genres           []int  `json:"genres"`
	Name             string `json:"name"`
	Platforms        []int  `json:"platforms"`
	Slug             string `json:"slug"`
	Summary          string `json:"summary"`
}

// Games returns a list of games that match the given criteria.
func (c *Client) Games(ctx context.Context, criteria string) ([]Game, error) {
	url := c.baseURL + "/games"

	payload := fmt.Sprintf("fields %s; search %q;", fieldsGamesEndpoint, criteria)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed making request to %q: %w", url, err)
	}
	defer resp.Body.Close() // nolint:errcheck,gosec

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body from %q: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"invalid response status from %q. got: %d, want: %d. body: %q",
			url,
			resp.StatusCode,
			http.StatusOK,
			string(body),
		)
	}

	result, err := ParseGamesResponse(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse results from %q: %s", url, err)
	}

	return result, nil
}

// ParseGamesResponse parses the given data and returns a list of games.
func ParseGamesResponse(data []byte) ([]Game, error) {
	var game []Game

	if err := json.Unmarshal(data, &game); err != nil {
		return nil, fmt.Errorf("failed to parse json response body: %s. body: %q", err, data)
	}

	return game, nil
}
