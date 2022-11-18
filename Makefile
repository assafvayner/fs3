.PHONY: FORCE

protoc_single=protoc --go_out=protos/$(1) --go_opt=paths=source_relative \
	--go-grpc_out=protos/$(1) --go-grpc_opt=paths=source_relative $(2) $(3) \
	./protos/$(1)/$(1).proto

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

protos_clean: FORCE
	rm protos/*/*.go