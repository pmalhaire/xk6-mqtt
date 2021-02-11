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
    consume
} from 'k6/x/mqtt'; // import mqtt plugin

let id = 0;

export default function () {
    const k6Topic = "test-k6-plugin-topic"; // Mqtt topic
    const k6Message = "k6-message-content" // Message content
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
        // timeout in sec
        10,
    )
    //console.log(`produce ${__ITER}`, id);
    let err_produce = produce(producer,
        k6Topic,
        1,
        k6Message,
        10
    );

    check(err_produce, {
        "is sent": err => err == undefined
    });
    if (err_produce != undefined) {
        console.log(`produce err ${__ITER}`, id);
        fail("produce error", err_produce);
    }

    //console.log(`reader ${__ITER}`, id);
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
    //console.log(`consume ${__ITER}`, id);
    // Read messages
    let err_recieve, message = consume(consumer,
        // Topic to consume messages from
        k6Topic,
        // The QoS of messages
        1,
        // timeout in sec
        10,
    );
    //console.log(`consume end ${__ITER}`, id);
    // let messages = consume(consumer, 1);
    check(err_recieve, {
        "is recieved": err => err == undefined
    });
    if (err_recieve != undefined) {
        console.log(`recive err ${__ITER}`, id);
        fail("recive error", err_recieve);
    }

    check(message, {
        "message returned": msg => msg == k6Message
    });
    //console.log(`check done ${__ITER}`, id);


    //console.log(`end ${__ITER}`, id);
}
