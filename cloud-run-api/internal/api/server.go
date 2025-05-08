// internal/api/server.go
package api

import (
	"fmt"
	"net/http"

	"github.com/duizendstra/dui-go/logging/cloudlogging" // New import
)

// SetupRoutes configures the HTTP routes and returns the handler.
func SetupRoutes(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	// Hello World GET handler
	helloHandlerFunc := http.HandlerFunc(handler.HandleHelloWorld)
	// Use WithCloudTraceContext from the new package
	handlerWithTraceHello := cloudlogging.WithCloudTraceContext(helloHandlerFunc)
	mux.Handle("/hello", handlerWithTraceHello)

	// Echo POST handler
	echoHandlerFunc := http.HandlerFunc(handler.HandleEcho)
	// Use WithCloudTraceContext from the new package
	handlerWithTraceEcho := cloudlogging.WithCloudTraceContext(echoHandlerFunc)
	mux.Handle("/echo", handlerWithTraceEcho)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Welcome to the Go Hello World API!")
		fmt.Fprintln(w, "Try /hello (GET) or /echo (POST)")
	})

	return mux
}
