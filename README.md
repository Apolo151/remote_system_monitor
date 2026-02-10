# Remote System Health Monitor

A real-time system monitoring application using gRPC server-side streaming.

## Features
- Real-time CPU and RAM monitoring
- gRPC server-side streaming
- Graceful shutdown support
- Configurable monitoring intervals

## Prerequisites
- Go 1.20+
- Protocol Buffers compiler (protoc)

## Installation

```bash
go mod download
make proto
make build
```

## Usage

### Run Server
```bash
# Default (port 50051, 5s interval)
make run-server

# Custom configuration
go run cmd/server/main.go -port 8080 -interval 10s
```

### Run Client
```bash
# Default (localhost:50051)
make run-client

# Custom server
go run cmd/client/main.go -server remote.example.com:50051
```

## Project Structure
```
cmd/         - Entry points
internal/    - Private application code
pkg/         - Generated protobuf code
proto/   - Protocol buffer definitions
```

## Testing
```bash
make test
```