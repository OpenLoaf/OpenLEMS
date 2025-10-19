#!/bin/bash
set -e

echo "构建 C++ 库..."
make build

echo "运行 MPC 卡尔曼滤波器测试..."
export CGO_ENABLED=1
export CGO_LDFLAGS="-L$(pwd)/build -lhexlib -Wl,-rpath,$(pwd)/build"

cd hexlib && go test -v
