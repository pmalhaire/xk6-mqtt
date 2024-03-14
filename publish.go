package mqtt

import (
	// "context"
	// "fmt"
	"time"

	"github.com/dop251/goja"
	// paho "github.com/eclipse/paho.golang/paho"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/metrics"
)

// Publish allow to publish one message
//
//nolint:gocognit


func (c *client) publishSync(
	topic string,
	qos int,
	message string,
) error {
	// if c.connectionManager == nil {
	// 	rt := c.vu.Runtime()
	// 	common.Throw(rt, ErrClient)
	// 	return ErrClient
	// }


	err := c.publishMessageMetric(float64(len(message)))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) publishMessageMetric(msgLen float64) error {
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
		TimeSeries: metrics.TimeSeries{Metric: c.metrics.SentMessages, Tags: c.metrics.TagsAndMeta.Tags},
		Time:       now,
		Value:      float64(1),
	})
	metrics.PushIfNotDone(ctx, state.Samples, metrics.Sample{
		TimeSeries: metrics.TimeSeries{Metric: c.metrics.SentBytes, Tags: c.metrics.TagsAndMeta.Tags},
		Time:       now,
		Value:      msgLen,
	})
	return nil
}

//nolint:nosnakecase // their choice not mine
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