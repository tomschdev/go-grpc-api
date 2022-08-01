# go-grpc-api
Collection of RPC APIs written in Go during the completion of Udemy course: Go & gRPC Masterclass. \
The repository holds multiple APIs, each implemented with the purpose of demonstrating a certain feature of gRPC. \
The Description and Execution sections below hold in the general case. \
The code held in most directories is self-explanatory, by directory name. \
**advanced** contains APIs under **errorhandle** (demo of error handling in gRPC), **deadline** (demo of timeouts enforced by client), and **ssl** (demo of proper ssl security implementation with certificate generation).\
**blog** holds code that implements a CRUD API using MongoDB. This involves server and client implementations as well as a docker-compose which allows an instance of MongoDB to be deployed in a docker container and accessed via localhost.\

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



