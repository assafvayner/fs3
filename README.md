# fs3

CSE 453 Data Center Systems final project
Assaf Vayner (assafv) + Tom Wu (zw237)

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
