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

func (*Mqtt) Subscribe(
	ctx context.Context,
	reader paho.Client,
	// Topic to consume messages from
	topic string,
	// The QoS of messages
	qos,
	// timeout in sec
	timeout int,
) (chan paho.Message, error) {
	state := lib.GetState(ctx)

	if state == nil {
		ReportError(ErrorNilState, "Subscribe Cannot determine state")
		return nil, ErrorNilState
	}
	recieved := make(chan paho.Message, 1)
	messageCB := func(client paho.Client, msg paho.Message) {
		go func(msg paho.Message) {
			recieved <- msg
		}(msg)
	}
	if reader == nil {
		ReportError(ErrorNilReader, "Subscribe Reader is not ready")
		return nil, ErrorNilReader
	}
	token := reader.Subscribe(topic, byte(qos), messageCB)
	if !token.WaitTimeout(time.Duration(timeout) * time.Second) {
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

func (*Mqtt) Consume(
	ctx context.Context,
	reader paho.Client,
	recieved chan paho.Message,
	timeout int,
) (string, error) {
	state := lib.GetState(ctx)

	if state == nil {
		ReportError(ErrorNilState, "Cannot determine state")
		return "", ErrorNilState
	}

	// force close async
	defer func() { go reader.Disconnect(0) }()

	select {
	case msg := <-recieved:
		return string(msg.Payload()), nil
	case <-time.After(time.Second * time.Duration(timeout)):
		ReportError(ErrorMessageRecieveTimeout, "Message never recieved")
		return "", ErrorMessageRecieveTimeout
	}
}
