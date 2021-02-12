package mqtt

import (
	"context"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/loadimpact/k6/js/modules"
	"github.com/loadimpact/k6/lib"
)

func init() {
	modules.Register("k6/x/mqtt", new(Mqtt))
}

// Subscribe to the given topic return a channel to wait the message
func (*Mqtt) Subscribe(
	ctx context.Context,
	// Mqtt client to be used
	client paho.Client,
	// Topic to consume messages from
	topic string,
	// The QoS of messages
	qos,
	// timeout ms
	timeout uint,
) (chan paho.Message, error) {
	state := lib.GetState(ctx)

	if state == nil {
		ReportError(ErrorNilState, "Subscribe Cannot determine state")
		return nil, ErrorNilState
	}
	recieved := make(chan paho.Message)
	messageCB := func(client paho.Client, msg paho.Message) {
		go func(msg paho.Message) {
			recieved <- msg
		}(msg)
	}
	if client == nil {
		ReportError(ErrorNilReader, "Subscribe Reader is not ready")
		return nil, ErrorNilReader
	}
	token := client.Subscribe(topic, byte(qos), messageCB)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		ReportError(ErrorReaderTimeout, "Subscribe timeout")
		return nil, ErrorReaderTimeout
	}
	err := token.Error()
	if err != nil {
		ReportError(err, "Subscribe failed")
		return nil, err
	}
	return recieved, nil
}

// Consume will wait for one message to arrive
func (*Mqtt) Consume(
	ctx context.Context,
	recieved chan paho.Message,
	// timeout ms
	timeout uint,
) (string, error) {
	state := lib.GetState(ctx)
	if state == nil {
		ReportError(ErrorNilState, "Cannot determine state")
		return "", ErrorNilState
	}
	if recieved == nil {
		ReportError(ErrorNilReader, "Consume token is invalid")
		return "", ErrorNilReader
	}

	select {
	case msg := <-recieved:
		return string(msg.Payload()), nil
	case <-time.After(time.Millisecond * time.Duration(timeout)):
		ReportError(ErrorMessageRecieveTimeout, "Message never recieved")
		return "", ErrorMessageRecieveTimeout
	}
}
