package metrics

import (
	"fmt"
	"net/http"

	"github.com/jesusvico/http-uptime/internal/collector"
	"github.com/jesusvico/http-uptime/internal/endpoint"
	"github.com/jesusvico/http-uptime/internal/metrics/http_response_time"
	"github.com/jesusvico/http-uptime/internal/metrics/http_status"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var collectors []collector.Collector

func Init() {
	// Add all the collectors
	collectors = append(collectors, http_response_time.Collector{})
	collectors = append(collectors, http_status.Collector{})

	// Register all the collectors
	for _, c := range collectors {
		c.Register()
	}
}

func Collect(e endpoint.Endpoint) error {
	// Create the data
	d := collector.NewData(e)

	// Collect the metrics
	for _, c := range collectors {
		if err := c.Collect(d); err != nil {
			return fmt.Errorf("Error collecting data: %v", err)
		}
	}
	return nil
}

func Handler() http.Handler {
	return promhttp.Handler()
}
