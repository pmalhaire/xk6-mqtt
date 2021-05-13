package mqtt

import (
	"context"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/lib"
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
) {
	state := lib.GetState(ctx)
	if state == nil {
		common.Throw(common.GetRuntime(ctx), ErrorState)
		return
	}
	if client == nil {
		common.Throw(common.GetRuntime(ctx), ErrorClient)
		return
	}

	token := client.Publish(topic, byte(qos), retain, message)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		common.Throw(common.GetRuntime(ctx), ErrorTimeout)
		return
	}
	if err := token.Error(); err != nil {
		common.Throw(common.GetRuntime(ctx), ErrorPublish)
		return
	}

	return
}
