import * as VehicleStateProtos from './generated/tools/proto/hive/ota/vehicle_state_pb';

const mqtt = require('k6/x/mqtt');

const rnd_count = 2000;

// create random number to create a new topic at each run
let rnd = Math.random() * rnd_count;

// conection timeout (ms)
let connectTimeout = 2000

// publish timeout (ms)
let publishTimeout = 2000

// subscribe timeout (ms)
let subscribeTimeout = 2000

// connection close timeout (ms)
let closeTimeout = 2000

// Mqtt topic one per VU
const k6Topic = "vehicle_state_ota/820fb0a0-9966-4dcf-9745-028bd047fde7";
// Connect IDs one connection per VU
const k6SubId = `k6-sub-${rnd}-${__VU}`;
const k6PubId = `k6-pub-${rnd}-${__VU}`;

// number of message pusblished and receives at each iteration
const messageCount = 3;

const host = "mqtts://mqtt.mosaic.staging.applied.dev";
const port = "8883";


// create publisher client
console.log("in test.js, creating client")
let publisher = new mqtt.Client(
    // The list of URL of  MQTT server to connect to
    [host + ":" + port],
    // A username to authenticate to the MQTT server
    "admin-user",
    // Password to match username
    "REDACTED",
    // clean session setting
    false,
    // Client id for reader
    k6PubId,
    // timeout in ms
    connectTimeout,
)
let err;

const myVehicleState = new VehicleStateProtos.VehicleState();
myVehicleState.setDoorsLocked(false);
const spoiler_state = new VehicleStateProtos.SpoilerState();
spoiler_state.setPosition(7);
myVehicleState.setSpoilerState(spoiler_state);

const date = new proto.google.protobuf.Timestamp()
let curr_date = new Date()
date.fromDate(curr_date)
console.log("date is: ", curr_date)
myVehicleState.setTimestampUtc(date);

try {
    console.log("in test.js connecting to broker")
    publisher.connect()
}
catch (error) {
    err = error
}

export default function () {
    for (let i = 0; i < messageCount; i++) {
        // publish count messages
        let err_publish;
        // console.log("in test.js, will publish message")
        try {
            publisher.publish(
                // topic to be used
                k6Topic,
                // The QoS of messages
                1,
                // Message to be sent
                myVehicleState.serializeBinary(),
                // retain policy on message
                false,
                // timeout in ms
                publishTimeout,
            );
        } catch (error) {
            console.log("We failed to publish!: ", error)
            err_publish = error
        }
    }
}

export function teardown() {
    publisher.close(closeTimeout);
}
