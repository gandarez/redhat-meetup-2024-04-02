package twitch

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/api"
)

type (
	// Cache interface for caching.
	Cache interface {
		Get(ctx context.Context, key string) (string, error)
		Set(ctx context.Context, key string, value any, expiration time.Duration) error
	}

	// Client communicates with the Twitch api.
	Client struct {
		baseURL      string
		cache        Cache
		cacheKey     string
		client       *api.Client
		clientID     string
		clientSecret string
		logger       *slog.Logger
		// doFunc allows api client options to manipulate request/response handling.
		// default function will be set in constructor.
		//
		// wrapping by api options should be performed as follows:
		//
		//	next := c.doFunc
		//	c.doFunc = func(c *Client, req *http.Request) (*http.Response, error) {
		//		// do something
		//		resp, err := next(c, req)
		//		// do more
		//		return resp, err
		//	}
		doFunc func(c *Client, req *http.Request) (*http.Response, error)
	}

	// Config contains celcoin client configurations.
	Config struct {
		BaseURL      string
		Cache        Cache
		CacheKey     string
		ClientID     string
		ClientSecret string
		Logger       *slog.Logger
	}

	// Twitch is the interface for the Twitch client.
	Twitch interface {
		Authenticate(context.Context) (string, error)
	}
)

// NewClient creates a new Client. Any number of Options can be provided.
func NewClient(config Config) *Client {
	apiclient := api.NewClient(api.Config{
		BaseURL: config.BaseURL,
	})

	c := &Client{
		baseURL:      config.BaseURL,
		cache:        config.Cache,
		cacheKey:     config.CacheKey,
		client:       apiclient,
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
		logger:       config.Logger,
		doFunc: func(c *Client, req *http.Request) (*http.Response, error) {
			return c.client.Do(req)
		},
	}

	return c
}

// Do executes c.doFunc(), which in turn allows wrapping c.client.Do() and manipulating
// the request behavior of the api client.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.doFunc(c, req)
}

// ClientID returns the client id.
func (c *Client) ClientID() string {
	return c.clientID
}
