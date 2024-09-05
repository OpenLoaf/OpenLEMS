#!/bin/bash



# 获取当前脚本所在目录
SOURCE_DIR="$(dirname "$(realpath "$0")")"
echo "SOURCE_DIR: $SOURCE_DIR"
cd $SOURCE_DIR || exit 1



if [ "$1" == "arm" ]; then
  echo "Building for ARM"
  export pluginPath=${SOURCE_DIR}/arm-build
#  export GOARCH=arm
#  export GOOS=linux
#  export GOARM=7
  # 先编译ems，再编译插件
  go build -ldflags "-s -w" -o ${pluginPath}/ems ${SOURCE_DIR}/application/main.go || exit 1
  echo "Ems build success! To ${pluginPath}/ems"

  export CGO_ENABLED=1
#  export CC=arm-linux-gnueabihf-gcc
#  export CXX=arm-linux-gnueabihf-g++
elif [ "$1" == "debug" ]; then
  echo "Building in debug mode"
  export IS_DEBUG=true
#else
  export pluginPath=${SOURCE_DIR}/application/manifest
else
  export pluginPath=${SOURCE_DIR}/application/manifest
fi

#export pluginPath=/app/

echo "Building plugins to ${pluginPath}"

# 查找包含 //go:generate 指令的所有 Go 文件
files=$(grep -rl '//go:generate' plug_drivers --include \*.go)

currentPath=$(pwd)

# 在每个文件所在的目录中执行 go generate 文件名
for file in $files; do
    cd "$currentPath/$(dirname $file)" || exit 1
    echo "Running go generate $file "
    go generate  $(basename ${file}) || exit 1
done
#echo "Finish!Save To ${pluginPath}"
