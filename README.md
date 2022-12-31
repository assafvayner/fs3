# fs3

## Summary
A distributed remote file system with a grpc interface and command line tool to access it.
Jwt-dispensing authorization service that allows per user separation of access to files.
Http frontend server to utilize the system from a web context over cli/scripts.

Began as a UW CSE 453 (Data Center Systems) final project.

## Key generation notes
To run the app the internal worker nodes require tls to communicate securely, and to enable this we must generate some keys.
To do this, enter the `keys/grpc-tls` directory and run `./make_keys.sh` then all the requisite keys will be generated for tls within the containers using grpc.

To enable jwt usage, you will need a keys to use to sign or verify tokens.
You can generate such keys by entering the `keys/jwt` directory and running `go run keygen.go`.
Note that if you want to turn down the service then up again and be able to reuse old but non-expired jwt tokens then you will need to keep the same keys.

## Intended Runtime Configuration
The server nodes for this project are intended to be managed via docker swarm and ran on cloudlab servers.
There are make commands as well as instructions below to make running the application simply
While we do backup files we do not currently support a mechanism to make the backup node into the primary node and support allowing the backup to fail indefinitely.
The client code is also intended to be ran from cloudlab nodes so that they have access to communication without the public internet.

The client consists of 2 completely separate components, firstly a cli tool to make individual requests as proof of concept of the initial intended usage.
Secondly a python program used to make grpc calls and measure their duration as a method of analyzing performance.

We have used BloomRPC to test grpc APIs as well as Postman to test http apis

## Instructions to run
### Server in development
Tested on wsl/ubuntu mounted and mac, run the following commands (you will need docker installed)
- `sudo make up` - build the application and run it
  - if you wish to use the authservice, you will have to run `go run keygen.go` in the keys directory first
- `sudo make up_no_build` will pull prebuilt images from dockerhub so you don't need to build the application.

You can run the clients locally, compile the cli via `make fs3_client`

### Server on cloudlab
On cloudlab, first run `./start_docker.sh` in `/local/repository`
- then you can run `sudo make stack_up` which is an alias for `sudo docker stack deply -c fs3.yml fs3` to utilize the swarm
- you can also run on vms by running our `start_vms.sh` and then following the instructions from [CSE 453 lab 0](https://gitlab.cs.washington.edu/syslab/cse453-cloud-project/-/blob/main/docs/lab0.md#run-deathstarbench-in-vms-and-test-it) except change the last greyed out command to:
  - `sudo lxc exec vm1 -- docker stack deploy --compose-file fs3.yml fs3`

### Testing on cloudlab client
- requires bringing profile.py into repo root from 453scripts dir
- then on cloudlab creating a profile from repo
- option 1 running the python testing scripts:
  - we require some dependencies that are missing under the current config, run the following:
    - `wget https://bootstrap.pypa.io/get-pip.py`
    - `python3 get-pip.py`
    - `python3 -m pip install -r requirements.txt`
- option 2 using cli
  - first install go, run the following commands:
    - `wget https://go.dev/dl/go1.19.4.linux-amd64.tar.gz`
    - `sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz`
  - then run `make fs3_client`
  - the cli will be the resulting executable in `cli-go/fs3`
  - use `./fs3 -h` to get the options you can use

### CLI client
- run `make fs3_client`
- Then you will have the `fs3` executable available in the `cli-go` directory
  - use `./fs3 -h` to see how to use it; briefly:
    - `./fs3 cp <local_src> <remote_dst>`
    - `./fs3 rm <remote_path>`
    - `./fs3 get <remote_src> <local_dst>`
    - all params are paths
  - There are other commands for using the authservice as well as describing resources
    - run `./fs3 -h` to see them

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

# TBD, in no particular order:
- Host the service on non-cloudlab platform.
- Get a signed certificate (letsencrypt) to set frontend service serving https.
- Build a gui frontend, likely web based
- Deploy with kubernetes rather than docker swarm
- Set up failover state transfer, or set up stronger replication with a paxos variant
