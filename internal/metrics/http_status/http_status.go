package http_status

import (
	"github.com/jesusvico/http-uptime/internal/collector"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct{}

var HTTPStatusCode *prometheus.GaugeVec

// Register the collector
func (c Collector) Register() {
	HTTPStatusCode = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_status_code",
			Help: "Latest HTTP status code for different websites",
		},
		[]string{"name", "url"},
	)

	prometheus.MustRegister(HTTPStatusCode)
}

// Collect the metrics
func (c Collector) Collect(d *collector.CollectorData) error {
	if !d.HasData() {
		// There is no data, so we set the status code to -1
		HTTPStatusCode.WithLabelValues(d.GetName(), d.GetUrl()).Set(-1)
		return nil
	}

	HTTPStatusCode.WithLabelValues(d.GetName(), d.GetUrl()).Set(float64(d.Response.StatusCode))

	return nil
}
