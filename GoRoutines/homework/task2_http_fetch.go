package homework

import (
	"time"
)

// Task 2: HTTP API with Concurrency
//
// OBJECTIVE: Fetch multiple URLs concurrently with timeout
//
// PACKAGES TO USE:
// - net/http: https://pkg.go.dev/net/http
//   client := &http.Client{Timeout: 10 * time.Second}
//   resp, err := client.Get(url)
//   statusCode := resp.StatusCode
//
// - context: https://pkg.go.dev/context
//   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//   defer cancel()
//
// HINT: Launch goroutine per URL, collect results in channel

// FetchURLs fetches multiple URLs concurrently and returns their status codes
func FetchURLs(urls []string, timeout time.Duration) (map[string]int, error) {
	// TODO: Implement concurrent URL fetching
	// 1. Create HTTP client with timeout
	// 2. Create result channel for collecting responses
	// 3. Create WaitGroup for tracking goroutines
	// 4. For each URL, launch goroutine that fetches and sends result
	// 5. Wait for all goroutines to complete
	// 6. Collect results into map of URL to status code
	// 7. Return map and any errors
	return nil, nil
}

// FetchWithRetry fetches a URL with retry logic
func FetchWithRetry(url string, maxRetries int, timeout time.Duration) (int, error) {
	// TODO: Implement fetch with exponential backoff retry
	// 1. Create HTTP client with timeout
	// 2. Try to fetch URL
	// 3. If successful, return status code
	// 4. If error, check if retryable
	// 5. Wait with exponential backoff
	// 6. Retry up to maxRetries times
	// 7. Return final result or error
	return 0, nil
}
