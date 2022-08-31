#!/bin/sh

set -e

cd "$(dirname "$0")/.."

script/update.sh

echo "==> Compiling bootstrapper..."
go build -o bin/bootstrapper cmd/bootstrapper/main.go

echo "==> Bootstrapper is now ready to go!"
