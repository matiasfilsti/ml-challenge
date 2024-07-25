package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "NumberOfRequest",
		Help: "number of total request",
	})
)

func RecordMetrics() {
	go func() {
		opsProcessed.Inc()
	}()

}
