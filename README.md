# xk6-mqtt

This is a [k6](https://github.com/loadimpact/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ---------------------------------------------------------------------------------------------------------------------------- |

This project is a k6 extension that can be used to load test Mqtt. Per each connection to Mqtt, many messages can be sent. These messages are an array of strings. There is also a consumer for testing purposes, i.e. to make sure you send the correct data to Mqtt. The consumer is not meant to be used for testing Mqtt under load.

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


running (1m00.0s), 00/50 VUs, 570339 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  1m0s

     ✓ is pub connected
     ✓ is sub connected
     ✓ is subscribed
     ✓ is sent
     ✓ is received
     ✓ is content correct

     █ teardown

     checks...............: 100.00% ✓ 3422034 ✗ 0   
     data_received........: 0 B     0 B/s
     data_sent............: 0 B     0 B/s
     iteration_duration...: avg=5.25ms min=5.96µs med=4.49ms max=59.65ms p(90)=7.39ms p(95)=11.6ms
     iterations...........: 570339  9504.706041/s
     publish_time.........: avg=1.73ms min=0s     med=1ms    max=41ms    p(90)=3ms    p(95)=4ms   
     subscribe_time.......: avg=2.86ms min=0s     med=2ms    max=43ms    p(90)=4ms    p(95)=6ms   
     vus..................: 50      min=50    max=50
     vus_max..............: 50      min=50    max=50

```
