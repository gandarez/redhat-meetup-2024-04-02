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

func TestLoad_DatabaseHost(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("DATABASE_HOST")

	assert.Equal(t, "localhost", c.Database.Host)
}

func TestLoad_DatabaseUser(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("DATABASE_USER")

	assert.Equal(t, "username", c.Database.User)
}

func TestLoad_DatabasePassword(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("DATABASE_PASSWORD")

	assert.Equal(t, "password", c.Database.Password)
}

func TestLoad_DatabasePort(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("DATABASE_PORT")

	assert.Equal(t, 9999, c.Database.Port)
}

func TestLoad_DatabaseName(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("DATABASE_NAME")

	assert.Equal(t, "database_name", c.Database.Name)
}

func TestLoad_RedisHost(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("REDIS_HOST")

	assert.Equal(t, "localhost:6379", c.Redis.Host)
}

func TestLoad_RedisPassword(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("REDIS_PASSWORD")

	assert.Equal(t, "secret", c.Redis.Password)
}

func TestLoad_RedisDB(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("REDIS_DB")

	assert.Equal(t, 9, c.Redis.DB)
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

func TestLoad_VendorIGDBHost(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("VENDOR_IGDB_HOST")

	assert.Equal(t, "some-api.igdb.com", c.VendorIGDB.Host)
}

func TestLoad_VendorTwitchHost(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("VENDOR_TWITCH_HOST")

	assert.Equal(t, "some-api.twitch.com", c.VendorTwitch.Host)
}

func TestLoad_VendorTwitchClientID(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("VENDOR_TWITCH_CLIENT_ID")

	assert.Equal(t, "some-client-id", c.VendorTwitch.ClientID)
}

func TestLoad_VendorTwitchClientSecret(t *testing.T) {
	c, err := config.Load(context.Background(), "testdata/env")
	require.NoError(t, err)

	defer os.Unsetenv("VENDOR_TWITCH_CLIENT_SECRET")

	assert.Equal(t, "some-client-secret", c.VendorTwitch.ClientSecret)
}
