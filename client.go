// Package mqtt xk6 extenstion to suppor mqtt with k6
package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/grafana/sobek"
	"github.com/mstoykov/k6-taskqueue-lib/taskqueue"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

type client struct {
	vu         modules.VU
	metrics    *mqttMetrics
	conf       conf
	pahoClient paho.Client
	obj        *sobek.Object // the object that is given to js to interact with the WebSocket

	// listeners
	// this return sobek.value *and* error in order to return error on exception instead of panic
	// https://pkg.go.dev/github.com/dop251/goja#hdr-Functions
	messageListener func(sobek.Value) (sobek.Value, error)
	errorListener   func(sobek.Value) (sobek.Value, error)
	tq              *taskqueue.TaskQueue
	messageChan     chan paho.Message
	subRefCount     int
}

type conf struct {
	// The list of URL of  MQTT server to connect to
	servers []string
	// A username to authenticate to the MQTT server
	user string
	// Password to match username
	password string
	// clean session setting
	cleansess bool
	// Client id for reader
	clientid string
	// timeout ms
	timeout uint
	// caRoot PEM string or path
	caRoot string
	// client cert PEM string or path
	clientCert string
	// client cert PEM string or path
	clientCertKey string
	// wether to skip the cert validity check
	skipTLSValidation bool
}

const (
	sentBytesLabel             = "mqtt_sent_bytes"
	receivedBytesLabel         = "mqtt_received_bytes"
	sentMessagesCountLabel     = "mqtt_sent_messages_count"
	receivedMessagesCountLabel = "mqtt_received_messages_count"
)

func getLabels(labelsArg sobek.Value, rt *sobek.Runtime) mqttMetricsLabels {
	labels := mqttMetricsLabels{}
	metricsLabels := labelsArg
	if metricsLabels == nil || sobek.IsUndefined(metricsLabels) {
		// set default values
		labels.SentBytesLabel = sentBytesLabel
		labels.ReceivedBytesLabel = receivedBytesLabel
		labels.SentMessagesCountLabel = sentMessagesCountLabel
		labels.ReceivedMessagesCountLabel = receivedMessagesCountLabel
		return labels
	}

	labelsJS, ok := metricsLabels.Export().(map[string]any)
	if !ok {
		common.Throw(rt, fmt.Errorf("invalid metricsLabels %#v", metricsLabels.Export()))
	}
	labels.SentBytesLabel, ok = labelsJS["sentBytesLabel"].(string)
	if !ok {
		common.Throw(rt, fmt.Errorf("invalid metricsLabels sentBytesLabel %#v", metricsLabels.Export()))
	}
	labels.ReceivedBytesLabel, ok = labelsJS["receivedBytesLabel"].(string)
	if !ok {
		common.Throw(rt, fmt.Errorf("invalid metricsLabels receivedBytesLabel %#v", metricsLabels.Export()))
	}
	labels.SentMessagesCountLabel, ok = labelsJS["sentMessagesCountLabel"].(string)
	if !ok {
		common.Throw(rt, fmt.Errorf("invalid metricsLabels sentMessagesCountLabel %#v", metricsLabels.Export()))
	}
	labels.ReceivedMessagesCountLabel, ok = labelsJS["receivedMessagesCountLabel"].(string)
	if !ok {
		common.Throw(rt, fmt.Errorf("invalid metricsLabels receivedMessagesCountLabel %#v", metricsLabels.Export()))
	}

	return labels
}

