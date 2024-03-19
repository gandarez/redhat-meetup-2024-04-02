package igdb

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/api"
)

type (
	// Client communicates with the igdb api.
	Client struct {
		baseURL      string
		client       *api.Client
		logger       *slog.Logger
		twitchClient TwitchClient
		// doFunc allows api client options to manipulate request/response handling.
		// default function will be set in constructor.
		//
		// wrapping by api options should be performed as follows:
		//
		//	next := c.doFunc
		//	c.doFunc = func(ctx context.Context, c *Client, req *http.Request) (*http.Response, error) {
		//		// do something
		//		resp, err := next(ctx, c, req)
		//		// do more
		//		return resp, err
		//	}
		doFunc func(ctx context.Context, c *Client, req *http.Request) (*http.Response, error)
	}

	// Config contains igdb client configurations.
	Config struct {
		BaseURL      string
		Logger       *slog.Logger
		TwitchClient TwitchClient
	}

	// TwitchClient is the interface for the Twitch client.
	TwitchClient interface {
		Authenticate(ctx context.Context) (string, error)
		ClientID() string
	}
)

// NewClient creates a new igdb Client.
func NewClient(config Config) *Client {
	apiclient := api.NewClient(api.Config{
		BaseURL: config.BaseURL,
	})

	c := &Client{
		baseURL:      config.BaseURL,
		client:       apiclient,
		logger:       config.Logger,
		twitchClient: config.TwitchClient,
		doFunc: func(ctx context.Context, c *Client, req *http.Request) (*http.Response, error) {
			token, err := c.twitchClient.Authenticate(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get authorization token: %w", err)
			}

			req.Header.Add("Client-ID", c.twitchClient.ClientID())
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

			return c.client.Do(req)
		},
	}

	return c
}

// Do executes c.doFunc(), which in turn allows wrapping c.client.Do() and manipulating
// the request behavior of the api client.
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return c.doFunc(ctx, c, req)
}
