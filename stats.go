package mqtt

import (
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type mqttMetrics struct {
	SentBytes        *metrics.Metric
	ReceivedBytes    *metrics.Metric
	SentMessages     *metrics.Metric
	ReceivedMessages *metrics.Metric
}

// registerMetrics registers the metrics for the mqtt module in the metrics registry
func registerMetrics(vu modules.VU) (mqttMetrics, error) {
	var err error
	registry := vu.InitEnv().Registry
	m := mqttMetrics{}

	m.SentBytes, err = registry.NewMetric("mqtt.sent.bytes", metrics.Counter)
	if err != nil {
		return m, err
	}

	m.ReceivedBytes, err = registry.NewMetric("mqtt.received.bytes", metrics.Counter)
	if err != nil {
		return m, err
	}

	m.SentMessages, err = registry.NewMetric("mqtt.sent.messages.count", metrics.Counter)
	if err != nil {
		return m, err
	}

	m.ReceivedMessages, err = registry.NewMetric("mqtt.received.messages.count", metrics.Counter)
	if err != nil {
		return m, err
	}
	return m, nil
}
