# fs3

CSE 453 Data Center Systems final project
Assaf Vayner (assafv) + Tom Wu (zw237)

## Summary
A distributed remote file system with a grpc interface and command line tool to access it.

## Intended Runtime Configuration
The server nodes for this project are intended to be managed via docker swarm and ran on cloudlab servers.
While we do backup files we do not currently support a mechanism to make the backup node into the primary node and support allowing the backup to fail indefinitely.
The client code is also intended to be ran from cloudlab nodes so that they have access to communication without the public internet.

The client consists of 2 completely separate components, firstly a cli tool to make individual requests as proof of concept of the initial intended usage.
Secondly a python program used to make grpc calls and measure their duration as a method of analyzing performance.

## Instructions to run server
Tested on wsl/ubuntu mounted, run the following commands (you will need docker installed)
- `sudo make build_p`
- `sudo make build_b`
- `sudo make up`

## Performance analysis
We intend to analyze performance by varying frequency of requests, and the size of request/response payloads.
We will attempt to run the server application on containers running on bare metal as well as vms with virtualized storage i/o paths so as to showcase the performance gain of getting direct access to storage hardware.

### Network performance analysis
We may attempt to test network performance drops by measuring RTT's per forwarding requests.
We will keep a request Id for each forwarding request, and record the total time of the forwarding request on the server and the total time of processing the request on the backup, then by subtracting the latter from the former we will be able to gauge the time of packets on the wire (barring de/serialization time taken by grpc libraries).
The we will be able to visualize the performance drops based on network performance from aggregating the data from both nodes.

## Running the applications in development mode
### Server
To build the server run `make server` which will generate the `server` executable in the `server` directory (`server/server`).

To run the primary run: `make primary`

To run the backup run: `make backup`

In local development mode you'll want to save the files in the `data` directory as opposed to the privileged `/data` directory. To do this add `stage=dev` to the make commands.

#### *tl;dr*
`make primary stage=dev`

`make backup stage=dev`

### CLI client
- run `make fs3_client`
- Then you will have the `fs3` executable available in the `cli-go` directory
  - use `./fs3 -h` to see how to use it; briefly:
    - `./fs3 cp <local_src> <remote_dst>`
    - `./fs3 rm <remote_path>`
    - `./fs3 get <remote_src> <local_dst>`
    - all params are paths

## Generating Protos
After acquiring the prerequisites, simply run `make protos` to regenerate.

### Prerequisites
#### Go
- decently recent protoc, we downloaded through `apt-get install protobuf-compiler` on ubuntu (wsl)
- `protoc-gen-go` and `protoc-gen-go-grpc`
  - see [grpc quickstart for go](https://grpc.io/docs/languages/go/quickstart/)

#### Python
- python 3.8 (untested with other versions, higher versions likely to work)
- install requirements.txt (`pip install -r requirements.txt`)
  - we recommend using a virtual environment
- reference: [grpc python basics](https://grpc.io/docs/languages/python/basics/)
