// Package mqtt xk6 extenstion to suppor mqtt with k6
package mqtt

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/dop251/goja"
	paho "github.com/eclipse/paho.golang/paho"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/mstoykov/k6-taskqueue-lib/taskqueue"
	// "go.uber.org/zap"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

// To preserve the connection
type client struct {
	vu         modules.VU
	metrics    *mqttMetrics
	conf       conf
	obj        *goja.Object // the object that is given to js to interact with the WebSocket
	tq              *taskqueue.TaskQueue
	connectionManager *autopaho.ConnectionManager
	clientConfig 	autopaho.ClientConfig
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
	// path to caRoot path
	caRootPath string
	// path to client cert file
	clientCertPath string
	// path to client cert key file
	clientCertKeyPath string
}

//nolint:nosnakecase // their choice not mine
func (m *MqttAPI) client(c goja.ConstructorCall) *goja.Object {
	fmt.Println("In client.go, Creating client")
	serversArray := c.Argument(0)
	rt := m.vu.Runtime()
	if serversArray == nil || goja.IsUndefined(serversArray) {
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
	if userValue == nil || goja.IsUndefined(userValue) {
		common.Throw(rt, errors.New("Client requires a user value"))
	}
	clientConf.user = userValue.String()
	passwordValue := c.Argument(2)
	if userValue == nil || goja.IsUndefined(passwordValue) {
		common.Throw(rt, errors.New("Client requires a password value"))
	}
	clientConf.password = passwordValue.String()
	cleansessValue := c.Argument(3)
	if cleansessValue == nil || goja.IsUndefined(cleansessValue) {
		common.Throw(rt, errors.New("Client requires a cleaness value"))
	}
	clientConf.cleansess = cleansessValue.ToBoolean()

	clientIDValue := c.Argument(4)
	if clientIDValue == nil || goja.IsUndefined(clientIDValue) {
		common.Throw(rt, errors.New("Client requires a clientID value"))
	}
	clientConf.clientid = clientIDValue.String()

	timeoutValue := c.Argument(5)
	if timeoutValue == nil || goja.IsUndefined(timeoutValue) {
		common.Throw(rt, errors.New("Client requires a timeout value"))
	}
	clientConf.timeout = uint(timeoutValue.ToInteger())

	// optional args
	if caRootPathValue := c.Argument(6); caRootPathValue == nil || goja.IsUndefined(caRootPathValue) {
		clientConf.caRootPath = ""
	} else {
		clientConf.caRootPath = caRootPathValue.String()
	}
	if clientCertPathValue := c.Argument(7); clientCertPathValue == nil || goja.IsUndefined(clientCertPathValue) {
		clientConf.clientCertPath = ""
	} else {
		clientConf.clientCertPath = clientCertPathValue.String()
	}
	if clientCertKeyPathValue := c.Argument(8); clientCertKeyPathValue == nil || goja.IsUndefined(clientCertKeyPathValue) {
		clientConf.clientCertKeyPath = ""
	} else {
		clientConf.clientCertKeyPath = clientCertKeyPathValue.String()
	}

	client := &client{
		vu:      m.vu,
		metrics: &m.metrics,
		conf:    clientConf,
		obj:     rt.NewObject(),
	}
	must := func(err error) {
		if err != nil {
			common.Throw(rt, err)
		}
	}

	must(client.obj.DefineDataProperty(
		"connect", rt.ToValue(client.Connect), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))
	must(client.obj.DefineDataProperty(
		"publish", rt.ToValue(client.Publish), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))

	must(client.obj.DefineDataProperty(
		"close", rt.ToValue(client.Close), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))

	return client.obj
}

// type ConnectionPreserver struct {
// 	connectionManager *autopaho.ConnectionManager
// }

// func (c *ConnectionPreserver) init_connection(ctx context.Context, cliCfg *autopaho.ClientConfig) error {
//     connection, err := autopaho.NewConnection(ctx, *cliCfg)
//     if err != nil {
//         return err
//     }
//     if err = connection.AwaitConnection(ctx); err != nil {
//         return fmt.Errorf("failed to connect to broker: %w", err)
//     }
//     c.connectionManager = connection
//     return nil
// }

// Connect create a connection to mqtt
func (c *client) Connect() error {
	fmt.Println("connecting client, ")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	parsed_urls := []*url.URL{}
	for _, server := range c.conf.servers {
		parsed_url, err := url.Parse(server)
		if err != nil {
			panic(err)
		}
		parsed_urls = append(parsed_urls, parsed_url)
	}

	topic := "hello, test_topic"


	cliCfg := autopaho.ClientConfig{
		ServerUrls: parsed_urls,
		KeepAlive:  20, // Keepalive message should be sent every 20 seconds
		// CleanStartOnInitialConnection defaults to false. Setting this to true will clear the session on the first connection.
		CleanStartOnInitialConnection: false,
		// SessionExpiryInterval - Seconds that a session will survive after disconnection.
		// It is important to set this because otherwise, any queued messages will be lost if the connection drops and
		// the server will not queue messages while it is down. The specific setting will depend upon your needs
		// (60 = 1 minute, 3600 = 1 hour, 86400 = one day, 0xFFFFFFFE = 136 years, 0xFFFFFFFF = don't expire)
		SessionExpiryInterval: 60,
		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
			fmt.Println("mqtt connection up")
			// Subscribing in the OnConnectionUp callback is recommended (ensures the subscription is reestablished if
			// the connection drops)
			if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
				Subscriptions: []paho.SubscribeOptions{
					{Topic: topic, QoS: 1},
				},
			}); err != nil {
				fmt.Printf("failed to subscribe (%s). This is likely to mean no messages will be received.", err)
			}
			// fmt.Println("mqtt subscription made")
		},
		OnConnectError: func(err error) { fmt.Printf("error whilst attempting connection: %s\n", err) },
		// eclipse/paho.golang/paho provides base mqtt functionality, the below config will be passed in for each connection
		ClientConfig: paho.ClientConfig{
			// If you are using QOS 1/2, then it's important to specify a client id (which must be unique)
			ClientID: c.conf.clientid,
			// OnPublishReceived is a slice of functions that will be called when a message is received.
			// You can write the function(s) yourself or use the supplied Router
			OnPublishReceived: []func(paho.PublishReceived) (bool, error){
				func(pr paho.PublishReceived) (bool, error) {
					fmt.Printf("received message on topic %s; body: %s (retain: %t)\n", pr.Packet.Topic, pr.Packet.Payload, pr.Packet.Retain)
					return true, nil
				}},
			OnClientError: func(err error) { fmt.Printf("client error: %s\n", err) },
			OnServerDisconnect: func(d *paho.Disconnect) {
				fmt.Println("WE ARE DISCONNECTED")
				if d.Properties != nil {
					fmt.Printf("server requested disconnect: %s\n", d.Properties.ReasonString)
				} else {
					fmt.Printf("server requested disconnect; reason code: %d\n", d.ReasonCode)
				}
			},
		},
	}
	cliCfg.ConnectUsername = c.conf.user
	cliCfg.ConnectPassword = []byte(c.conf.password)

	// connection, err := autopaho.NewConnection(ctx, cliCfg) // starts process; will reconnect until context cancelled
	// if err != nil {
	// 	panic(err)
	// }
	// // Wait for the connection to come up
	// if err = connection.AwaitConnection(ctx); err != nil {
	// 	fmt.Println("failed to connect to broker!!!!!!")
	// 	panic(err)
	// }

	// // preserved_connection := ConnectionPreserver{}
	// // preserved_connection.init_connection(ctx, &cliCfg)

	// c.connectionManager = connection

	var err error
	c.connectionManager, err = autopaho.NewConnection(ctx, cliCfg)
	if err != nil {
		panic(err)
	}
	if err = c.connectionManager.AwaitConnection(ctx); err != nil {
		fmt.Println("failed to connect to broker!!!!!!")
		panic(err)
	}

	c.clientConfig = cliCfg

	// Publish a test message (use PublishViaQueue if you don't want to wait for a response)
	// fmt.Println("Publishing message, ", "hello")
	// _, publish_error := c.connectionManager.Publish(ctx, &paho.Publish{
	// 	QoS:     1,
	// 	Topic:  "vehicle_state_ota/8b9dbede-27fc-485a-a55b-e20a72bcb257",
	// 	Payload: []byte("hello not thissssss"),
	// })
	// if (publish_error != nil) {
	// 	return publish_error
	// }
	return nil
}

