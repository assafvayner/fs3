.PHONY: FORCE

fs3_proto: protos/fs3/fs3.proto
	protoc \
		--proto_path=protos/ \
		--go_out=protos \
		--go_opt=paths=source_relative \
		--go-grpc_out=protos \
		--go-grpc_opt=paths=source_relative \
		./protos/fs3/fs3.proto
	python -m grpc_tools.protoc \
		-Iprotos \
		--python_out=protos \
		--pyi_out=protos \
		--grpc_python_out=protos \
		protos/fs3/fs3.proto

primarybackup_proto: protos/fs3/fs3.proto protos/primarybackup/primarybackup.proto
	protoc \
		--proto_path=protos/ \
		--proto_path=protos/fs3 \
		--proto_path=protos/primarybackup \
		--go_out=protos \
		--go_opt=paths=source_relative \
		--go-grpc_out=protos \
		--go-grpc_opt=paths=source_relative \
		./protos/primarybackup/primarybackup.proto

protos: fs3_proto primarybackup_proto

SERVER_FILES := $(shell find server -name "*.go")

server: $(SERVER_FILES)
	go build -o server/server server/server.go

primary: server
	./server/server primary $(stage)

backup: server
	./server/server backup $(stage)

CLIENT_GO_FILES := $(shell find cli-go -name "*.go")

fs3_client: $(CLIENT_GO_FILES)
	go build -o cli-go/fs3 cli-go/fs3.go

protos_clean: FORCE
	rm -f protos/*/*.go protos/*/*.py protos/*/*.pyi

# don't remove protos unless you are able to regenerate them
clean: FORCE
	rm -f server/server
	rm -f cli-go/fs3

# docker related tasks
FS3_NET="fs3-net"
PRIMARY_NAME="primary_container"
BACKUP_NAME="backup_container"

create_network: FORCE
	docker network create $(FS3_NET)

remove_network: FORCE
	docker network rm $(FS3_NET)

build_p: FORCE
	docker build -t fs3/primary -f Dockerfile_p .

run_p: FORCE
	docker run -d -p 127.0.0.1:5000:5000 --hostname primary.fs3 --network $(FS3_NET) --name $(PRIMARY_NAME) fs3/primary

build_b: FORCE
	docker build -t fs3/backup -f Dockerfile_b .

run_b: FORCE
	docker run -d -p 127.0.0.1:50000:50000 --hostname backup.fs3 --network $(FS3_NET) --name $(BACKUP_NAME) fs3/backup

shutdown_p: FORCE
	docker shutdown $(PRIMARY_NAME)
	docker rm $(PRIMARY_NAME)

shutdown_b: FORCE
	docker shutdown $(BACKUP_NAME)
	docker rm $(BACKUP_NAME)

up: FORCE
	docker compose -f fs3.yml up -d --build

down: FORCE
	docker compose -f fs3.yml down