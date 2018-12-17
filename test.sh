#!/bin/bash

which go || {
    echo "go command was not found in this shell"
    exit 1
}

go test -v -cover -timeout 30s