//nolint:nosnakecase // their choice not mine
func (m *MqttAPI) client(c sobek.ConstructorCall) *sobek.Object {
	serversArray := c.Argument(0)
	rt := m.vu.Runtime()
	if serversArray == nil || sobek.IsUndefined(serversArray) {
		common.Throw(rt, errors.New("Client requires a server list"))
	}
	var servers []string
	var clientConf conf
	err := rt.ExportTo(serversArray, &servers)
	if err != nil {
		common.Throw(rt,
			fmt.Errorf("Client requires valid server list, but got %q which resulted in %w", serversArray, err))
	}
	clientConf.servers = servers
	userValue := c.Argument(1)
	if userValue == nil || sobek.IsUndefined(userValue) {
		common.Throw(rt, errors.New("Client requires a user value"))
	}
	clientConf.user = userValue.String()
	passwordValue := c.Argument(2)
	if passwordValue == nil || sobek.IsUndefined(passwordValue) {
		common.Throw(rt, errors.New("Client requires a password value"))
	}
	clientConf.password = passwordValue.String()
	cleansessValue := c.Argument(3)
	if cleansessValue == nil || sobek.IsUndefined(cleansessValue) {
		common.Throw(rt, errors.New("Client requires a cleansess value"))
	}
	clientConf.cleansess = cleansessValue.ToBoolean()

	clientIDValue := c.Argument(4)
	if clientIDValue == nil || sobek.IsUndefined(clientIDValue) {
		common.Throw(rt, errors.New("Client requires a clientID value"))
	}
	clientConf.clientid = clientIDValue.String()

	timeoutValue := c.Argument(5)
	if timeoutValue == nil || sobek.IsUndefined(timeoutValue) {
		common.Throw(rt, errors.New("Client requires a timeout value"))
	}
	clientConf.timeout = uint(timeoutValue.ToInteger())

	// optional args
	if caRootValue := c.Argument(6); caRootValue == nil || sobek.IsUndefined(caRootValue) {
		clientConf.caRoot = ""
	} else {
		clientConf.caRoot = caRootValue.String()
	}
	if clientCertValue := c.Argument(7); clientCertValue == nil || sobek.IsUndefined(clientCertValue) {
		clientConf.clientCert = ""
	} else {
		clientConf.clientCert = clientCertValue.String()
	}
	if clientCertKeyValue := c.Argument(8); clientCertKeyValue == nil ||
		sobek.IsUndefined(clientCertKeyValue) {
		clientConf.clientCertKey = ""
	} else {
		clientConf.clientCertKey = clientCertKeyValue.String()
	}
	labels := getLabels(c.Argument(9), rt)
	metrics, err := registerMetrics(m.vu, labels)
	if err != nil {
		common.Throw(m.vu.Runtime(), err)
	}

	skipTLS := c.Argument(10)
	clientConf.skipTLSValidation = skipTLS.ToBoolean()

	client := &client{
		vu:      m.vu,
		metrics: &metrics,
		conf:    clientConf,
		obj:     rt.NewObject(),
	}
	must := func(err error) {
		if err != nil {
			common.Throw(rt, err)
		}
	}

	// TODO add onmessage,onclose and so on
	must(client.obj.DefineDataProperty(
		"addEventListener", rt.ToValue(client.AddEventListener), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(client.obj.DefineDataProperty(
		"subContinue", rt.ToValue(client.SubContinue), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(client.obj.DefineDataProperty(
		"connect", rt.ToValue(client.Connect), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(client.obj.DefineDataProperty(
		"isConnected", rt.ToValue(client.IsConnected), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(client.obj.DefineDataProperty(
		"publish", rt.ToValue(client.Publish), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(client.obj.DefineDataProperty(
		"subscribe", rt.ToValue(client.Subscribe), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))

	must(client.obj.DefineDataProperty(
		"close", rt.ToValue(client.Close), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))

	return client.obj
}

// Connect create a connection to mqtt
func (c *client) Connect() error {
	opts := paho.NewClientOptions()

	var tlsConfig *tls.Config
	// Use root CA if specified
	if len(c.conf.caRoot) > 0 {
		var caContent []byte
		if strings.HasPrefix(c.conf.caRoot, "-----BEGIN CERTIFICATE-----") {
			// PEM content
			caContent = []byte(c.conf.caRoot)
		} else {
			// probably path to pem file
			var err error
			caContent, err = os.ReadFile(c.conf.caRoot)
			if err != nil {
				panic(err)
			}
		}
		rootCA := x509.NewCertPool()
		loadCA := rootCA.AppendCertsFromPEM(caContent)
		if !loadCA {
			panic("failed to parse root certificate")
		}
		tlsConfig = &tls.Config{
			RootCAs:    rootCA,
			MinVersion: tls.VersionTLS13,
		}
	}
	// Use local cert if specified
	if len(c.conf.clientCert) > 0 {
		var cert tls.Certificate
		var err error
		if strings.HasPrefix(c.conf.clientCert, "-----BEGIN CERTIFICATE-----") && strings.HasPrefix(c.conf.clientCertKey, "-----BEGIN RSA PRIVATE KEY-----") {
			// PEM content
			cert, err = tls.X509KeyPair([]byte(c.conf.clientCert), []byte(c.conf.clientCertKey))
		} else {
			// probably paths to pem files
			cert, err = tls.LoadX509KeyPair(c.conf.clientCert, c.conf.clientCertKey)
		}
		if err != nil {
			panic("failed to parse client certificate")
		}
		if tlsConfig != nil {
			tlsConfig.Certificates = []tls.Certificate{cert}
		} else {
			tlsConfig = &tls.Config{
				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS13,
			}
		}
	}

	// set tls if skip is forced
	if c.conf.skipTLSValidation {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: c.conf.skipTLSValidation, //nolint:gosec
		}
	}

	if tlsConfig != nil {
		opts.SetTLSConfig(tlsConfig)
	}
	for i := range c.conf.servers {
		opts.AddBroker(c.conf.servers[i])
	}
	opts.SetClientID(c.conf.clientid)
	opts.SetUsername(c.conf.user)
	opts.SetPassword(c.conf.password)
	opts.SetCleanSession(c.conf.cleansess)
	client := paho.NewClient(opts)
	token := client.Connect()
	rt := c.vu.Runtime()
	if !token.WaitTimeout(time.Duration(c.conf.timeout) * time.Millisecond) {
		common.Throw(rt, ErrTimeout)
		return ErrTimeout
	}
	if token.Error() != nil {
		common.Throw(rt, token.Error())
		return token.Error()
	}
	c.pahoClient = client
	return nil
}

// Close the given client
// wait for pending connections for timeout (ms) before closing
func (c *client) Close() {
	// exit subscribe task queue if running
	if c.tq != nil {
		c.tq.Close()
	}
	// disconnect client
	if c.pahoClient != nil && c.pahoClient.IsConnected() {
		c.pahoClient.Disconnect(c.conf.timeout)
	}
}

// IsConnected the given client
func (c *client) IsConnected() bool {
	if c.pahoClient == nil || !c.pahoClient.IsConnected() {
		return false
	}
	return true
}

// error event for async
//
//nolint:nosnakecase // their choice not mine
func (c *client) newErrorEvent(msg string) *sobek.Object {
	rt := c.vu.Runtime()
	o := rt.NewObject()
	must := func(err error) {
		if err != nil {
			common.Throw(rt, err)
		}
	}

	must(o.DefineDataProperty("type", rt.ToValue("error"), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(o.DefineDataProperty("message", rt.ToValue(msg), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	return o
}
