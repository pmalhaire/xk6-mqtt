# xk6-mqtt

This is a [k6](https://github.com/loadimpact/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept  at WIP state, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ------------------------------------------------------------------------------------------------------------------------------------------ |

This project is a k6 extension that can be used to load test Mqtt, using a producer. Per each connection to Mqtt, many messages can be sent. These messages are an array of strings. There is also a consumer for testing purposes, i.e. to make sure you send the correct data to Mqtt. The consumer is not meant to be used for testing Mqtt under load. The extension supports producing and consuming messages in Avro format, given a schema for key and/or value.

The real purpose of this extension is not only to test an Mqtt server, but also the system you've designed that uses Apache Mqtt. So, you can test your consumers, and hence your system, by auto-generating messages and sending them to your system via Mqtt.

In order to build the source, you should have the latest version of Go (go1.15) installed. I recommend you to have [gvm](https://github.com/moovweb/gvm) installed.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [gvm](https://github.com/moovweb/gvm)
- [Git](https://git-scm.com/)

Then, install [xk6](https://github.com/k6io/xk6) and build your custom k6 binary with the Mqtt extension:

1. Install `xk6`:
  ```shell
  $ go get -u github.com/k6io/xk6/cmd/xk6
  ```

2. Build the binary: from latest version
  ```shell
  $ xk6 build v0.30.0 --with github.com/pmalhaire/xk6-mqtt
  ```

3. Build the binary: from source code
  ```shell
  git clone github.com/pmalhaire/xk6-mqtt
  cd xk6-mqtt
  $ xk6 build v0.30.0 --with github.com/pmalhaire/xk6-mqtt=.
  ```

## Run & Test

First, you need to have your Mqtt development environment setup.

For example you can use vernemq

```
docker run -p 1883:1883 -e "DOCKER_VERNEMQ_ACCEPT_EULA=yes" -e DOCKER_VERNEMQ_ALLOW_ANONYMOUS=on --name vernemq -d vernemq/vernemq
```

### k6 Test

see [test file](test.js)

You can run k6 with the Mqtt extension using the following command:

```bash
$ ./k6 run --vus 10 --duration 10s test.js
```

And here's the test result output:

```bash

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: test.js
     output: -

  scenarios: (100.00%) 1 scenario, 50 max VUs, 1m30s max duration (incl. graceful stop):
           * default: 50 looping VUs for 1m0s (gracefulStop: 30s)


running (1m00.4s), 00/50 VUs, 6554 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  1m0s

    ✓ is sent
    ✓ 10 messages returned

    checks.........................: 100.00% ✓ 661954 ✗ 0
    data_received..................: 0 B     0 B/s
    data_sent......................: 0 B     0 B/s
    iteration_duration.............: avg=459.31ms min=188.19ms med=456.26ms max=733.67ms p(90)=543.22ms p(95)=572.76ms
    iterations.....................: 6554    108.563093/s
    kafka.reader.dial.count........: 6554    108.563093/s
    kafka.reader.error.count.......: 0       0/s
    kafka.reader.fetches.count.....: 6554    108.563093/s
    kafka.reader.message.bytes.....: 6.4 MB  106 kB/s
    kafka.reader.message.count.....: 77825   1289.124612/s
    kafka.reader.rebalance.count...: 0       0/s
    kafka.reader.timeouts.count....: 0       0/s
    kafka.writer.dial.count........: 6554    108.563093/s
    kafka.writer.error.count.......: 0       0/s
    kafka.writer.message.bytes.....: 54 MB   890 kB/s
    kafka.writer.message.count.....: 655400  10856.309293/s
    kafka.writer.rebalance.count...: 6554    108.563093/s
    kafka.writer.write.count.......: 655400  10856.309293/s
    vus............................: 50      min=50   max=50
    vus_max........................: 50      min=50   max=50
```
