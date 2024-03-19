package twitch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"time"
)

const grantType = "client_credentials"

// AuthResponse contains the response for the authentication endpoint.
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// Authenticate returns an access token for the Twitch API.
func (c *Client) Authenticate(ctx context.Context) (string, error) {
	if token, err := c.cache.Get(ctx, c.cacheKey); err == nil && token != "" {
		return token, nil
	}

	url := c.baseURL + "/oauth2/token"

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	if err := writer.WriteField("client_id", c.clientID); err != nil {
		return "", err
	}

	if err := writer.WriteField("client_secret", c.clientSecret); err != nil {
		return "", err
	}

	if err := writer.WriteField("grant_type", grantType); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := c.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed making request to %q: %s", url, err)
	}
	defer resp.Body.Close() // nolint:errcheck,gosec

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body from %q: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"invalid response status from %q. got: %d, want: %d. body: %q",
			url,
			resp.StatusCode,
			http.StatusOK,
			string(body),
		)
	}

	result, err := ParseAuthenticateResponse(body)
	if err != nil {
		return "", fmt.Errorf("failed to parse results from: %w", err)
	}

	err = c.cache.Set(ctx, c.cacheKey, result.AccessToken, time.Duration(result.ExpiresIn)*time.Second)
	if err != nil {
		c.logger.Error("failed to cache authorization token", slog.Any("error", err))
	} else {
		c.logger.Info("authorization token cached")
	}

	return result.AccessToken, nil
}

// ParseAuthenticateResponse parses the Authenticate response into AuthResponse.
func ParseAuthenticateResponse(data []byte) (*AuthResponse, error) {
	var body AuthResponse

	if err := json.Unmarshal(data, &body); err != nil {
		return nil, fmt.Errorf("failed to parse json response body: %w. body: %q", err, data)
	}

	return &body, nil
}
