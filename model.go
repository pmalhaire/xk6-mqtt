package mqtt

import (
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

// TODO improve this StartTime is not a good way
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
