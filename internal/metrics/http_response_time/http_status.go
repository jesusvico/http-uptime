package http_response_time

import (
	"github.com/jesusvico/http-uptime/internal/collector"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct{}

var HTTPResposeTime *prometheus.GaugeVec

// Register the collector
func (c Collector) Register() {
	HTTPResposeTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_response_time_seconds",
			Help: "Latest HTTP response time in seconds for different websites",
		},
		[]string{"name", "url"},
	)

	prometheus.MustRegister(HTTPResposeTime)
}

// Collect the metrics
func (c Collector) Collect(d *collector.CollectorData) error {

	HTTPResposeTime.WithLabelValues(d.GetName(), d.GetUrl()).Set(float64(d.ResponseTime.Seconds()))

	return nil
}
