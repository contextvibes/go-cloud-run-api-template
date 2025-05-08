// internal/api/handlers.go
package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	// "reflect" // No longer needed
	// "strings"   // No longer needed for dataset name
	"time"

	// "cloud.google.com/go/bigquery" // No longer needed

	"your-module-name/internal/config"
	"your-module-name/internal/models" // Keep for our new models
)

// Handler holds dependencies required by the HTTP handlers.
type Handler struct {
	Logger    *slog.Logger
	AppConfig config.Config
	// BQClient BQClientInterface // Removed
	// SchemaTypeMap map[string]reflect.Type // Removed
}

// NewHandler creates and returns a new Handler instance with its dependencies initialized.
func NewHandler(logger *slog.Logger, appConfig config.Config) *Handler { // Removed bqClient
	return &Handler{
		Logger:    logger,
		AppConfig: appConfig,
	}
}

// HandleHelloWorld is a simple GET handler.
func (h *Handler) HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	h.Logger.InfoContext(ctx, "Hello world request received", "path", r.URL.Path)

	response := models.HelloWorldResponse{
		Message:   fmt.Sprintf("Hello, World from %s!", h.AppConfig.ServiceName),
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.ErrorContext(ctx, "Failed to encode hello world response", "error", err)
		// Hard to send an error to client if headers already sent and partially written.
	}
}

// HandleEcho is a POST handler that echoes back part of the request.
func (h *Handler) HandleEcho(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	h.Logger.InfoContext(ctx, "Echo request received", "path", r.URL.Path)

	var echoReq models.EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&echoReq); err != nil {
		h.Logger.ErrorContext(ctx, "Failed to decode echo request body", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if echoReq.TextToEcho == "" {
		h.Logger.WarnContext(ctx, "Echo request with empty text_to_echo")
		http.Error(w, "text_to_echo is required", http.StatusBadRequest)
		return
	}

	response := models.EchoResponse{
		ReceivedText: echoReq.TextToEcho,
		Reply:        fmt.Sprintf("Service '%s' received your message: '%s'", h.AppConfig.ServiceName, echoReq.TextToEcho),
		Timestamp:    time.Now().UTC().Format(time.RFC3339Nano),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.ErrorContext(ctx, "Failed to encode echo response", "error", err)
	}
}
