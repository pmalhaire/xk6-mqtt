// NOTE : all tests here suppose a running server
package mqtt

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/require"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/eventloop"
	"go.k6.io/k6/js/modulestest"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/testutils/httpmultibin"
	"go.k6.io/k6/metrics"
	"gopkg.in/guregu/null.v3"
)

// test server params
const host = "localhost"

const (
	port    = "1883"
	timeout = "2000"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		//nolint:gosec
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type testState struct {
	rt      *goja.Runtime
	tb      *httpmultibin.HTTPMultiBin
	state   *lib.State
	samples chan metrics.SampleContainer
	ev      *eventloop.EventLoop
}

func newTestState(t testing.TB) testState {
	tb := httpmultibin.NewHTTPMultiBin(t)

	root, err := lib.NewGroup("", nil)
	require.NoError(t, err)

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper{})

	samples := make(chan metrics.SampleContainer, 1000)

	state := &lib.State{
		Group:  root,
		Dialer: tb.Dialer,
		Options: lib.Options{
			SystemTags: metrics.NewSystemTagSet(
				metrics.TagURL,
				metrics.TagProto,
				metrics.TagStatus,
				metrics.TagSubproto,
			),
			UserAgent: null.StringFrom("TestUserAgent"),
		},
		Samples:        samples,
		TLSConfig:      tb.TLSClientConfig,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(metrics.NewRegistry()),
		Tags:           lib.NewTagMap(nil),
	}

	vu := &modulestest.VU{
		CtxField:     tb.Context,
		InitEnvField: &common.InitEnvironment{},
		RuntimeField: rt,
		StateField:   state,
	}
	m := new(RootModule).NewModuleInstance(vu)
	require.NoError(t, rt.Set("mqtt", m.Exports().Named))
	ev := eventloop.New(vu)
	vu.RegisterCallbackField = ev.RegisterCallback

	return testState{
		rt:      rt,
		tb:      tb,
		state:   state,
		samples: samples,
		ev:      ev,
	}
}

func TestBasic(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	rndStr := RandStringRunes(10)
	str := `const k6Topic = "gotest-k6-topic-` + rndStr + `";
	const k6SubId = "gotest-k6-SubID-` + rndStr + `";
	const k6PubId = "gotest-k6-PubID-` + rndStr + `";

	const host = "` + host + `";
	const port = "` + port + `";
	const timeout = ` + timeout + `;
	let err;

	let client = new mqtt.Client(
		// The list of URL of  MQTT server to connect to
		[host + ":" + port],
		// A username to authenticate to the MQTT server
		"",
		// Password to match username
		"",
		// clean session setting
		false,
		// Client id for reader
		k6PubId,
		// timeout in ms
		timeout,
	)

	try {
		client.connect()
	} catch (error) {
			err = error
	}

	if (err != undefined) {
		throw new Error("Unexpected client connect error:", err)
	}

	try {
		client.close(timeout)
	} catch (error) {
			err = error
	}

	if (err != undefined) {
		throw new Error("expected close error")
	}
	`

	err := ts.ev.Start(func() error {
		_, err := ts.rt.RunString(str)
		return err
	})
	require.NoError(t, err, "please check that your server is running")
}

func TestSubscribeNoPublish(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	rndStr := RandStringRunes(10)
	str := `const k6Topic = "gotest-k6-topic-` + rndStr + `";
	const k6SubId = "gotest-k6-SubID-` + rndStr + `";
	const k6PubId = "gotest-k6-PubID-` + rndStr + `";

	const host = "` + host + `";
	const port = "` + port + `";
	const timeout = ` + timeout + `;
	let err;

	let client = new mqtt.Client(
		// The list of URL of  MQTT server to connect to
		[host + ":" + port],
		// A username to authenticate to the MQTT server
		"",
		// Password to match username
		"",
		// clean session setting
		false,
		// Client id for reader
		k6PubId,
		// timeout in ms
		timeout,
	)

	try {
		client.connect()
	} catch (error) {
			err = error
	}

	if (err != undefined) {
		throw new Error("Unexpected client connect error:", err)
	}

	client.subscribe(
		// topic to be used
		k6Topic,
		// The QoS of messages
		1,
		// timeout in ms
		timeout,
	)

	try {
		client.close(timeout)
	} catch (error) {
			err = error
	}

	if (err != undefined) {
		throw new Error("expected close error")
	}
	`

	err := ts.ev.Start(func() error {
		_, err := ts.rt.RunString(str)
		return err
	})
	require.NoError(t, err, "please check that your server is running")
}

