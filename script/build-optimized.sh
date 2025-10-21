#!/bin/bash
set -e

# EMS 优化构建脚本
# 使用 BuildKit 缓存优化构建性能

echo "🚀 开始优化构建 EMS 项目..."

# 设置环境变量
export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

# 创建缓存目录
mkdir -p /tmp/.buildx-cache

# 清理旧的缓存（可选，根据需要启用）
# echo "🧹 清理旧缓存..."
# docker buildx prune --all --force

# 构建镜像
echo "🔨 构建 Docker 镜像..."
docker-compose -f script/docker-compose.dev.yml --profile amd64 up --build

# 移动缓存（临时修复，参考 GitHub Actions 最佳实践）
echo "📦 优化缓存..."
if [ -d "/tmp/.buildx-cache-new" ]; then
    rm -rf /tmp/.buildx-cache
    mv /tmp/.buildx-cache-new /tmp/.buildx-cache
    echo "✅ 缓存已优化"
fi

echo "🎉 构建完成！"
echo "💡 提示：下次构建将使用缓存，速度会更快"
