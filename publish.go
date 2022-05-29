package mqtt

import (
	"time"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
)

// Publish allow to publish one message
//nolint:gocognit
func (c *client) Publish(
	topic string,
	qos int,
	message string,
	retain bool,
	timeout uint,
	success func(goja.Value) (goja.Value, error),
	failure func(goja.Value) (goja.Value, error),
) error {
	if success == nil && failure == nil {
		if c.pahoClient == nil || !c.pahoClient.IsConnected() {
			rt := c.vu.Runtime()
			common.Throw(rt, ErrClient)
			return ErrClient
		}
		token := c.pahoClient.Publish(topic, byte(qos), retain, message)
		// sync case
		if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
			rt := c.vu.Runtime()
			common.Throw(rt, ErrTimeout)
			return ErrTimeout
		}
		if err := token.Error(); err != nil {
			rt := c.vu.Runtime()
			common.Throw(rt, ErrPublish)
			return ErrPublish
		}
		return nil
	}
	// async case
	callback := c.vu.RegisterCallback()
	go func() {
		if c.pahoClient == nil || !c.pahoClient.IsConnected() {
			callback(func() error {
				ev := c.newErrorEvent("publish not connected")
				if failure != nil {
					if _, err := failure(ev); err != nil {
						return err
					}
				}
				return nil
			})
			return
		}
		token := c.pahoClient.Publish(topic, byte(qos), retain, message)
		if !token.WaitTimeout(time.Duration(timeout) * time.Millisecond) {
			callback(func() error {
				ev := c.newErrorEvent("publish timeout")

				if failure != nil {
					if _, err := failure(ev); err != nil {
						return err
					}
				}
				return nil
			})
			return
		}
		if err := token.Error(); err != nil {
			callback(func() error {
				ev := c.newErrorEvent(err.Error())
				if failure != nil {
					if _, err := failure(ev); err != nil {
						return err
					}
				}
				return nil
			})
			return
		}
		callback(func() error {
			ev := c.newPublishEvent(topic)
			if success != nil {
				if _, err := success(ev); err != nil {
					return err
				}
			}
			return nil
		})
	}()
	return nil
}

func (c *client) newPublishEvent(topic string) *goja.Object {
	rt := c.vu.Runtime()
	o := rt.NewObject()
	must := func(err error) {
		if err != nil {
			common.Throw(rt, err)
		}
	}

	must(o.DefineDataProperty("type", rt.ToValue("publish"), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))
	must(o.DefineDataProperty("topic", rt.ToValue(topic), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_TRUE))
	return o
}
