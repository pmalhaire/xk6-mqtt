package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"go.k6.io/k6/js/common"
)

// Connect create a connection to mqtt
func (m *Mqtt) Connect(
	//ctx context.Context,
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
	// path to local cert
	certPath string,

) paho.Client {
	// TODO fix connect is done when no VU is available
	state := m.vu.State()
	rt := m.vu.Runtime()
	if state == nil {
		common.Throw(rt, ErrorState)
		return nil
	}
	opts := paho.NewClientOptions()
	// Use local cert if specified
	if len(certPath) > 0 {
		mqtt_tls_ca, err := os.ReadFile(certPath)
		if err != nil {
			panic(err)
		}

		root_ca := x509.NewCertPool()
		load_ca := root_ca.AppendCertsFromPEM([]byte(mqtt_tls_ca))
		if !load_ca {
			panic("failed to parse root certificate")
		}
		tlsConfig := &tls.Config{RootCAs: root_ca}
		opts.SetTLSConfig(tlsConfig)
	}
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
		rt := m.vu.Runtime()
		common.Throw(rt, ErrorTimeout)
		return nil
	}
	if token.Error() != nil {
		rt := m.vu.Runtime()
		common.Throw(rt, token.Error())
		return nil
	}
	return client
}

// Close the given client
func (m *Mqtt) Close(
	//ctx context.Context,
	// Mqtt client to be closed
	client paho.Client,
	// timeout ms
	timeout uint,
) {
	state := m.vu.State()
	rt := m.vu.Runtime()
	if state == nil {
		common.Throw(rt, ErrorState)
		return
	}
	client.Disconnect(timeout)
	return
}
