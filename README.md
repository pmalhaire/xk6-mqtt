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

### k6 Test

The [test file](test.js) does the following.


Creates per VU (concurrent clients) :
- one topic
- one subscribe connection (done at first iteration)
- one publish connection (done at first iteration)

Per iteration :
Subscribe to topic
Publish to topic
Read and check message result

You can run the test using the following command:

```bash
./k6 run --vus 50 --duration 1m test.js
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


running (1m00.0s), 00/50 VUs, 326051 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  1m0s

     ✓ is pub connected
     ✓ is sub connected
     ✓ is subscribed
     ✓ is sent
     ✓ is received
     ✓ is content correct

     █ teardown

     checks...............: 100.00% ✓ 1956306     ✗ 0
     data_received........: 14 MB   238 kB/s
     data_sent............: 14 MB   238 kB/s
     iteration_duration...: avg=9.19ms min=9.17µs med=8.39ms max=141.6ms p(90)=12.8ms p(95)=16.3ms
     iterations...........: 326051  5433.374851/s
     publish_time.........: avg=3.23ms min=0s     med=3ms    max=41ms    p(90)=5ms    p(95)=6ms
     subscribe_time.......: avg=4.94ms min=0s     med=5ms    max=51ms    p(90)=7ms    p(95)=9ms
     vus..................: 50      min=50        max=50
     vus_max..............: 50      min=50        max=50


```
