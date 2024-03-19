package database_test

import (
	"testing"

	"github.com/gandarez/redhat-meetup-2024-04-02/internal/database"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := database.NewClient(database.Configuration{
		DbName:   "test",
		Host:     "localhost",
		Password: "password",
		Port:     5432,
		User:     "user",
	})
	require.NotNil(t, client)

	assert.Equal(
		t,
		"postgres://user:password@localhost:5432/test?sslmode=disable",
		client.ConnectionString)
}
