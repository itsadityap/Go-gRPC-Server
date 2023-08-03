## Go - gRPC Introduction

## Tech Stack
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)


### What is gRPC?
gRPC is a modern, open source, high-performance remote procedure call (RPC) framework that can run anywhere.

### What is an RPC?
A remote procedure call is when a computer program causes a procedure (subroutine) to execute in a different address space (commonly on another computer on a shared network), which is coded as if it were a normal (local) procedure call, without the programmer explicitly coding the details for the remote interaction. That is, the programmer writes essentially the same code whether the subroutine is local to the executing program, or remote. This is a form of client–server interaction (caller is client, executor is server), typically implemented via a request–response message-passing system. In the object-oriented programming paradigm, RPCs are represented by remote method invocation (RMI). The RPC model implies a level of location transparency, namely that calling procedures is largely the same whether it is local or remote, but usually they are not identical, so local calls can be distinguished from remote calls. Remote calls are usually orders of magnitude slower and less reliable than local calls, so distinguishing them is important.

### What is Protocol Buffers?
Protocol buffers are a language-neutral, platform-neutral extensible mechanism for serializing structured data. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages. You can even update your data structure without breaking deployed programs that are compiled against the "old" format.

### What is a gRPC Server?
A gRPC server is a server that implements RPCs using the gRPC framework. A gRPC server is created like any other server, but it uses a gRPC server instance to handle RPCs instead of a normal HTTP handler.

### What is a gRPC Client?
A gRPC client is a client that implements RPCs using the gRPC framework. A gRPC client is created like any other client, but it uses a gRPC client instance to handle RPCs instead of a normal HTTP client.

### What is a gRPC Stub?
A gRPC stub is a client-side interface that is used to make RPC calls to a gRPC server. A gRPC stub is created from a gRPC client instance and is used to call RPC methods on the server.

### What is a gRPC Service?
A gRPC service is a service that implements RPCs using the gRPC framework. A gRPC service is created like any other service, but it uses a gRPC service instance to handle RPCs instead of a normal HTTP service.


## Installation in Windows
### Install Go
1. Download the Go installer from the [Downloads page](https://golang.org/dl/).
2. Open the installer and follow the prompts to install Go.
3. Verify that you've installed Go by opening a command prompt and typing the following command:
```bash
$ go version
```
4. Confirm that the command prints the installed version of Go.
