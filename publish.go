package mqtt

import (
	"context"
	"time"
	"encoding/hex"

	paho "github.com/eclipse/paho.mqtt.golang"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/lib"
)


func asHex(
	stringVal string,
	convertToBin bool,
) interface{} {
	if convertToBin == true {
		data, err := hex.DecodeString(stringVal)
			if err != nil {
				panic(err)
			}
			return data
	}
	return stringVal
}

// Publish allow to publish one message
func (*Mqtt) Publish(
	ctx context.Context,
	client paho.Client,
	topic string,
	qos int,
	message string,
	retain bool,
	timeout int,
	hex2bin bool,
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
	outboundMessage := asHex(message, hex2bin)
	token := client.Publish(topic, byte(qos), retain, outboundMessage)
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
