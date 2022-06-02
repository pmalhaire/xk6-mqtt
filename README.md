# xk6-mqtt

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ---------------------------------------------------------------------------------------------------------------------------- |

This project is a k6 extension that can be used to load test Mqtt. Per each connection to Mqtt, many messages can be sent. These messages are an array of strings. There is also a consumer for testing purposes, i.e. to make sure you send the correct data to Mqtt. The consumer is not meant to be used for testing Mqtt under load.

In order to build the source, you should have the latest version of Go (go1.15) installed. I recommend you to have [gvm](https://github.com/moovweb/gvm) installed.

## k6 version

This extension is uptodate with k6 version v0.38.0 (breaking change) previous version is tagged at tag v0.37.0

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [gvm](https://github.com/moovweb/gvm)
- [Git](https://git-scm.com/)

Then, install [xk6](https://github.com/grafana/xk6) and build your custom k6 binary with the Mqtt extension:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary: from latest version
  ```shell
  $ xk6 build --with github.com/pmalhaire/xk6-mqtt@latest
  ```

3. Build the binary: from source code
  ```shell
  git clone github.com/pmalhaire/xk6-mqtt
  cd xk6-mqtt
  $ xk6 build --with github.com/pmalhaire/xk6-mqtt=.
  ```

## Run & Test

First, you need to have your Mqtt development environment setup.

For example you can use vernemq

```
docker run -p 1883:1883 -e "DOCKER_VERNEMQ_ACCEPT_EULA=yes" -e DOCKER_VERNEMQ_ALLOW_ANONYMOUS=on --name vernemq -d vernemq/vernemq
```

or Mosquitto

```
docker run --name mosquitto -d -p 1883:1883 eclipse-mosquitto mosquitto -c /mosquitto-no-auth.conf
```

### k6 Test

To use this extension write you test as a js file :

The sample [test file](examples/test.js) does the following.

Creates per VU (concurrent clients) :
- one topic
- one subscribe connection
- one publish connection

Per iteration :
- Subscribe to topic
- Publish to topic `messageCount` messages
- Read and check message `messageCount` results

You can run the test using the following command:

```bash
./k6 run --vus 50 --duration 1m examples/test.js
```

Sample test result output:

```bash

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: examples/test.js
     output: -

  scenarios: (100.00%) 1 scenario, 50 max VUs, 1m30s max duration (incl. graceful stop):
           * default: 50 looping VUs for 1m0s (gracefulStop: 30s)


running (1m00.0s), 00/50 VUs, 148575 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  1m0s

     ✓ is publisher connected
     ✓ is subcriber connected
     ✓ is sent
     ✓ message received
     ✓ is content correct

     █ teardown

     checks...............: 100.00% ✓ 1634325     ✗ 0
     data_received........: 0 B     0 B/s
     data_sent............: 0 B     0 B/s
     iteration_duration...: avg=20.18ms min=149.12µs med=18.92ms max=100.09ms p(90)=27.41ms p(95)=33.05ms
     iterations...........: 148575  2475.613925/s
     vus..................: 50      min=50        max=50
     vus_max..............: 50      min=50        max=50


```
