// internal/api/server_test.go
package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock" // Removed

	"your-module-name/internal/config"
	"your-module-name/internal/models"
)

func TestSetupRoutes(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	// Ensure ProjectID is set for config loading, as it's checked by Load()
	// For server tests, the specific value isn't critical unless a handler uses it directly.
	cfg := config.Config{ProjectID: "test-project-server", ServiceName: "TestServer"}
	handler := NewHandler(logger, cfg)
	router := SetupRoutes(handler)

	t.Run("HealthCheck Endpoint", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "text/plain", rr.Header().Get("Content-Type"))
		assert.Equal(t, "ok\n", rr.Body.String())

		reqPost := httptest.NewRequest(http.MethodPost, "/healthz", nil)
		rrPost := httptest.NewRecorder()
		router.ServeHTTP(rrPost, reqPost)
		assert.Equal(t, http.StatusMethodNotAllowed, rrPost.Code)
	})

	t.Run("Root Endpoint", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "text/plain", rr.Header().Get("Content-Type"))
		assert.Contains(t, rr.Body.String(), "Welcome to the Go Hello World API!")
	})

	t.Run("NonExistent Path", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/nonexistentpath", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("Hello Endpoint GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		var resp models.HelloWorldResponse
		err := json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Contains(t, resp.Message, "Hello, World from TestServer!")
	})

	t.Run("Hello Endpoint POST Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/hello", bytes.NewBufferString(`{}`))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	})

	t.Run("Echo Endpoint POST", func(t *testing.T) {
		echoPayload := `{"text_to_echo": "Live Test"}`
		req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(echoPayload))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		var resp models.EchoResponse
		err := json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "Live Test", resp.ReceivedText)
		assert.Contains(t, resp.Reply, "received your message: 'Live Test'")
	})

	t.Run("Echo Endpoint GET Not Allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/echo", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	})
}
