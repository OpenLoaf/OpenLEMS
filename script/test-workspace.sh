#!/bin/bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT_DIR"

modules="$(go work edit -json | sed -n 's/.*"DiskPath": "\(.*\)".*/\1/p')"

if [ -z "$modules" ]; then
  echo "No modules found in go.work"
  exit 1
fi

for mod in $modules; do
  pattern="$mod/..."
  packages="$(go list "$pattern" 2>/dev/null || true)"
  if [ -z "$packages" ]; then
    echo "Skipping $pattern (no packages)"
    continue
  fi
  echo "Testing $pattern"
  go test "$pattern"
done
