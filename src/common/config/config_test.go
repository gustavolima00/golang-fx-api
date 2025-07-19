package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("PORT", "9000")

	cfg := NewConfig()

	assert.NotNil(t, cfg)
	assert.Equal(t, "9000", cfg.Port)

	// Clean up environment variables
	os.Unsetenv("PORT")
}
