package collector

import (
	"net/http"
	"time"

	"github.com/jesusvico/http-uptime/internal/endpoint"
)

type CollectorData struct {
	endpoint     endpoint.Endpoint
	Response     *http.Response
	ResponseTime time.Duration
}

type Collector interface {
	Register()
	Collect(d *CollectorData) error
}

func NewData(e endpoint.Endpoint) *CollectorData {
	startTime := time.Now()

	res, err := e.Request()
	resTime := time.Since(startTime)

	if err != nil {
		res = nil
	}

	c := CollectorData{
		endpoint:     e,
		Response:     res,
		ResponseTime: resTime,
	}

	return &c
}

func (c *CollectorData) HasData() bool {
	return c.Response != nil
}

func (c *CollectorData) GetName() string {
	return c.endpoint.Name
}

func (c *CollectorData) GetUrl() string {
	return c.endpoint.Url.String()
}
