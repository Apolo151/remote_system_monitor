#!/bin/bash

# Generate proto files
# This script compiles .proto files to generate Go code

set -e

PROTO_DIR="${PROTO_DIR:-.}"
OUT_DIR="${OUT_DIR:-.}"

if [ ! -d "$PROTO_DIR" ]; then
    echo "Error: Proto directory not found: $PROTO_DIR"
    exit 1
fi

# Find and compile all .proto files
for proto_file in "$PROTO_DIR"/*.proto; do
    if [ -f "$proto_file" ]; then
        echo "Generating from: $proto_file"
        protoc --proto_path="$PROTO_DIR" \
       --go_out="$OUT_DIR" \
       --go_opt=paths=source_relative \
       --go-grpc_out="$OUT_DIR" \
       --go-grpc_opt=paths=source_relative \
       "$proto_file"
    fi
done

echo "Proto generation complete"