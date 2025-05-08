// internal/config/config.go
package config

import (
	"fmt"
	// "os" // No longer directly needed for Getenv

	"github.com/duizendstra/dui-go/env" // Use your library
)

// Config holds application configuration values loaded from the environment.
// Struct tags define the corresponding environment variables, defaults, and requirements.
type Config struct {
	ServiceName string `env:"API_SERVICE_NAME" envDefault:"go-hello-world-api"`
	Port        string `env:"PORT" envDefault:"8080"`
	// GOOGLE_CLOUD_PROJECT is needed by the cloudlogging library internally,
	// but env.Process doesn't strictly need to load it into *this* struct
	// unless other parts of *your* application code need it directly.
	// If only the logger needs it, we might not need it here.
	// Let's assume for now your app *might* need it elsewhere, or for clarity.
	ProjectID string `env:"GOOGLE_CLOUD_PROJECT" envRequired:"true"`
}

// Load configuration from environment variables using the dui-go/env library.
func Load() (Config, error) {
	var cfg Config
	err := env.Process(&cfg) // Use the new Process function
	if err != nil {
		// Wrap the error for more context, preserving the original error type if possible
		return Config{}, fmt.Errorf("failed to load config from environment: %w", err)
	}

	// Add any custom cross-field validation here if needed after loading
	// e.g., if cfg.Port had to be within a certain range (though it's a string here).

	return cfg, nil
}
