package mqtt

import (
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

// RootModule is the root module for the mqtt API
type RootModule struct{}

func init() {
	modules.Register("k6/x/mqtt", new(RootModule))
}

// MqttAPI is the k6 extension implementing the Mqtt api
type MqttAPI struct { //nolint:revive
	vu              modules.VU
	initEnvironment *common.InitEnvironment
}

// NewModuleInstance implements the modules.Module interface and returns
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &MqttAPI{vu: vu, initEnvironment: vu.InitEnv()}
}

// Exports exposes the given object in ts
func (m *MqttAPI) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Client": m.client,
		},
	}
}

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &MqttAPI{}
	_ modules.Module   = &RootModule{}
)
