package main

import (
	"fmt"

	"github.com/TheMatako/KernWarden/internal/probes/web"
)

func main() {
	target := "https://github.com"
	fmt.Printf("Starting active probe on: %s\n", target)

	result := web.CheckHealth(target)

	if result.IsHealthy {
		fmt.Printf("[OK] target healthy. Status: %d | Latency: %v\n", result.StatusCode, result.Latency)
	} else {
		fmt.Printf("[CRITICAl] Failed. Target Unhealthy. Status: %d | Latency: %v\n", result.StatusCode, result.Latency)
	}
}
