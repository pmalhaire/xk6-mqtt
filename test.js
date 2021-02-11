/*

This is a k6 test script that imports the xk6-mqtt and
tests Mqtt with a 100 messages per connection.

*/

import {
    check, fail
} from 'k6';
import {
    writer,
    produce,
    reader,
    consume,
    subscribe
} from 'k6/x/mqtt'; // import mqtt plugin

let seed = Math.random() * 1000;
let id = seed;

export default function () {
    const k6Topic = "test-k6-plugin-topic" + id; // Mqtt topic
    const k6Message = "k6-message-content" + id; // Message content
    const k6SubId = "k6-sub" + id++
    const k6PubId = "k6-pub" + id++
    //console.log(`writer ${__ITER}`, id);
    const producer = writer(
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
        // retain or not the message if no consumer subscribed
        true,
        // timeout in sec
        10,
    )

    const consumer = reader(
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
        // timeout in sec
        10,
    )

    // subscribe first
    let err_subscribe, recieved = subscribe(
        // consume object
        consumer,
        // topic to be used
        k6Topic,
        // The QoS of messages
        1,
        // timeout in sec
        10,
    )
    check(err_subscribe, {
        "is subscribed": err => err == undefined
    });
    // produce message
    let err_produce = produce(
        // producer object
        producer,
        // topic to be used
        k6Topic,
        // The QoS of messages
        1,
        // Message to be sent
        k6Message,
        // retain policy on message
        true,
        // timeout in sec
        10
    );

    check(err_produce, {
        "is sent": err => err == undefined
    });
    if (err_produce != undefined) {
        console.log(`produce err ${__ITER}`, id);
    }
    // Read messages
    let err_recieve, message = consume(
        consumer,
        // token to recieve message
        recieved,
        // timeout in sec
        10,
    );

    check(err_recieve, {
        "is recieved": err => err == undefined
    });
    if (err_recieve != undefined) {
        console.log(`recive err ${__ITER}`, id);
    }

    check(message, {
        "message returned": msg => msg == k6Message
    });
}
