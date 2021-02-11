/*

This is a k6 test script that imports the xk6-mqtt and
tests Mqtt with a 100 messages per connection.

*/

import {
    check
} from 'k6';
import {
    writer,
    produce,
    reader,
    consume
} from 'k6/x/mqtt'; // import mqtt plugin


export default function () {
    const producer = writer(
        ["localhost:9092"], // bootstrap servers
        "test-k6-plugin-topic", // Mqtt topic
    )

    for (let index = 0; index < 100; index++) {
        let error = produce(producer, "topic", 1,
            ["msg 1",
                "msg 2"]);

        check(error, {
            "is sent": err => err == undefined
        });
    }

    producer.close();

    const consumer = reader(
        ["localhost:9092"], // bootstrap servers
        "test-k6-plugin-topic", // Mqtt topic
    )

    // Read 10 messages only
    let messages = consume(consumer);
    // let messages = consume(consumer, 1);

    // console.log(JSON.stringify(messages));
    // check(messages, {
    //     "10 messages returned": msgs => msgs.length == 10
    // })

    consumer.close();
}
