package metrics_prometheus

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsPrometheus struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

func NewMetricsPrometheus(name string) (*MetricsPrometheus, error) {
	metr := MetricsPrometheus{}
	metr.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_hits_total",
	})

	if err := prometheus.Register(metr.HitsTotal); err != nil {
		return nil, err
	}

	metr.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name + "_hits",
		},
		[]string{"status", "method"},
	)

	if err := prometheus.Register(metr.Hits); err != nil {
		return nil, err
	}

	metr.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name + "_times",
		},
		[]string{"status", "method"},
	)

	if err := prometheus.Register(metr.Times); err != nil {
		return nil, err
	}

	if err := prometheus.Register(collectors.NewBuildInfoCollector()); err != nil {
		return nil, err
	}

	return &metr, nil
}

func (m *MetricsPrometheus) GetHandler() http.Handler {
	return promhttp.Handler()
}

func (m *MetricsPrometheus) IncHits(status metrics.Status, method string) {
	m.HitsTotal.Inc()
	m.Hits.WithLabelValues(status.String(), method).Inc()
}

func (m *MetricsPrometheus) ObserveResponseTime(status metrics.Status, method string, observeTime float64) {
	m.Times.WithLabelValues(status.String(), method).Observe(observeTime)
}

var _ metrics.Metrics = &MetricsPrometheus{}
