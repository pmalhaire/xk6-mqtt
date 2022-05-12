package mqtt

import (
	"time"

	"go.k6.io/k6/js/common"
)

// Publish allow to publish one message
func (m *client) Publish(
	topic string,
	qos int,
	message string,
	retain bool,
	timeout uint,
) error {
	if m.pahoClient == nil || !m.pahoClient.IsConnected() {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrClient)
		return ErrClient
	}

	token := m.pahoClient.Publish(topic, byte(qos), retain, message)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrTimeout)
		return ErrTimeout
	}
	if err := token.Error(); err != nil {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrPublish)
		return ErrPublish
	}
	return nil
}
