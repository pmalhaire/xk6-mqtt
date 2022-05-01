package mqtt

import (
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/mqtt", New())
}

type (
	// MqttModule is the global module instance that will create Mqtt
	// instances for each VU.
	MqttModule struct{}

	// Mqtt represents an instance of the JS module.
	MqttInstance struct {
		// modules.VU provides some useful methods for accessing internal k6
		// objects like the global context, VU state and goja runtime.
		vu modules.VU
		// Mqtt is the exported module instance.
		*Mqtt
	}
)

// New returns a pointer to a new RootModule instance.
func New() *MqttModule {
	return &MqttModule{}
}

// Mqtt is the objet to be used in tests
type Mqtt struct {
	vu modules.VU
}

// NewModuleInstance implements the modules.Module interface and returns

// a new instance for each VU.
func (*MqttModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &MqttInstance{vu: vu, Mqtt: &Mqtt{vu: vu}}
}

func (m *MqttInstance) Exports() modules.Exports {
	return modules.Exports{Default: m}

}

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &MqttInstance{}
	_ modules.Module   = &MqttModule{}
)

// Subscribe to the given topic and returns a channel waiting for the message
func (m *Mqtt) Subscribe(
	// Mqtt client to be used
	client paho.Client,
	// Topic to consume messages from
	topic string,
	// The QoS of messages
	qos,
	// timeout ms
	timeout uint,
) chan paho.Message {
	recieved := make(chan paho.Message)
	messageCB := func(client paho.Client, msg paho.Message) {
		go func(msg paho.Message) {
			recieved <- msg
		}(msg)
	}
	if client == nil {
		common.Throw(nil, ErrorClient)
		return nil
	}
	token := client.Subscribe(topic, byte(qos), messageCB)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorTimeout)
		return nil
	}
	err := token.Error()
	if err != nil {
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorTimeout)
		return nil
	}
	return recieved
}

// Consume will wait for one message to arrive
func (m *Mqtt) Consume(
	token chan paho.Message,
	// timeout ms
	timeout uint,
) string {
	state := m.vu.State()
	rt := m.vu.Runtime()

	if state == nil {
		common.Throw(rt, ErrorState)
		return ""
	}
	if token == nil {
		common.Throw(rt, ErrorConsumeToken)
		return ""
	}

	select {
	case msg := <-token:
		payload := msg.Payload()
		return string(payload)
	case <-time.After(time.Millisecond * time.Duration(timeout)):
		common.Throw(rt, ErrorTimeout)
		return ""
	}
}
