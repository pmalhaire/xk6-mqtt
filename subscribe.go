package mqtt

import (
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/metrics"
)

// Subscribe to the given topic and returns a channel waiting for the message
func (m *Mqtt) Subscribe(
	// Mqtt client to be used
	client paho.Client,
	// Topic to consume messages from
	topic string,
	// The QoS of messages
	qos,
	// timeout ms
	timeout uint,
) chan paho.Message {
	recieved := make(chan paho.Message)
	messageCB := func(client paho.Client, msg paho.Message) {
		go func(msg paho.Message) {
			recieved <- msg
		}(msg)
	}
	if client == nil {
		common.Throw(nil, ErrorClient)
		return nil
	}
	token := client.Subscribe(topic, byte(qos), messageCB)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorTimeout)
		return nil
	}
	err := token.Error()
	if err != nil {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorTimeout)
		return nil
	}
	return recieved
}

// Consume will wait for one message to arrive
func (m *Mqtt) Consume(
	token chan paho.Message,
	// timeout ms
	timeout uint,
) string {
	state := m.vu.State()
	rt := m.vu.Runtime()

	if state == nil {
		common.Throw(rt, ErrorState)
		return ""
	}
	if token == nil {
		common.Throw(rt, ErrorConsumeToken)
		return ""
	}

	select {
	case msg := <-token:
		payload := msg.Payload()
		msgByteLen := len([]byte(payload))
		state := m.vu.State()
		state.BuiltinMetrics.DataReceived.Sink.Add(
			metrics.Sample{Metric: &metrics.Metric{}, Value: float64(msgByteLen), Time: time.Now()},
		)
		return string(payload)
	case <-time.After(time.Millisecond * time.Duration(timeout)):
		common.Throw(rt, ErrorTimeout)
		return ""
	}
}
