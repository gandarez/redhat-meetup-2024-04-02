package igdb_test

import (
	"context"
	"net/http"
	"net/http/httptest"
)

func setupTestServer() (string, *http.ServeMux, func()) {
	router := http.NewServeMux()
	srv := httptest.NewServer(router)

	return srv.URL, router, func() {
		srv.Close()
	}
}

type mockTwitchClient struct {
	AuthenticateFn      func(ctx context.Context) (string, error)
	AuthenticateFnCount int
	ClientIDFn          func() string
	ClientIDFnCount     int
}

func (m *mockTwitchClient) Authenticate(ctx context.Context) (string, error) {
	m.AuthenticateFnCount++
	return m.AuthenticateFn(ctx)
}

func (m *mockTwitchClient) ClientID() string {
	m.ClientIDFnCount++
	return m.ClientIDFn()
}