func TestBasicErr(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	sr := ts.tb.Replacer.Replace
	err := ts.ev.Start(func() error {
		_, err := ts.rt.RunString(sr(`
		const k6Topic = "dummy";
		const k6SubId = "dummy";
		const k6PubId = "dummy";

		const host = "dummy";
		const port = "1883";
		const timeout = ` + timeout + `;

		let err;
		// create client
		let client = new mqtt.Client(
			// The list of URL of  MQTT server to connect to
			[host + ":" + port],
			// A username to authenticate to the MQTT server
			"",
			// Password to match username
			"",
			// clean session setting
			false,
			// Client id for reader
			k6PubId,
			// timeout in ms
			timeout
		)

		try {
			client.connect()
		} catch (error) {
				err = error
		}

		if ( !err ) {
			throw new Error("expected connect error")
		}
	`))
		return err
	})
	require.NoError(t, err)
}

func TestPubSub(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	rndStr := RandStringRunes(10)
	err := ts.ev.Start(func() error {
		_, err := ts.rt.RunString(
			`
const k6Topic = "gotest-k6-topic-` + rndStr + `";
const k6SubId = "gotest-k6-SubID-` + rndStr + `";
const k6PubId = "gotest-k6-PubID-` + rndStr + `";
const k6Message = "gotest-k6-message";

const host = "` + host + `";
const port = "` + port + `";
const timeout = ` + timeout + `;
const messageCount = 2;


// create publisher client
let publisher = new mqtt.Client(
    // The list of URL of  MQTT server to connect to
    [host + ":" + port],
    // A username to authenticate to the MQTT server
    "",
    // Password to match username
    "",
    // clean session setting
    false,
    // Client id for reader
    k6PubId,
    // timeout in ms
    timeout,
)

publisher.connect()


// create subscriber client
let subscriber = new mqtt.Client(
    // The list of URL of  MQTT server to connect to
    [host + ":" + port],
    // A username to authenticate to the MQTT server
    "",
    // Password to match username
    "",
    // clean session setting
    false,
    // Client id for reader
    k6SubId,
    // timeout in ms
    timeout,
)

subscriber.connect()


// subscribe first
subscriber.subscribe(
	// topic to be used
	k6Topic,
	// The QoS of messages
	1,
	// timeout in ms
	timeout,
)

let count = messageCount;
subscriber.addEventListener("message", (obj) => {
	// closing as we received one message
	let message = obj.message
	if (message != k6Message ) {
		throw new Error("unexpected message content")
	}
	// tell the event listener to wait for more messages
	// remove this if you want to send only one message
	if (--count > 0) {
		subscriber.subContinue();
	} else {
		subscriber.close(timeout);
		publisher.close(timeout);
	}
})

subscriber.addEventListener("error", (err) => {
	throw new Error("message error");
})

for (let i = 0; i < messageCount; i++) {
	// publish count messages
	publisher.publish(
		// topic to be used
		k6Topic,
		// The QoS of messages
		1,
		// Message to be sent
		k6Message,
		// retain policy on message
		false,
		// timeout in ms
		timeout,
	);

}
`)
		return err
	})
	require.NoError(t, err)
}

