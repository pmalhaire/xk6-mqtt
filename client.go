package mqtt

import (
	"context"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/lib"
)

// Mqtt is the objet to be used in tests
type Mqtt struct {
}

// Connect create a connection to mqtt
func (*Mqtt) Connect(
	ctx context.Context,
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
	// timeout ms
	timeout uint,

) paho.Client {
	state := lib.GetState(ctx)
	if state == nil {
		common.Throw(common.GetRuntime(ctx), ErrorState)
		return nil
	}
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

	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		common.Throw(common.GetRuntime(ctx), ErrorTimeout)
		return nil
	}
	if token.Error() != nil {
		common.Throw(common.GetRuntime(ctx), ErrorClient)
		return nil
	}
	return client
}

// Close the given client
func (*Mqtt) Close(
	ctx context.Context,
	// Mqtt client to be closed
	client paho.Client,
	// timeout ms
	timeout uint,
) {
	state := lib.GetState(ctx)
	if state == nil {
		common.Throw(common.GetRuntime(ctx), ErrorState)
		return
	}
	client.Disconnect(timeout)
	return
}
