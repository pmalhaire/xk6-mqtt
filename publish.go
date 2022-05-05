package mqtt

import (
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/metrics"
)

// Publish allow to publish one message
func (m *Mqtt) Publish(
	client paho.Client,
	topic string,
	qos int,
	message string,
	retain bool,
	timeout int,
) {
	if client == nil {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorClient)
		return
	}

	token := client.Publish(topic, byte(qos), retain, message)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorTimeout)
		return
	}
	if err := token.Error(); err != nil {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorPublish)
		return
	}
	state := m.vu.State()
	msgByteLen := len([]byte(message))
	state.BuiltinMetrics.DataSent.Sink.Add(
		metrics.Sample{Metric: &metrics.Metric{}, Value: float64(msgByteLen), Time: time.Now()},
	)
	return
}
