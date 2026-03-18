#!/bin/bash

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

workspace_modules() {
  local modules

  modules="$(cd "$ROOT_DIR" && go work edit -json | sed -n 's/.*"DiskPath": "\(.*\)".*/\1/p')"

  if [ -z "$modules" ]; then
    echo "No modules found in go.work" >&2
    return 1
  fi

  printf '%s\n' "$modules"
}

prepare_hexlib() {
  local hexlib_dir="$ROOT_DIR/@cpp/hexlib"

  if [ ! -f "$hexlib_dir/Makefile" ]; then
    return 0
  fi

  if ! command -v make >/dev/null 2>&1; then
    echo "make is required to build @cpp/hexlib" >&2
    return 1
  fi

  echo "Building C++ support library (@cpp/hexlib)"
  make -C "$hexlib_dir" build
}
