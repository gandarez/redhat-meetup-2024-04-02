package api

import (
	"net/http"
)

type (
	// Client communicates with the any api.
	Client struct {
		baseURL string
		client  *http.Client
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

	// Config contains api client configurations.
	Config struct {
		BaseURL string
	}
)

// NewClient creates a new Client. Any number of Options can be provided.
func NewClient(config Config) *Client {
	c := &Client{
		baseURL: config.BaseURL,
		client: &http.Client{
			Transport: NewTransport(),
		},
		doFunc: func(c *Client, req *http.Request) (*http.Response, error) {
			req.Header.Set("Accept", "application/json")

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
