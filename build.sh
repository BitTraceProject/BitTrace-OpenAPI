#!/bin/bash

export GOOS="linux" # windows darwin linux

go build -v -o ./output/openapi ./cmd...
echo "build successfully!"
