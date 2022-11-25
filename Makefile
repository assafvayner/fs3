.PHONY: FORCE

protoc_single=protoc --go_out=protos/$(1) --go_opt=paths=source_relative \
	--go-grpc_out=protos/$(1) --go-grpc_opt=paths=source_relative $(2) $(3) \
	./protos/$(1)/$(1).proto

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

protos_clean: FORCE
	rm -f protos/*/*.go protos/*/*.py protos/*/*.pyi

clean:
	make protos_clean
	rm server/server
