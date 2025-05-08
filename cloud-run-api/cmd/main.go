// cmd/main.go
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"your-module-name/internal/api"
	"github.com/duizendstra/dui-go/logging/cloudlogging" // New import
	"your-module-name/internal/config"
)

// Package-level variables for application config and the logger.
var (
	appConfig config.Config
	logger    *slog.Logger
)

// init runs before main and is used for essential setup like loading configuration
// and initializing the global logger.
func init() {
	var err error
	appConfig, err = config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: Config load error: %v\n", err)
		os.Exit(1)
	}

	// Initialize the structured logger with the dui-go CloudLoggingHandler.
	// The NewCloudLoggingHandler from dui-go is expected to return an slog.Handler.
	// It will internally handle LOG_LEVEL and Project ID detection if designed that way.
	cloudHandler := cloudlogging.NewCloudLoggingHandler(appConfig.ServiceName)
	logger = slog.New(cloudHandler) // Pass the handler directly

	slog.SetDefault(logger)
	logger.Debug("Configuration loaded successfully")
	logger.Debug("Initialization complete.")
}

// main is the entry point of the application.
func main() {
	logger.Info(fmt.Sprintf("%s starting...", appConfig.ServiceName))

	apiHandler := api.NewHandler(logger, appConfig)
	httpHandler := api.SetupRoutes(apiHandler)

	addr := ":" + appConfig.Port
	logger.Info("Server listening", "address", addr)

	server := &http.Server{
		Addr:              addr,
		Handler:           httpHandler,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
