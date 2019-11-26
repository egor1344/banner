package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ApiCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_count",
		Help: "Client user api",
	})
)
