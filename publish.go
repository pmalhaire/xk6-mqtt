package mqtt

import (
	"context"
	"fmt"
	"time"

	paho "github.com/eclipse/paho.golang/paho"
	"go.k6.io/k6/metrics"
)

func (c *client) Publish(
	topic string,
	qos int,
	message []byte,
	retain bool,
	timeout uint,
) error {
	ctx := context.Background()

	_, publish_error := c.connectionManager.Publish(ctx, &paho.Publish{
		QoS:     1,
		Topic:  topic,
		Payload: message,
	})
	if (publish_error != nil) {
		fmt.Println("error publishing message: ", publish_error, " to topic: ", topic, " for message: ", message)
		return publish_error
	}
	return nil
}

func (c *client) publishSync(
	topic string,
	qos int,
	message string,
) error {
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