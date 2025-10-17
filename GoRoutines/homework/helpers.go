package homework

import (
	"net/http"
	"time"
)

// MockHTTPClient for testing HTTP operations
// This simulates HTTP responses without making real network calls
type MockHTTPClient struct {
	responses map[string]*http.Response
	delays    map[string]time.Duration
}

// NewMockHTTPClient creates a new mock HTTP client
func NewMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{
		responses: make(map[string]*http.Response),
		delays:    make(map[string]time.Duration),
	}
}

// Get simulates an HTTP GET request
func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	if delay, ok := m.delays[url]; ok {
		time.Sleep(delay)
	}
	if resp, ok := m.responses[url]; ok {
		return resp, nil
	}
	return &http.Response{StatusCode: 200}, nil
}
