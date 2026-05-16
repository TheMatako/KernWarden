package web

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Global declaration of Prometheus metric matrices (Vector Gauges)
var (
	ProbeStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kernwarden_probe_status",
			Help: "Binary health status of the target URL (1 = Healthy, 0 = Critical)",
		},
		[]string{"url"},
	)

	ProbeLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kernwarden_probe_latency_seconds",
			Help: "Response latency of the target URL in seconds",
		},
		[]string{"url"},
	)
)

func init() {
	// Strict registration of metrics into the default Prometheus registry
	prometheus.MustRegister(ProbeStatus)
	prometheus.MustRegister(ProbeLatency)
}

// Result represents the health status of a checked URL.
type Result struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	IsHealthy  bool
}

// CheckHealth performs an active HTTP/HTTPS probe against the target URL
// and updates the Prometheus metrics in RAM.
func CheckHealth(targetURL string) Result {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(targetURL)

	latency := time.Since(start).Seconds()

	if err != nil {
		ProbeStatus.WithLabelValues(targetURL).Set(0)
		ProbeLatency.WithLabelValues(targetURL).Set(latency)

		return Result{
			URL:       targetURL,
			Latency:   time.Since(start),
			IsHealthy: false,
		}
	}

	defer resp.Body.Close()

	isHealthy := resp.StatusCode >= 200 && resp.StatusCode < 300

	if isHealthy {
		ProbeStatus.WithLabelValues(targetURL).Set(1)
	} else {
		ProbeStatus.WithLabelValues(targetURL).Set(0)
	}
	ProbeLatency.WithLabelValues(targetURL).Set(latency)

	return Result{
		URL:        targetURL,
		StatusCode: resp.StatusCode,
		Latency:    time.Since(start),
		IsHealthy:  isHealthy,
	}
}
