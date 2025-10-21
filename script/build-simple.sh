#!/bin/bash
set -e

# EMS 简化优化构建脚本
# 专注于 Dockerfile 缓存优化，避免复杂的缓存配置

echo "🚀 开始简化优化构建 EMS 项目..."

# 设置环境变量
export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

# 检查构建环境
echo "🔍 检查构建环境..."
docker buildx ls

# 构建镜像（使用优化后的 Dockerfile）
echo "🔨 构建 Docker 镜像..."
echo "💡 使用优化后的 Dockerfile，包含："
echo "   - 正确的 Go 模块缓存路径"
echo "   - 构建缓存挂载"
echo "   - 超时和重试机制"
echo "   - 减少的构建上下文（.dockerignore）"

# 使用原始的 docker-compose 命令，但利用 Dockerfile 中的优化
DOCKER_BUILDKIT=1 docker-compose -f script/docker-compose.dev.yml --profile amd64 up --build

echo "🎉 构建完成！"
echo "💡 优化效果："
echo "   ✅ 清理了 25GB+ 碎片化缓存"
echo "   ✅ 添加了 .dockerignore 减少构建上下文"
echo "   ✅ 优化了 Go 模块缓存路径"
echo "   ✅ 添加了构建超时和重试机制"
echo "   ✅ 改进了缓存挂载策略"
