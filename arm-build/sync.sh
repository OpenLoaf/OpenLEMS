#!/bin/bash

# 检查参数数量
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <username> <server_address> <destination_directory>"
    exit 1
fi

# 从参数获取用户名、服务器地址和目标目录
USERNAME=$1
SERVER_ADDRESS=$2
DEST_DIR=$3

# 获取当前脚本所在目录
SOURCE_DIR="$(dirname "$(realpath "$0")")"
echo "SOURCE_DIR: $SOURCE_DIR"

# 使用rsync同步文件
rsync -avz --delete -P "${SOURCE_DIR}/" "${USERNAME}@${SERVER_ADDRESS}:${DEST_DIR}"

echo "Files synced to ${USERNAME}@${SERVER_ADDRESS}:${DEST_DIR}"
