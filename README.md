# gRPC PoC
This is a simple dummy service that implements both gRPC and JSON to compare them.

```sh
make run-grpc-server # start the gRPC server
```

Then on another terminal, run the client:
```shell
make run-golang-client
```

There is also a simple PHP client implemented:
```shell
make run-php-client
```

If you want to compare implementations, there is also a functional JSON HTTP server:
```shell
make run-http-server
```
