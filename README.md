# zipkin-tracing-test
Test out OpenTracing with Zipkin

## Dependencies

golang, docker

## Setup

1) Run `make setup` to download the zipkin docker-compose file from the [github repo](https://github.com/openzipkin/docker-zipkin)
2) Edit `docker-compose.yml` to include the following section under `services`:

 ```
 testserver:
   build: ./
   image: testserver:v1
   container_name: testserver
   ports:
     - "3000:3000"
```

3) Run `make start`. Navigate in your browser to the appropriate ip address + port (9411) to see your traces. If you are using docker-machine + docker (as I do) this should be:

`docker-machine ip dev`

Then navigate to the given ip_address:9411.
