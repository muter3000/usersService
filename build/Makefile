SHELL=/usr/bin/zsh

.PHONY: protos

protos:
	 protoc -I ../internal/proto-files/service ../internal/proto-files/service/*.proto \
	 --go_out=../internal/proto-files/service --go-grpc_out=../internal/proto-files/service