func (c *client) Publish(
	topic string,
	qos int,
	message string,
	retain bool,
	timeout uint,
) error {
	fmt.Println("inside publish, ", topic, qos, message, retain, timeout)
	// return nil
	// sync case no callback added
	// return nil

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, publish_error := c.connectionManager.Publish(ctx, &paho.Publish{
		QoS:     1,
		Topic:  topic,
		Payload: []byte(message),
	})
	if (publish_error != nil) {
		fmt.Println("error publishing message: ", publish_error)
		return publish_error
	}
	return nil
}

// Close the given client
// wait for pending connections for timeout (ms) before closing
func (c *client) Close() {
	fmt.Println("INSIDE CLOSE")
	// exit subscribe task queue if running
	c.connectionManager.Done()
	if c.tq != nil {
		c.tq.Close()
	}
}

// error event for async
//
//nolint:nosnakecase // their choice not mine
func (c *client) newErrorEvent(msg string) *goja.Object {
	rt := c.vu.Runtime()
	o := rt.NewObject()
	must := func(err error) {
		if err != nil {
			common.Throw(rt, err)
		}
	}

	must(o.DefineDataProperty("type", rt.ToValue("error"), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))
	must(o.DefineDataProperty("message", rt.ToValue(msg), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))
	return o
}
