package config

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	envconfig "github.com/sethvargo/go-envconfig"
)

type (
	// Configuration contains loaded environment variables.
	Configuration struct {
		ServiceName string `env:"SERVICE_NAME,required"`
		Server      Server `env:",prefix=SERVER_"`
	}

	// Server contains server environment variables.
	Server struct {
		// Port is the port number to listen on.
		Port int `env:"PORT,default=17020"`
	}
)

// Load loads from environment file first. If it fails, then load environment variables to Configuration struct.
func Load(ctx context.Context, fp string) (Configuration, error) {
	err := godotenv.Load(fp)
	if err != nil {
		log.Printf("failed to load environment file: %s", err)
	}

	var c Configuration

	if err := envconfig.Process(ctx, &c); err != nil {
		return Configuration{}, fmt.Errorf("failed to load environment variables: %s", err)
	}

	return c, nil
}