func TestPubAsyncSub(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	rndStr := RandStringRunes(10)
	err := ts.ev.Start(func() error {
		_, err := ts.rt.RunString(
			`
const k6Topic = "gotest-k6-topic-` + rndStr + `";
const k6SubId = "gotest-k6-SubID-` + rndStr + `";
const k6PubId = "gotest-k6-PubID-` + rndStr + `";
const k6Message = "gotest-k6-message";

const host = "` + host + `";
const port = "` + port + `";
const timeout = ` + timeout + `;
const messageCount = 3;


// create publisher client
let publisher = new mqtt.Client(
    // The list of URL of  MQTT server to connect to
    [host + ":" + port],
    // A username to authenticate to the MQTT server
    "",
    // Password to match username
    "",
    // clean session setting
    false,
    // Client id for reader
    k6PubId,
    // timeout in ms
    timeout,
)

publisher.connect()


// create subscriber client
let subscriber = new mqtt.Client(
    // The list of URL of  MQTT server to connect to
    [host + ":" + port],
    // A username to authenticate to the MQTT server
    "",
    // Password to match username
    "",
    // clean session setting
    false,
    // Client id for reader
    k6SubId,
    // timeout in ms
    timeout,
)

subscriber.connect()


// subscribe first
subscriber.subscribe(
	// topic to be used
	k6Topic,
	// The QoS of messages
	1,
	// timeout in ms
	timeout,
)

let count = messageCount;
subscriber.addEventListener("message", (obj) => {
	// closing as we received one message
	let message = obj.message
	if (message != k6Message ) {
		throw new Error("unexpected message content")
	}
	// tell the event listener to wait for more messages
	// remove this if you want to send only one message
	if (--count > 0) {
		subscriber.subContinue();
	} else {
		subscriber.close(timeout);
		publisher.close(timeout);
	}
})

subscriber.addEventListener("error", (err) => {
	throw new Error("message error");
})


// publish success and failure
publisher.publish(
	// topic to be used
	k6Topic,
	// The QoS of messages
	1,
	// Message to be sent
	k6Message,
	// retain policy on message
	false,
	// timeout in ms
	timeout,
	(obj) => {
		if (!obj.type || !obj.topic) {
			throw new Error("unexpected event");
		}
	},
	(err) => {
		if (!err.type || !err.message) {
			throw new Error("unexpected error", err);
		}
	}
);

// publish success only
publisher.publish(
	// topic to be used
	k6Topic,
	// The QoS of messages
	1,
	// Message to be sent
	k6Message,
	// retain policy on message
	false,
	// timeout in ms
	timeout,
	(obj) => {
		if (!obj.type || !obj.topic) {
			throw new Error("unexpected event");
		}
	}
);

// publish error only
publisher.publish(
	// topic to be used
	k6Topic,
	// The QoS of messages
	1,
	// Message to be sent
	k6Message,
	// retain policy on message
	false,
	// timeout in ms
	timeout,
	null,
	(err) => {
		if (!err.type || !err.message) {
			throw new Error("unexpected error");
		}
	}
);

publisher.close(timeout);
subscriber.close(timeout);

// create failer client
let failer = new mqtt.Client(
    // The list of URL of  MQTT server to connect to
    [host + ":" + port],
    // A username to authenticate to the MQTT server
    "",
    // Password to match username
    "",
    // clean session setting
    false,
    // Client id for reader
    k6PubId,
    // timeout in ms
    timeout,
)

// publish test failure
failer.publish(
	// topic to be used
	k6Topic,
	// The QoS of messages
	1,
	// Message to be sent
	k6Message,
	// retain policy on message
	false,
	// timeout in ms
	timeout,
	null,
	(err) => {
		if (!err.type || !err.message) {
			throw new Error("unexpected error", err);
		}
	}
);
`)
		return err
	})
	require.NoError(t, err)
}

func BenchmarkLoop(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ts := newTestState(t)
		rndStr := RandStringRunes(10)
		err := ts.ev.Start(func() error {
			_, err := ts.rt.RunString(
				`
const k6Topic = "gotest-k6-topic-` + rndStr + `";
const k6SubId = "gotest-k6-SubID-` + rndStr + `";
const k6PubId = "gotest-k6-PubID-` + rndStr + `";
const k6Message = "gotest-k6-message";

const host = "` + host + `";
const port = "` + port + `";
const timeout = ` + timeout + `;


let publisher = new mqtt.Client(
		// The list of URL of  MQTT server to connect to
		[host + ":" + port],
		// A username to authenticate to the MQTT server
		"",
		// Password to match username
		"",
		// clean session setting
		false,
		// Client id for reader
		k6PubId,
		// timeout in ms
		timeout,
	)

publisher.connect()

let subscriber = new mqtt.Client(
	// The list of URL of  MQTT server to connect to
	[host + ":" + port],
	// A username to authenticate to the MQTT server
	"",
	// Password to match username
	"",
	// clean session setting
	false,
	// Client id for reader
	k6SubId,
	// timeout in ms
	timeout,
)

subscriber.connect()

subscriber.subscribe(
		// topic to be used
		k6Topic,
		// The QoS of messages
		1,
		// timeout in ms
		timeout,
	)

subscriber.addEventListener("message", (msg) => {
		let message = msg.payload
	})
subscriber.addEventListener("error", (err) => {
		throw new Error("message not received error:", err)
	})

publisher.publish(
	// topic to be used
	k6Topic,
	// The QoS of messages
	1,
	// Message to be sent
	k6Message,
	// retain policy on message
	false,
	// timeout in ms
	timeout,
);
				`)
			return err
		})
		require.NoError(t, err)
	}
}
