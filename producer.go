package mqtt

import (
	"context"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/loadimpact/k6/lib"
)

func (*Mqtt) Writer(
	// The list of URL of  MQTT server to connect to
	servers []string,
	// A username to authenticate to the MQTT server
	user,
	// Password to match username
	password string,

	// clean session setting
	cleansess bool,

	// Client id for reader
	clientid string,

	// timeout sec
	timeout int,

) paho.Client {

	opts := paho.NewClientOptions()
	for i := range servers {
		opts.AddBroker(servers[i])
	}
	opts.SetClientID(clientid)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetCleanSession(cleansess)
	client := paho.NewClient(opts)
	token := client.Connect()
	if !token.WaitTimeout(time.Duration(timeout) * time.Second) {
		ReportError(token.Error(), "Connect timeout")
		return nil
	}
	if token.Error() != nil {
		ReportError(token.Error(), "Connect failed")
		return nil
	}
	return client
}

func (*Mqtt) Produce(
	ctx context.Context,
	writer paho.Client,
	topic string,
	qos int,
	message string,
	retain bool,
	timeout int,
) error {
	if writer == nil {
		ReportError(ErrorNilWriter, "Writer is not ready")
		return ErrorNilWriter
	}
	state := lib.GetState(ctx)
	if state == nil {
		ReportError(ErrorNilState, "Cannot determine state")
		return ErrorNilState
	}
	// force close async
	defer func() { go writer.Disconnect(0) }()

	token := writer.Publish(topic, byte(qos), retain, message)
	if !token.WaitTimeout(time.Duration(timeout) * time.Second) {
		ReportError(token.Error(), "Produce timeout")
		return ErrorWriterTimeout
	}
	if err := token.Error(); err != nil {
		ReportError(err, "Produce failed")
		return err
	}

	return nil
}
