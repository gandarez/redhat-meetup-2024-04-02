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
		ServiceName string   `env:"SERVICE_NAME,required"`
		Database    Database `env:",prefix=DATABASE_"`
		Server      Server   `env:",prefix=SERVER_"`
	}

	// Database contains database environment variables.
	Database struct {
		Host     string `env:"HOST,required"`
		Name     string `env:"NAME,required"`
		User     string `env:"USER,required"`
		Password string `env:"PASSWORD,required"`
		Port     int    `env:"PORT,default=5432"`
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
