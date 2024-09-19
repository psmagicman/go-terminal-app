package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Load from environment variables", func(t *testing.T) {
		os.Setenv("TEST_API_URL", "https://test.com")
		os.Setenv("TEST_USER_AGENT", "TestAgent/1.0")
		defer os.Unsetenv("TEST_API_URL")
		defer os.Unsetenv("TEST_USER_AGENT")

		cfg, err := LoadConfig("TEST_")
		assert.NoError(t, err)
		assert.Equal(t, "https://test.com", cfg.Get("api_url"))
		assert.Equal(t, "TestAgent/1.0", cfg.Get("user_agent"))
	})

	// Test with non existent keys
	t.Run("No keys with prefix exist", func(t *testing.T) {
		cfg, err := LoadConfig("NON_EXISTENT_")
		assert.NoError(t, err)
		assert.Equal(t, 0, len(cfg.data))
	})

}
