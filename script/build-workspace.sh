#!/bin/bash
set -euo pipefail

source "$(cd "$(dirname "$0")" && pwd)/workspace-common.sh"

cd "$ROOT_DIR"

prepare_hexlib
modules="$(workspace_modules)"

for mod in $modules; do
  pattern="$mod/..."
  packages="$(go list "$pattern" 2>/dev/null || true)"
  if [ -z "$packages" ]; then
    echo "Skipping $pattern (no packages)"
    continue
  fi
  echo "Building $pattern"
  go build "$pattern"
done
