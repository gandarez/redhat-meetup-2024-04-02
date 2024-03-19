package twitch_test

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/client/twitch"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticate(t *testing.T) {
	url, router, tearDown := setupTestServer()
	defer tearDown()

	var numCalls int

	router.HandleFunc(
		"/oauth2/token", func(w http.ResponseWriter, req *http.Request) {
			numCalls++

			// check request
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, []string{"application/json"}, req.Header["Accept"])

			require.Len(t, req.Header["Content-Type"], 1)
			assert.Contains(t, req.Header["Content-Type"][0], "multipart/form-data; boundary")

			// write response
			f, err := os.Open("testdata/twitch_authenticate_response.json")
			require.NoError(t, err)

			w.WriteHeader(http.StatusOK)
			_, err = io.Copy(w, f)
			require.NoError(t, err)
		})

	client := twitch.NewClient(twitch.Config{
		BaseURL: url,
		Cache: &mockCache{
			GetFn: func(_ context.Context, _ string) (string, error) {
				return "", nil
			},
			SetFn: func(_ context.Context, _ string, _ any, _ time.Duration) error {
				return nil
			},
		},
		Logger: slog.Default(),
	})

	token, err := client.Authenticate(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "access-token", token)
	assert.Eventually(t, func() bool { return numCalls == 1 }, time.Second, 50*time.Millisecond)
}

func TestAuthenticate_Cached(t *testing.T) {
	url, _, tearDown := setupTestServer()
	defer tearDown()

	client := twitch.NewClient(twitch.Config{
		BaseURL: url,
		Cache: &mockCache{
			GetFn: func(_ context.Context, _ string) (string, error) {
				return "access-token", nil
			},
			SetFn: func(_ context.Context, _ string, _ any, _ time.Duration) error {
				return nil
			},
		},
	})

	token, err := client.Authenticate(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "access-token", token)
}

func TestAuthenticate_MalformedRespnse(t *testing.T) {
	url, router, tearDown := setupTestServer()
	defer tearDown()

	var numCalls int

	router.HandleFunc(
		"/oauth2/token", func(w http.ResponseWriter, req *http.Request) {
			numCalls++

			// check request
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, []string{"application/json"}, req.Header["Accept"])

			require.Len(t, req.Header["Content-Type"], 1)
			assert.Contains(t, req.Header["Content-Type"][0], "multipart/form-data; boundary")

			// write response
			f, err := os.Open("testdata/twitch_authenticate_bad_response.json")
			require.NoError(t, err)

			w.WriteHeader(http.StatusOK)
			_, err = io.Copy(w, f)
			require.NoError(t, err)
		})

	client := twitch.NewClient(twitch.Config{
		BaseURL: url,
		Cache: &mockCache{
			GetFn: func(_ context.Context, _ string) (string, error) {
				return "", nil
			},
			SetFn: func(_ context.Context, _ string, _ any, _ time.Duration) error {
				return nil
			},
		},
	})

	_, err := client.Authenticate(context.Background())

	assert.EqualError(
		t,
		err,
		`failed to parse results from: failed to parse json response body: invalid character 'i' looking for beginning of value. body: "invalid"`, // nolint:revive
	)
}

func setupTestServer() (string, *http.ServeMux, func()) {
	router := http.NewServeMux()
	srv := httptest.NewServer(router)

	return srv.URL, router, func() { srv.Close() }
}

type mockCache struct {
	GetFn      func(ctx context.Context, key string) (string, error)
	GetFnCount int
	SetFn      func(ctx context.Context, key string, value any, expiration time.Duration) error
	SetFnCount int
}

func (m *mockCache) Get(ctx context.Context, key string) (string, error) {
	m.GetFnCount++
	return m.GetFn(ctx, key)
}

func (m *mockCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	m.SetFnCount++
	return m.SetFn(ctx, key, value, expiration)
}
