// tests/integration/main_api_test.go
package integration_test

import (
	"bytes"
	// "context" // Not directly needed for logger setup
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	// "time" // Not strictly needed here

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"your-module-name/internal/api"
	"github.com/duizendstra/dui-go/logging/cloudlogging" // New import
	"your-module-name/internal/config"
	"your-module-name/internal/models"
)

var (
	testServer *httptest.Server
	appConfig  config.Config
	logger     *slog.Logger // Logger for the test setup/teardown itself
)

// TestMain sets up the HTTP test server once for all integration tests in this package.
func TestMain(m *testing.M) {
	var err error
	appConfig, err = config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: Failed to load config for integration tests: %v\n", err)
		os.Exit(1)
	}
	appConfig.ServiceName = "IntegrationTestService" // Override for test clarity

	logOutput := io.Discard
	// For the test's own logger, a simple TextHandler is fine.
	// The actual application handler will use the dui-go cloud logger.
	if os.Getenv("INTEGRATION_TEST_LOG_OUTPUT") == "stderr" {
		logOutput = os.Stderr
	}
	logger = slog.New(slog.NewTextHandler(logOutput, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Setup the handler and router
	// The API handler will be configured with the dui-go logger
	// The dui-go NewCloudLoggingHandler should handle its own LOG_LEVEL env var processing.
	handlerLogger := slog.New(cloudlogging.NewCloudLoggingHandler(appConfig.ServiceName))
	apiHandler := api.NewHandler(handlerLogger, appConfig)
	httpHandler := api.SetupRoutes(apiHandler) // SetupRoutes uses dui-go's WithCloudTraceContext
	testServer = httptest.NewServer(httpHandler)

	logger.Info("Integration test server started", "url", testServer.URL)

	exitCode := m.Run()

	logger.Info("Stopping integration test server...")
	testServer.Close()
	logger.Info("Integration test server stopped.")
	os.Exit(exitCode)
}

func TestIntegration_APIEndpoints(t *testing.T) {
	require.NotNil(t, testServer, "Test server should be initialized")

	client := testServer.Client()

	t.Run("GET /healthz", func(t *testing.T) {
		resp, err := client.Get(testServer.URL + "/healthz")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, "ok\n", string(bodyBytes))
	})

	t.Run("GET /hello", func(t *testing.T) {
		resp, err := client.Get(testServer.URL + "/hello")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		var helloResp models.HelloWorldResponse
		err = json.NewDecoder(resp.Body).Decode(&helloResp)
		require.NoError(t, err)
		assert.Contains(t, helloResp.Message, "Hello, World from "+appConfig.ServiceName)
		assert.NotEmpty(t, helloResp.Timestamp)
	})

	t.Run("POST /echo", func(t *testing.T) {
		echoPayload := models.EchoRequest{TextToEcho: "Integration Echo Test"}
		payloadBytes, err := json.Marshal(echoPayload)
		require.NoError(t, err)

		httpResp, err := client.Post(testServer.URL+"/echo", "application/json", bytes.NewBuffer(payloadBytes))
		require.NoError(t, err)
		defer httpResp.Body.Close()

		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		assert.Equal(t, "application/json", httpResp.Header.Get("Content-Type"))

		var echoResp models.EchoResponse
		err = json.NewDecoder(httpResp.Body).Decode(&echoResp)
		require.NoError(t, err)
		assert.Equal(t, "Integration Echo Test", echoResp.ReceivedText)
		assert.Contains(t, echoResp.Reply, "received your message: 'Integration Echo Test'")
		assert.NotEmpty(t, echoResp.Timestamp)
	})

	t.Run("POST /echo - Empty Text", func(t *testing.T) {
		echoPayload := models.EchoRequest{TextToEcho: ""}
		payloadBytes, err := json.Marshal(echoPayload)
		require.NoError(t, err)

		httpResp, err := client.Post(testServer.URL+"/echo", "application/json", bytes.NewBuffer(payloadBytes))
		require.NoError(t, err)
		defer httpResp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, httpResp.StatusCode)
		bodyBytes, _ := io.ReadAll(httpResp.Body)
		assert.Contains(t, string(bodyBytes), "text_to_echo is required")
	})

	t.Run("GET /", func(t *testing.T) {
		resp, err := client.Get(testServer.URL + "/")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(bodyBytes), "Welcome to the Go Hello World API!")
	})

	t.Run("GET /notfound", func(t *testing.T) {
		resp, err := client.Get(testServer.URL + "/notfound")
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
