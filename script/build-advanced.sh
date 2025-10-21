#!/bin/bash
set -e

# EMS 高级优化构建脚本
# 包含网络优化、重试机制和详细日志

echo "🚀 开始高级优化构建 EMS 项目..."

# 设置环境变量
export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1
export BUILDKIT_PROGRESS=plain

# 创建缓存目录
mkdir -p /tmp/.buildx-cache

# 检查 Docker 和 BuildKit 状态
echo "🔍 检查构建环境..."
docker buildx ls
docker system df

# 设置 Go 代理优化（在 Dockerfile 中已设置，这里作为备用）
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

# 构建函数，包含重试机制
build_with_retry() {
    local max_attempts=3
    local attempt=1
    
    while [ $attempt -le $max_attempts ]; do
        echo "🔄 构建尝试 $attempt/$max_attempts..."
        
        if docker-compose -f script/docker-compose.dev.yml --profile amd64 up --build; then
            echo "✅ 构建成功！"
            return 0
        else
            echo "❌ 构建失败，尝试 $attempt/$max_attempts"
            if [ $attempt -lt $max_attempts ]; then
                echo "⏳ 等待 10 秒后重试..."
                sleep 10
                # 清理失败的构建缓存
                docker builder prune -f
            fi
            attempt=$((attempt + 1))
        fi
    done
    
    echo "💥 构建失败，已尝试 $max_attempts 次"
    return 1
}

# 执行构建
if build_with_retry; then
    # 移动缓存（临时修复，参考 GitHub Actions 最佳实践）
    echo "📦 优化缓存..."
    if [ -d "/tmp/.buildx-cache-new" ]; then
        rm -rf /tmp/.buildx-cache
        mv /tmp/.buildx-cache-new /tmp/.buildx-cache
        echo "✅ 缓存已优化"
    fi
    
    echo "🎉 构建完成！"
    echo "💡 提示：下次构建将使用缓存，速度会更快"
    
    # 显示构建统计
    echo "📊 构建统计："
    docker system df
    docker buildx du | head -20
else
    echo "💥 构建失败，请检查错误信息"
    exit 1
fi
