package mqtt

import (
	"errors"
	"time"

	"github.com/dop251/goja"
	paho "github.com/eclipse/paho.golang/paho"
	"github.com/mstoykov/k6-taskqueue-lib/taskqueue"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/metrics"
)

// Subscribe to the given topic message will be received using addEventListener
func (c *client) Subscribe(
	// Topic to consume messages from
	topic string,
	// The QoS of messages
	qos,
	// timeout ms
	timeout uint,
) error {
	rt := c.vu.Runtime()
	if c.pahoClient == nil || !c.pahoClient.IsConnected() {
		common.Throw(rt, ErrClient)
		return ErrClient
	}
	c.messageChan = make(chan paho.Message)
	messageCB := func(client paho.Client, msg paho.Message) {
		go func(msg paho.Message) {
			c.messageChan <- msg
		}(msg)
	}
	token := c.pahoClient.Subscribe(topic, byte(qos), messageCB)
	if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
		common.Throw(rt, ErrTimeout)
		return ErrTimeout
	}
	if err := token.Error(); err != nil {
		common.Throw(rt, err)
		return ErrTimeout
	}
	registerCallback := func() func(func() error) {
		callback := c.vu.RegisterCallback()
		return func(f func() error) {
			callback(f)
		}
	}
	c.tq = taskqueue.New(registerCallback)
	go c.loop(c.messageChan, timeout)
	return nil
}

func (c *client) receiveMessageMetric(msgLen float64) error {
	// publish metrics
	now := time.Now()
	state := c.vu.State()
	if state == nil {
		return ErrState
	}

	ctx := c.vu.Context()
	if ctx == nil {
		return ErrState
	}
	metrics.PushIfNotDone(ctx, state.Samples, metrics.Sample{
		TimeSeries: metrics.TimeSeries{Metric: c.metrics.ReceivedMessages, Tags: c.metrics.TagsAndMeta.Tags},
		Time:       now,
		Value:      float64(1),
	})
	metrics.PushIfNotDone(ctx, state.Samples, metrics.Sample{
		TimeSeries: metrics.TimeSeries{Metric: c.metrics.ReceivedBytes, Tags: c.metrics.TagsAndMeta.Tags},
		Time:       now,
		Value:      msgLen,
	})
	return nil
}

//nolint:gocognit // todo improve this
func (c *client) loop(messageChan <-chan paho.Message, timeout uint) {
	ctx := c.vu.Context()
	stop := make(chan struct{})
	defer c.tq.Close()
	for {
		select {
		case msg, ok := <-messageChan:
			if !ok {
				// wanted exit in case of chan close
				return
			}
			c.tq.Queue(func() error {
				payload := string(msg.Payload())
				ev := c.newMessageEvent(msg.Topic(), payload)
				// publish associated metric
				err := c.receiveMessageMetric(float64(len(payload)))
				if err != nil {
					return err
				}
				// TODO authorize multiple listeners
				if c.messageListener != nil {
					if _, err := c.messageListener(ev); err != nil {
						return err
					}
				}
				// if the client is waiting for multiple messages
				// TODO handle multiple // subscribe case
				if c.subRefCount > 0 {
					c.subRefCount--
				} else {
					// exit the handle from evloop async
					stop <- struct{}{}
				}
				return nil
			})
		case <-stop:
			return
		// TODO handle the context better in case of interuption
		case <-ctx.Done():
			c.tq.Queue(func() error {
				ev := c.newErrorEvent("message vu cancel occurred")

				if c.errorListener != nil {
					if _, err := c.errorListener(ev); err != nil {
						// only seen in case of sigint
						return err
					}
				}
				return nil
			})
			return
		case <-time.After(time.Millisecond * time.Duration(timeout)):
			c.tq.Queue(func() error {
				ev := c.newErrorEvent("message timeout")

				if c.errorListener != nil {
					if _, err := c.errorListener(ev); err != nil {
						return err
					}
				}

				return nil
			})
			return
		}
	}
}

// AddEventListener expose the js method to listen for events
func (c *client) AddEventListener(event string, listener func(goja.Value) (goja.Value, error)) {
	switch event {
	case "message":
		c.messageListener = listener
	case "error":
		c.errorListener = listener
	default:
		rt := c.vu.Runtime()
		common.Throw(rt, errors.New("event: "+event+" does not exists"))
	}
}

// SubContinue to be call in message callback to wait for on more message
// be careful this must be called only in the event loop and it not thread safe
func (c *client) SubContinue() {
	c.subRefCount++
}

//nolint:nosnakecase // their choice not mine
func (c *client) newMessageEvent(topic, msg string) *goja.Object {
	rt := c.vu.Runtime()
	o := rt.NewObject()
	must := func(err error) {
		if err != nil {
			common.Throw(rt, err)
		}
	}

	must(o.DefineDataProperty("topic", rt.ToValue(topic), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))
	must(o.DefineDataProperty("message", rt.ToValue(msg), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))
	return o
}
