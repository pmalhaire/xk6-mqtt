/*

This is a k6 test script that imports the xk6-mqtt and
tests Mqtt with a 100 messages per connection.

*/

import {
    check, fail
} from 'k6';
import { TLS_1_0 } from 'k6/http';
import {
    // connect to mqtt
    connect,
    // close connection
    close,
    // subscribe to topic
    subscribe,
    // consume one message
    consume,
    // publish one message
    publish,
} from 'k6/x/mqtt'; // import mqtt plugin

// create random number to create a new topic at each run
let rnd = Math.random() * 1000;

let vus_connections = {}


export default function () {
    // Mqtt topic one per VU
    const k6Topic = `test-k6-plugin-topic ${rnd} ${__VU}`;
    // Message content one per ITER and per VU
    const k6Message = `k6-message-content-${rnd} ${__VU}:${__ITER}`;
    const k6SubId = `k6-sub-${__VU}`;
    const k6PubId = `k6-pub-${__VU}`;

    let pub_client;
    // use one connection per vu
    if (k6PubId in vus_connections) {
        pub_client = vus_connections[k6PubId];
    } else {
        pub_client = connect(
            // The list of URL of  MQTT server to connect to
            ["localhost:1883"],
            // A username to authenticate to the MQTT server
            "",
            // Password to match username
            "",
            // clean session setting
            false,
            // Client id for reader
            k6PubId,
            // timeout in ms
            1000,
        )
        vus_connections[k6PubId] = pub_client;
    }

    let sub_client;
    // use one connection per vu
    if (k6SubId in vus_connections) {
        sub_client = vus_connections[k6SubId];
    } else {
        sub_client = connect(
            // The list of URL of  MQTT server to connect to
            ["localhost:1883"],
            // A username to authenticate to the MQTT server
            "",
            // Password to match username
            "",
            // clean session setting
            false,
            // Client id for reader
            k6SubId,
            // timeout in ms
            1000,
        )
        vus_connections[k6SubId] = sub_client;
    }

    // subscribe first
    let err_subscribe, consume_token = subscribe(
        // consume object
        sub_client,
        // topic to be used
        k6Topic,
        // The QoS of messages
        1,
        // timeout in ms
        1000,
    )
    check(err_subscribe, {
        "is subscribed": err => err == undefined
    });
    // publish message
    let err_publish = publish(
        // producer object
        pub_client,
        // topic to be used
        k6Topic,
        // The QoS of messages
        1,
        // Message to be sent
        k6Message,
        // retain policy on message
        false,
        // timeout in ms
        1000,
    );

    check(err_publish, {
        "is sent": err => err == undefined
    });

    // Read one message
    let err_consume, message = consume(
        // token to recieve message
        consume_token,
        // timeout in ms
        1000,
    );

    check(err_consume, {
        "is recieved": err => err == undefined
    });
    if (err_consume != undefined) {
        console.log(`recive err ${__ITER} `, id);
    }

    check(message, {
        "is content correct": msg => msg == k6Message
    });
    if (message != k6Message) {
        console.log(`received ${message} expected ${k6Message}`);
    }

}

export function teardown() {
    for (client in vus_connections) {
        close(client, 1000);
    }
}