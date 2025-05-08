// internal/models/models.go
package models

// HelloWorldResponse defines the structure for a hello world JSON response.
type HelloWorldResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp,omitempty"`
}

// EchoRequest defines a simple structure for a POST request to be echoed.
type EchoRequest struct {
	TextToEcho string `json:"text_to_echo"`
}

// EchoResponse defines the structure for the echo response.
type EchoResponse struct {
	ReceivedText string `json:"received_text"`
	Reply        string `json:"reply"`
	Timestamp    string `json:"timestamp,omitempty"`
}
