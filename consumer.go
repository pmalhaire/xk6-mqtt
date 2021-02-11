package mqtt

import (
	"context"
	"fmt"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/loadimpact/k6/js/modules"
	"github.com/loadimpact/k6/lib"
)

func init() {
	modules.Register("k6/x/mqtt", new(Mqtt))
}

type Mqtt struct {
}

func (*Mqtt) Reader(
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
	// timeout in sec
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
	opts.SetDefaultPublishHandler(func(client paho.Client, msg paho.Message) {
		fmt.Println("unexpected message recieved on topic", msg.Topic())
	})
	client := paho.NewClient(opts)
	token := client.Connect()
	if !token.WaitTimeout(time.Duration(timeout) * time.Second) {
		ReportError(token.Error(), "Connection timeout")
		return nil
	}
	err := token.Error()
	if err != nil {
		ReportError(err, "Connection failed")
		return nil
	}
	return client
}

func (*Mqtt) Consume(
	ctx context.Context, reader paho.Client,
	// Topic to consume messages from
	topic string,
	// The QoS of messages
	qos,
	// timeout in sec
	timeout int,
) (string, error) {
	state := lib.GetState(ctx)

	if state == nil {
		ReportError(ErrorNilState, "Cannot determine state")
		return "", ErrorNilState
	}
	recieved := make(chan paho.Message)
	messageCB := func(client paho.Client, msg paho.Message) {
		go func(msg paho.Message) { recieved <- msg }(msg)
	}
	if reader == nil {
		ReportError(ErrorNilReader, "Reader is not ready")
		return "", ErrorNilReader
	}
	token := reader.Subscribe(topic, byte(qos), messageCB)
	if !token.WaitTimeout(time.Duration(timeout) * time.Second) {
		ReportError(token.Error(), "Consume timeout")
		return "", ErrorReaderTimeout
	}
	err := token.Error()
	if err != nil {
		ReportError(err, "Subscribe failed")
		return "", err
	}

	// force disconnect after read
	defer func() { go reader.Disconnect(0) }()
	select {
	case msg := <-recieved:
		return string(msg.Payload()), nil
	case <-time.After(time.Second * time.Duration(timeout)):
		return "", ErrorReaderTimeout
	}
}
