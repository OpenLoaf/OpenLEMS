#!/bin/bash

# EMS 项目 ARMv7l Linux 交叉编译脚本
# 构建目标: Linux ARMv7l 平台
# 输出目录: application/out/v1.0.0/
# 编译标签: dev

set -e  # 遇到错误立即退出

echo "开始交叉编译 EMS 项目为 Linux ARMv7l 平台..."

# 设置交叉编译环境变量
export GOOS=linux
export GOARCH=arm
export GOARM=7

# 创建输出目录
mkdir -p application/out/v1.0.0

# 交叉编译命令
echo "编译参数:"
echo "  GOOS=$GOOS"
echo "  GOARCH=$GOARCH"
echo "  GOARM=$GOARM"
echo "  标签: dev"
echo "  输出: application/out/v1.0.0/ems-linux-armv7l"
echo ""

# 执行交叉编译
go build -tags=dev -o application/out/v1.0.0/ems-linux-armv7l ./application

# 检查编译结果
if [ -f "application/out/v1.0.0/ems-linux-armv7l" ]; then
    echo "✅ 编译成功!"
    echo "📁 输出文件: application/out/v1.0.0/ems-linux-armv7l"
    echo "📊 文件大小: $(ls -lh application/out/v1.0.0/ems-linux-armv7l | awk '{print $5}')"
    echo "🔍 架构信息: $(file application/out/v1.0.0/ems-linux-armv7l | cut -d: -f2-)"
else
    echo "❌ 编译失败!"
    exit 1
fi

echo "编译完成！"
