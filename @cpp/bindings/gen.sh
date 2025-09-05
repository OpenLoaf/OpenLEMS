#!/usr/bin/env bash
set -euo pipefail

# Generate Go bindings using c-for-go

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CPP_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

if ! command -v c-for-go >/dev/null 2>&1; then
  echo "c-for-go not found. Install: go install github.com/xlab/c-for-go@latest" >&2
  exit 1
fi

pushd "${SCRIPT_DIR}" >/dev/null
c-for-go -out .. cforgo.yml
popd >/dev/null

echo "Bindings generated under @cpp/hexlib"


