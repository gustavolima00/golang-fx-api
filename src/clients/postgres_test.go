package clients

import (
	"go-api/src/common/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// MockGormDB is a mock implementation of *gorm.DB for testing purposes.
type MockGormDB struct {
	gorm.DB
}

func TestNewPostgresClient(t *testing.T) {
	t.Run("should return error with invalid config (e.g., empty host)", func(t *testing.T) {
		cfg := &config.Config{
			PostgresHost:     "", // Invalid host
			PostgresUser:     "user",
			PostgresPassword: "password",
			PostgresDB:       "testdb",
			PostgresPort:     "5432",
		}

		client, err := NewPostgresClient(PostgresClientParams{Config: cfg})

		assert.Error(t, err)
		assert.Nil(t, client)
	})
}
