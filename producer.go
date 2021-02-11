package mqtt

import (
	"context"
	"errors"
	"fmt"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/loadimpact/k6/lib"
)

func (*Mqtt) Writer(brokers []string, topic string) paho.Client {
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
	return client
}

func (*Mqtt) Produce(
	ctx context.Context, writer paho.Client, topic string, qos int, messages []string) error {
	state := lib.GetState(ctx)
	err := errors.New("State is nil")

	if state == nil {
		ReportError(err, "Cannot determine state")
		return err
	}

	for _, message := range messages {
		token := writer.Publish(topic, byte(qos), false, message)
		token.Wait()
		panic(token.Error())
	}

	return nil
}
