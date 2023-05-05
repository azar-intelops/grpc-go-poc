#!/bin/bash

# The path to the protobuf file relative to the root of your project.
PROTO_FILE="api/proto/*.proto"

# The output directory for the generated protobuf code.
OUTPUT_DIR="gen/pb"

# Generate the protobuf code.
if [ -d "gen/pb" ]; then rm -rf gen/pb; fi && mkdir gen/pb && protoc --go_out="$OUTPUT_DIR" --go-grpc_out="$OUTPUT_DIR" "$PROTO_FILE"
