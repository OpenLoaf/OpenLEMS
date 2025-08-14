#!/bin/bash

# go-work-update.sh - 更新 Go 工作区所有模块依赖，可排除指定目录

# 配置需要排除的目录（支持通配符）
EXCLUDE_DIRS=(
  "*/excluded_dir"    # 排除所有名为 excluded_dir 的目录
  "testdata/*"        # 排除所有 testdata 目录下的内容
  "vendor"            # 排除 vendor 目录
)

# 生成 exclude 参数
exclude_args=()
for pattern in "${EXCLUDE_DIRS[@]}"; do
  exclude_args+=( "-e" "$pattern" )
done

# 获取工作区中所有模块目录（排除指定目录）
modules=()
while IFS= read -r dir; do
  modules+=("$dir")
done < <(go work edit -json | jq -r '.Use[].DiskPath' | grep -v "${exclude_args[@]}")

if [ ${#modules[@]} -eq 0 ]; then
  echo "没有找到需要更新的 Go 模块"
  exit 0
fi

echo "将更新以下 Go 模块:"
printf " - %s\n" "${modules[@]}"
echo

# 更新每个模块
for dir in "${modules[@]}"; do
  echo "正在更新模块: $dir"
  (
    cd "$dir" || exit 1
    go get -u ./...
    go mod tidy
  )
  echo
done

# 同步工作区
echo "正在同步工作区..."
go work sync

echo "所有模块更新完成!"