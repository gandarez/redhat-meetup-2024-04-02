package igdb_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/client/igdb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Games(t *testing.T) {
	url, router, tearDown := setupTestServer()
	defer tearDown()

	var numCalls int

	router.HandleFunc(
		"/games", func(w http.ResponseWriter, req *http.Request) {
			numCalls++

			// check request
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, []string{"application/json"}, req.Header["Accept"])
			assert.Equal(t, []string{"application/json"}, req.Header["Content-Type"])

			// check body
			body, err := io.ReadAll(req.Body)
			require.NoError(t, err)

			assert.Equal(t, "fields first_release_date,name,slug,summary,genres,platforms; search \"Mario\";", string(body))

			// write response
			f, err := os.Open("testdata/igdb_games_response.json")
			require.NoError(t, err)

			w.WriteHeader(http.StatusOK)
			_, err = io.Copy(w, f)
			require.NoError(t, err)
		})

	c := igdb.NewClient(igdb.Config{
		BaseURL: url,
		TwitchClient: &mockTwitchClient{
			AuthenticateFn: func(_ context.Context) (string, error) {
				return "access-token", nil
			},
			ClientIDFn: func() string {
				return "client-id"
			},
		},
	})

	games, err := c.Games(context.Background(), "Mario")
	require.NoError(t, err)

	assert.Len(t, games, 10)

	assert.Eventually(t, func() bool { return numCalls == 1 }, time.Second, 50*time.Millisecond)
}
