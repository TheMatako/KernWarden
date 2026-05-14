package web

import (
	"net/http"
	"time"
)

// Result represents the health status of a checked URL.
type Result struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	IsHealthy  bool
}

// CheckHealth performs an active HTTP Get probe against the target URL.
func CheckHealth(targetURL string) Result {
	// Phase 1: Client instantiation with a strict timeout to prevent goroutine leaks
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now() // Get precise starting time

	// Phase 2: start HTTP GET request
	resp, err := client.Get(targetURL)

	// Layer 3/4 Error Handler (ex: Dead DNS, Timeout...)
	if err != nil {
		return Result{
			URL:       targetURL,
			Latency:   time.Since(start), // time before fail
			IsHealthy: false,
		}
	}

	// Memory hygiene (closing file descriptor)
	defer resp.Body.Close()

	return Result{
		URL:        targetURL,
		StatusCode: resp.StatusCode,
		Latency:    time.Since(start),
		IsHealthy:  resp.StatusCode >= 200 && resp.StatusCode < 300,
	}
}
