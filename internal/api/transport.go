package api

import (
	"net/http"
	"time"
)

const (
	// DefaultTimeoutSecs is the default timeout used for requests to the api.
	DefaultTimeoutSecs = 60
)

// NewTransport initializes a new http.Transport.
func NewTransport() *http.Transport {
	return &http.Transport{
		ForceAttemptHTTP2:   true,
		MaxConnsPerHost:     1,
		MaxIdleConns:        1,
		MaxIdleConnsPerHost: 1,
		Proxy:               nil,
		TLSHandshakeTimeout: DefaultTimeoutSecs * time.Second,
	}
}

// LazyCreateNewTransport uses the client's Transport if exists, or creates a new one.
func LazyCreateNewTransport(c *Client) *http.Transport {
	if c != nil && c.client != nil && c.client.Transport != nil {
		return c.client.Transport.(*http.Transport).Clone()
	}

	return NewTransport()
}
