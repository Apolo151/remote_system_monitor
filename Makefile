.PHONY: proto build run-server run-client test clean

proto:
	PROTO_DIR=proto	OUT_DIR=pkg/monitorpb	./scripts/generate_proto.sh

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/main.go

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/