package redis

import (
	"context"
	"fmt"
	"time"

	redisClient "github.com/redis/go-redis/v9"
)

type (
	// Client is the redis client.
	Client struct {
		rc *redisClient.Client
	}

	// Configuration contains redis client configurations.
	Configuration struct {
		Addr     string
		Password string
		DB       int
	}
)

// NewClient creates a new redis client.
func NewClient(cfg Configuration) *Client {
	c := redisClient.NewClient(&redisClient.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Client{c}
}

// Set saves a key/value pair on cache.
func (c *Client) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	if err := c.rc.Set(ctx, key, value, expiration); err != nil {
		return fmt.Errorf("failed to set cache: %s", err.Err())
	}

	return nil
}

// Get retrieves a value from cache.
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rc.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get from cache: %w", err)
	}

	return val, nil
}
