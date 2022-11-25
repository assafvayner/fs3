.PHONY: FORCE

fs3_proto: protos/fs3/fs3.proto
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		./protos/fs3/fs3.proto

primarybackup_proto: protos/fs3/fs3.proto protos/primarybackup/primarybackup.proto
	protoc \
		--proto_path=protos/fs3 \
		--proto_path=protos/primarybackup \
		--go_out=protos/primarybackup \
		--go_opt=paths=source_relative \
		--go-grpc_out=protos/primarybackup \
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
	rm protos/*/*.go

# don't remove protos unless you are able to regenerate them
clean:
	rm server/server
	rm cli-go/fs3
