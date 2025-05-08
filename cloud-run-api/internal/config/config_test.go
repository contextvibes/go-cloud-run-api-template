// internal/config/config_test.go
package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setEnvForTest remains the same helper
func setEnvForTest(t *testing.T, key, value string) {
	t.Helper()
	originalValue, wasSet := os.LookupEnv(key)
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatalf("Failed to set environment variable %s: %v", key, err)
	}
	t.Cleanup(func() {
		if wasSet {
			err := os.Setenv(key, originalValue)
			if err != nil {
				t.Errorf("Failed to restore original env var %s: %v", key, err)
			}
		} else {
			err := os.Unsetenv(key)
			if err != nil {
				t.Errorf("Failed to unset env var %s: %v", key, err)
			}
		}
	})
}

func TestLoadConfig(t *testing.T) {

	t.Run("Defaults", func(t *testing.T) {
		// Set required var, unset others to test defaults via env.Process
		setEnvForTest(t, "GOOGLE_CLOUD_PROJECT", "test-project-defaults")
		os.Unsetenv("API_SERVICE_NAME")
		os.Unsetenv("PORT")

		cfg, err := Load() // Load calls env.Process internally
		require.NoError(t, err, "Load() with defaults failed unexpectedly")

		assert.Equal(t, "go-hello-world-api", cfg.ServiceName, "Default ServiceName mismatch")
		assert.Equal(t, "8080", cfg.Port, "Default Port mismatch")
		assert.Equal(t, "test-project-defaults", cfg.ProjectID, "ProjectID mismatch")
	})

	t.Run("Overrides", func(t *testing.T) {
		// Set all vars to custom values
		setEnvForTest(t, "API_SERVICE_NAME", "my-custom-api-override")
		setEnvForTest(t, "PORT", "9999")
		setEnvForTest(t, "GOOGLE_CLOUD_PROJECT", "overridden-project-id")

		cfg, err := Load()
		require.NoError(t, err, "Load() with overrides failed unexpectedly")

		assert.Equal(t, "my-custom-api-override", cfg.ServiceName)
		assert.Equal(t, "9999", cfg.Port)
		assert.Equal(t, "overridden-project-id", cfg.ProjectID)
	})

	t.Run("Missing Required ProjectID", func(t *testing.T) {
		// Ensure the required variable is unset
		originalProjectID, projectIDSet := os.LookupEnv("GOOGLE_CLOUD_PROJECT")
		os.Unsetenv("GOOGLE_CLOUD_PROJECT")
		t.Cleanup(func() {
			if projectIDSet {
				os.Setenv("GOOGLE_CLOUD_PROJECT", originalProjectID)
			} else {
				os.Unsetenv("GOOGLE_CLOUD_PROJECT")
			}
		})

		_, err := Load()
		require.Error(t, err, "Load() succeeded when GOOGLE_CLOUD_PROJECT was missing")
		// Check that the error message originates from the env library's required check
		assert.Contains(t, err.Error(), "required environment variable GOOGLE_CLOUD_PROJECT is not set or is empty", "Error message mismatch")
		// Check for the wrapping text added in our Load function
		assert.True(t, strings.HasPrefix(err.Error(), "failed to load config from environment: "), "Error should be wrapped")
	})
}
