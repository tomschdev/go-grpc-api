# go-grpc-api
RPC API written in Go.

## Description:
The greet directory holds an implementation for simple greet service comprising of a client, server and protocol buffer definitions. Messages and services are defined in greet/proto/. Messages can be seen as the 'data types' in which the client/server deal. Services can be seen as a collection of interfaces which the server must implement and the client may directly call.

## Execution:
After setting up messages and services in greet/proto/greet.proto, run:
```bash
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/tomschdev/go-grpc-api --go-grpc_out=. --go-grpc_opt=module=github.com/tomschdev/go-grp
c-api greet/proto/greet.proto
```
to generate service and helper functions that enable serialisation/deserialisation and transmission of payloads.

Thereafter, when server and client is implemented, build packages and run their executables:
From greet/server dir:
```bash
go build
```

```bash
./server
```

Repeat for client.

