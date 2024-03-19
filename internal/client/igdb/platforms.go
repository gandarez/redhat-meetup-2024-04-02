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

const fieldsPlatformsEndpoint = "name"

// Platform represents a IGDB game platform.
type Platform struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Platforms returns a list of platforms that match the given criteria.
func (c *Client) Platforms(ctx context.Context, ids []int) ([]Platform, error) {
	url := c.baseURL + "/platforms"

	payload := fmt.Sprintf("fields %s; where id = (%s);",
		fieldsPlatformsEndpoint,
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

	result, err := ParsePlatformsResponse(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse results from %q: %s", url, err)
	}

	return result, nil
}

// ParsePlatformsResponse parses the given data and returns a list of platforms.
func ParsePlatformsResponse(data []byte) ([]Platform, error) {
	var platforms []Platform

	if err := json.Unmarshal(data, &platforms); err != nil {
		return nil, fmt.Errorf("failed to parse json response body: %s. body: %q", err, data)
	}

	return platforms, nil
}
