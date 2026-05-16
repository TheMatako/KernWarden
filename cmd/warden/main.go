package main

import (
	"log"
	"net/http"
	"time"

	"github.com/TheMatako/KernWarden/internal/probes/web"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	targetURL := "https://github.com"
	probeInterval := 15 * time.Second

	// 1. Asynchronous execution: Launch the infinite probe loop in a separate Goroutine
	go func() {
		ticker := time.NewTicker(probeInterval)
		defer ticker.Stop()

		// Force an immediate first probe before the ticker's first tick
		log.Printf("Executing initial probe against %s...", targetURL)

		for range ticker.C {
			log.Printf("Probing %s...", targetURL)
			web.CheckHealth(targetURL)
		}
	}()

	// 2. Synchronous execution: Bind the Prometheus handler to the /metrics route
	http.Handle("/metrics", promhttp.Handler())

	port := ":8080"
	log.Printf("Starting KernWarden metrics server on port %s", port)

	// This is a strictly blocking call. If ListenAndServer returns, the server has crashed.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Critical failure: HTTP server crashed: %v", err)
	}
}
