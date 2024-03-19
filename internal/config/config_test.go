package config_test

import (
	"context"
	"os"
	"testing"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad_Err(t *testing.T) {
	_, err := config.Load(context.Background(), "testdata/env-empty")

	assert.EqualError(t, err, "failed to load environment variables: ServiceName: missing required value: SERVICE_NAME")
}

func TestLoad_RequiredFields(t *testing.T) {
	_, err := config.Load(context.Background(), "testdata/env-minimal")

	assert.NoError(t, err)
}

func TestLoad_ServiceName(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("SERVICE_NAME")

	assert.Equal(t, "some-service", c.ServiceName)
}

func TestLoad_ServerPort(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("SERVER_PORT")

	assert.Equal(t, 8081, c.Server.Port)
}

func TestLoad_ServerPort_Default(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env-minimal")
	require.NoError(t, err)

	defer os.Unsetenv("SERVER_PORT")

	assert.Equal(t, 17020, c.Server.Port)
}
