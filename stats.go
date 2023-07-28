package mqtt

import (
	"errors"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type mqttMetrics struct {
	SentBytes        *metrics.Metric
	ReceivedBytes    *metrics.Metric
	SentMessages     *metrics.Metric
	ReceivedMessages *metrics.Metric
	TagsAndMeta      *metrics.TagsAndMeta
}

// registerMetrics registers the metrics for the mqtt module in the metrics registry
func registerMetrics(vu modules.VU) (mqttMetrics, error) {
	var err error
	m := mqttMetrics{}
	env := vu.InitEnv()
	if env == nil {
		return m, errors.New("missing env")
	}
	registry := env.Registry
	if registry == nil {
		return m, errors.New("missing registry")
	}
	m.SentBytes, err = registry.NewMetric("mqtt_sent_bytes", metrics.Counter)
	if err != nil {
		return m, err
	}

	m.ReceivedBytes, err = registry.NewMetric("mqtt_received_bytes", metrics.Counter)
	if err != nil {
		return m, err
	}

	m.SentMessages, err = registry.NewMetric("mqtt_sent_messages_count", metrics.Counter)
	if err != nil {
		return m, err
	}

	m.ReceivedMessages, err = registry.NewMetric("mqtt_received_messages_count", metrics.Counter)
	if err != nil {
		return m, err
	}
	m.TagsAndMeta = &metrics.TagsAndMeta{
		Tags: registry.RootTagSet(),
	}
	return m, nil
}
