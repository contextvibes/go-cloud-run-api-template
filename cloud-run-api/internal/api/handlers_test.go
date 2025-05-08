// internal/api/handlers_test.go
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
	"github.com/stretchr/testify/require"

	"your-module-name/internal/config"
	"your-module-name/internal/models"
)

// testDeps holds dependencies for handler tests. Simplified.
type testDeps struct {
	handler *Handler
}

// newTestDeps creates a Handler for testing. Simplified.
func newTestDeps(t *testing.T, cfg config.Config) testDeps {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := NewHandler(logger, cfg) // No BQ client needed
	return testDeps{handler: handler}
}

// --- Handler Unit Tests ---

func TestHandleHelloWorld(t *testing.T) {
	testConfig := config.Config{ServiceName: "TestService", Port: "8080", ProjectID: "test-project"}
	deps := newTestDeps(t, testConfig)

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rr := httptest.NewRecorder()

	deps.handler.HandleHelloWorld(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code mismatch")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var resp models.HelloWorldResponse
	err := json.NewDecoder(rr.Body).Decode(&resp)
	require.NoError(t, err, "Failed to decode response body")

	assert.Contains(t, resp.Message, "Hello, World from TestService!")
	assert.NotEmpty(t, resp.Timestamp, "Timestamp should not be empty")

	// Test wrong method
	reqWrongMethod := httptest.NewRequest(http.MethodPost, "/hello", nil)
	rrWrongMethod := httptest.NewRecorder()
	deps.handler.HandleHelloWorld(rrWrongMethod, reqWrongMethod)
	assert.Equal(t, http.StatusMethodNotAllowed, rrWrongMethod.Code)
}

func TestHandleEcho(t *testing.T) {
	testConfig := config.Config{ServiceName: "EchoService", Port: "8080", ProjectID: "test-project"}
	deps := newTestDeps(t, testConfig)

	tests := []struct {
		name                 string
		requestBody          string
		expectedStatus       int
		expectedReceivedText string // Added field for clearer assertion
		expectedReplyPart    string
		expectBodyContains   string // For error messages
	}{
		{
			name:                 "Valid Echo",
			requestBody:          `{"text_to_echo": "Testing 123"}`,
			expectedStatus:       http.StatusOK,
			expectedReceivedText: "Testing 123", // Expect this back
			expectedReplyPart:    "received your message: 'Testing 123'",
		},
		{
			name:               "Empty TextToEcho",
			requestBody:        `{"text_to_echo": ""}`,
			expectedStatus:     http.StatusBadRequest,
			expectBodyContains: "text_to_echo is required",
		},
		{
			name:               "Malformed JSON",
			requestBody:        `{"text_to_echo": "valid",`,
			expectedStatus:     http.StatusBadRequest,
			expectBodyContains: "Invalid request payload",
		},
		{
			name:               "Missing text_to_echo field",
			requestBody:        `{}`,
			expectedStatus:     http.StatusBadRequest,
			expectBodyContains: "text_to_echo is required", // TextToEcho defaults to empty string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			deps.handler.HandleEcho(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "Status code mismatch")

			if tt.expectedStatus == http.StatusOK {
				assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
				var resp models.EchoResponse
				err := json.NewDecoder(rr.Body).Decode(&resp)
				require.NoError(t, err, "Failed to decode response body")
				// --- Improved Assertion ---
				assert.Equal(t, tt.expectedReceivedText, resp.ReceivedText, "ReceivedText mismatch")
				// --- End Improvement ---
				assert.Contains(t, resp.Reply, tt.expectedReplyPart)
				assert.NotEmpty(t, resp.Timestamp, "Timestamp should not be empty")
			} else {
				if tt.expectBodyContains != "" {
					bodyStr := rr.Body.String()
					assert.Contains(t, bodyStr, tt.expectBodyContains, "Response body error message mismatch")
				}
			}
		})
	}

	// Test wrong method
	reqWrongMethod := httptest.NewRequest(http.MethodGet, "/echo", nil)
	rrWrongMethod := httptest.NewRecorder()
	deps.handler.HandleEcho(rrWrongMethod, reqWrongMethod)
	assert.Equal(t, http.StatusMethodNotAllowed, rrWrongMethod.Code)
}
