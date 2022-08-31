#!/bin/sh

set -e

cd "$(dirname "$0")/.."

script/setup.sh

echo "==> Running bootstrapper..."
./bin/bootstrapper
