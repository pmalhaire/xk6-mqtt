package mqtt

import (
	"context"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/loadimpact/k6/lib"
)

// Publish allow to publish one message
func (*Mqtt) Publish(
	ctx context.Context,
	client paho.Client,
	topic string,
	qos int,
	message string,
	retain bool,
	timeout int,
) error {
	if client == nil {
		ReportError(ErrorNilWriter, "Writer is not ready")
		return ErrorNilWriter
	}
	state := lib.GetState(ctx)
	if state == nil {
		ReportError(ErrorNilState, "Cannot determine state")
		return ErrorNilState
	}
	token := client.Publish(topic, byte(qos), retain, message)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		ReportError(token.Error(), "Produce timeout")
		return ErrorWriterTimeout
	}
	if err := token.Error(); err != nil {
		ReportError(err, "Produce failed")
		return err
	}

	return nil
}
