/*

This is a k6 test script that imports the xk6-mqtt and
tests Mqtt with a 100 messages per connection.

*/

import {
    check
} from 'k6';

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


const rnd_count = 2000;
// create random number to create a new topic at each run
let rnd = Math.random() * rnd_count;

// keep connection made by VU
let vus_connections = {}

// default timeout (ms)
let timeout = 2000

export default function () {
    // Mqtt topic one per VU
    const k6Topic = `test-k6-plugin-topic ${rnd} ${__VU}`;
    // Message content one per ITER and per VU
    const k6Message = `k6-message-content-${rnd} ${__VU}:${__ITER}`;
    const k6SubId = `k6-sub-${__VU}`;
    const k6PubId = `k6-pub-${__VU}`;

    let err_pub_client, pub_client;
    const host = "localhost";
    const port = "1883";
    // use one connection per vu
    if (k6PubId in vus_connections) {
        pub_client = vus_connections[k6PubId];
    } else {
        try {
            pub_client = connect(
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
            vus_connections[k6PubId] = pub_client;
        } catch (error) {
            err_pub_client = error;
        }
    }
    check(err_pub_client, {
        "is pub connected": err => err == undefined
    });

    let err_sub_client, sub_client;
    // use one connection per vu
    if (k6SubId in vus_connections) {
        sub_client = vus_connections[k6SubId];
    } else {
        try {
        sub_client = connect(
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
        vus_connections[k6SubId] = sub_client;
        } catch (error) {
            err_sub_client = error;
        }
    }
    check(err_sub_client, {
        "is sub connected": err => err == undefined
    });

    // subscribe first
    let err_subscribe, consume_token;
    try {
        consume_token = subscribe(
            // consume object
            sub_client,
            // topic to be used
            k6Topic,
            // The QoS of messages
            1,
            // timeout in ms
            timeout,
        )
    } catch (error) {
        err_subscribe = error
    }
    check(err_subscribe, {
        "is subscribed": err => err == undefined
    });
    // publish message
    let err_publish;

    try {
        publish(
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
            timeout,
        );
    } catch (error) {
        err_publish = error
    }

    check(err_publish, {
        "is sent": err => err == undefined
    });
    let err_consume, message
    try {
        // Read one message
        message = consume(
            // token to recieve message
            consume_token,
            // timeout in ms
            timeout,
        );
    } catch (error) {
        err_consume = error
    }

    check(err_consume, {
        "is received": err => err == undefined
    });

    check(message, {
        "is content correct": msg => msg == k6Message
    });

}

export function teardown() {
    for (client in vus_connections) {
        close(client, timeout);
    }
}
