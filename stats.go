package mqtt

import "github.com/loadimpact/k6/stats"

var (
	ReaderData = stats.New("mqtt.reader.data_recieved", stats.Counter)
	WriterData = stats.New("mqtt.writer.data_sent", stats.Counter)
)
