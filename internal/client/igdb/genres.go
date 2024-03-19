package igdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const fieldsGenresEndpoint = "name"

// Genre represents a IGDB game genre.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Genres returns a list of genres that match the given criteria.
func (c *Client) Genres(ctx context.Context, ids []int) ([]Genre, error) {
	url := c.baseURL + "/genres"

	payload := fmt.Sprintf("fields %s; where id = (%s);",
		fieldsGenresEndpoint,
		strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"),
	)

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

	result, err := ParseGenresResponse(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse results from %q: %s", url, err)
	}

	return result, nil
}

// ParseGenresResponse parses the given data and returns a list of genres.
func ParseGenresResponse(data []byte) ([]Genre, error) {
	var genres []Genre

	if err := json.Unmarshal(data, &genres); err != nil {
		return nil, fmt.Errorf("failed to parse json response body: %s. body: %q", err, data)
	}

	return genres, nil
}
