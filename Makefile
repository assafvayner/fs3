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

authservice_proto: protos/authservice/authservice.proto
	protoc \
		--proto_path=protos/ \
		--go_out=protos \
		--go_opt=paths=source_relative \
		--go-grpc_out=protos \
		--go-grpc_opt=paths=source_relative \
		./protos/authservice/authservice.proto
	python -m grpc_tools.protoc \
		-Iprotos \
		--python_out=protos \
		--pyi_out=protos \
		--grpc_python_out=protos \
		protos/authservice/authservice.proto

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

protos: fs3_proto primarybackup_proto authservice_proto

SHARED_SERVER_FILES := $(shell find server/shared -name "*.go")

SERVER_FILES := $(shell find server/app -name "*.go")

server: $(SERVER_FILES) $(SHARED_SERVER_FILES)
	go build -o server/app/server server/app/server.go

AUTH_SERVER_FILES := $(shell find server/auth -name "*.go")

authserver: $(AUTH_SERVER_FILES) $(SHARED_SERVER_FILES)
	go build -o server/auth/authserver server/auth/server.go

FRONTEND_SERVER_FILES := $(shell find server/frontend -name "*.go")

frontendserver: $(FRONTEND_SERVER_FILES) $(SHARED_SERVER_FILES)
	go build -o server/frontend/frontendserver server/frontend/main.go

CLIENT_GO_FILES := $(shell find cli-go -name "*.go")

fs3_client: $(CLIENT_GO_FILES)
	go build -o cli-go/fs3 cli-go/fs3.go

protos_clean: FORCE
	rm -f protos/*/*.go protos/*/*.py protos/*/*.pyi

# don't remove protos unless you are able to regenerate them
clean: FORCE
	rm -f server/app/server
	rm -f server/auth/authserver
	rm -f cli-go/fs3

# docker related tasks
FS3_NET="fs3-net"
PRIMARY_NAME="primary_container"
BACKUP_NAME="backup_container"

build_server_image: FORCE
	docker build -t assafvayner/fs3:server -f Dockerfile .

push_server_image: FORCE
	docker push assafvayner/fs3:server

up: FORCE
	docker compose -f fs3.yml up -d --build

down: FORCE
	docker compose -f fs3.yml down