package mqtt

import (
	"context"
	"fmt"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/loadimpact/k6/js/modules"
	"github.com/loadimpact/k6/lib"
)

func init() {
	modules.Register("k6/x/mqtt", new(Mqtt))
}

type Mqtt struct {
}

func (*Mqtt) Reader() paho.Client {
	// The full URL of the MQTT server to connect to"
	server := "127.0.0.1:1883"
	// A username to authenticate to the MQTT server
	user := "username"
	// Password to match username
	password := "password"

	// clean session setting
	cleansess := false

	// Client id for reader
	clientid := "1"

	// Topic to publish and receive the messages on
	topic := "topic"
	// The QoS to send the messages at
	qos := 1

	opts := paho.NewClientOptions()
	opts.AddBroker(server)
	opts.SetClientID(clientid)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetCleanSession(cleansess)
	opts.SetDefaultPublishHandler(func(client paho.Client, msg paho.Message) {
		fmt.Println("msg")
	})
	client := paho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := client.Subscribe(topic, byte(qos), nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func (*Mqtt) Consume(
	ctx context.Context, reader *paho.Client) {
	state := lib.GetState(ctx)

	if state == nil {
		ReportError(nil, "Cannot determine state")
		return
	}
}
