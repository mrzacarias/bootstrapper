#!/bin/sh

set -e

cd "$(dirname "$0")/.."

[ -z "$DEBUG" ] || set -x

script/update.sh

echo "==> Running tests…"
if [ -n "$1" ]; then
  testdir="/$*"
fi

docker-compose run --rm app go test -cover -race .$testdir/...
