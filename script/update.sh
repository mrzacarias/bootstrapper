#!/bin/sh

set -e

cd "$(dirname "$0")/.."

echo "==> Ensuring packages..."
go mod tidy -v

echo "==> Generating compiled templates package..."
go install github.com/go-bindata/go-bindata/...
go-bindata -o templates/templates.go -pkg templates -ignore templates.go templates/...
