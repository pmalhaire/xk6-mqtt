package mqtt

import "github.com/loadimpact/k6/stats"

var (
	// TODO not implemented yet
	ReaderData = stats.New("mqtt.reader.data_recieved", stats.Counter)
	// TODO not implemented yet
	WriterData = stats.New("mqtt.writer.data_sent", stats.Counter)
)
