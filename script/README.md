# LEMS Docker 多架构构建环境

支持 **AMD64**、**ARM64**、**ARM32** 三种架构的多阶段 Docker 构建，适用于服务器、树莓派等多种硬件平台。

## 📋 文件说明

- `Dockerfile.dev` - 多架构多阶段构建的 Docker 镜像配置
- `docker-compose.dev.yml` - 多架构 Docker Compose 配置
- `README.md` - 使用说明文档

## 🏗️ 架构特点

### 支持的平台

| 架构 | 平台标识 | 适用设备 | 镜像名称 | 容器名称 | 端口 |
|------|---------|---------|---------|----------|------|
| AMD64 | linux/amd64 | x86 服务器、PC | lems:amd64 | lems-amd64 | 8080 |
| ARM64 | linux/arm64 | 树莓派 4/5、Apple Silicon | lems:arm64 | lems-arm64 | 8081 |
| ARM32 | linux/arm/v7 | 树莓派 3、嵌入式设备 | lems:arm32 | lems-arm32 | 8082 |

### 多阶段构建流程

```
第一阶段 (builder)
  ├─ 安装 C++ 和 Go 构建依赖
  ├─ 下载对应架构的 Go 二进制 (1.25.3)
  ├─ 编译 C++ hexlib 库
  ├─ 下载 Go 模块依赖
  └─ 编译 EMS 应用程序

第二阶段 (runtime)
  ├─ 使用 Alpine Linux 轻量级镜像
  ├─ 复制编译好的二进制和库文件
  ├─ 设置环境变量 (LD_LIBRARY_PATH)
  └─ 配置健康检查
```

### 优化特性

- ✅ **BuildKit 缓存**: 使用 `--mount=type=cache` 加速依赖下载
- ✅ **动态架构适配**: 根据 `TARGETARCH` 自动选择对应的 Go 版本
- ✅ **YAML 锚点复用**: 避免配置重复,易于维护
- ✅ **独立数据卷**: 每个架构独立的输出目录
- ✅ **固化环境变量**: LD_LIBRARY_PATH 在 Dockerfile 中固定

## 🚀 快速开始

### 1. 构建单个架构

```bash
# 构建 AMD64 版本
docker-compose -f script/docker-compose.dev.yml --profile amd64 build

# 构建 ARM64 版本
docker-compose -f script/docker-compose.dev.yml --profile arm64 build

# 构建 ARM32 版本
docker-compose -f script/docker-compose.dev.yml --profile arm32 build
```

### 2. 构建所有架构

```bash
# 使用 'all' profile 构建全部架构
docker-compose -f script/docker-compose.dev.yml --profile all build
```

### 3. 启动服务

```bash
# 启动 AMD64 版本 (端口 8080)
docker-compose -f script/docker-compose.dev.yml --profile amd64 up -d

# 启动 ARM64 版本 (端口 8081)
docker-compose -f script/docker-compose.dev.yml --profile arm64 up -d

# 启动 ARM32 版本 (端口 8082)
docker-compose -f script/docker-compose.dev.yml --profile arm32 up -d

# 启动所有版本
docker-compose -f script/docker-compose.dev.yml --profile all up -d
```

### 4. 访问应用

根据启动的架构,访问对应端口:

- **AMD64**: http://localhost:8080
- **ARM64**: http://localhost:8081
- **ARM32**: http://localhost:8082

健康检查: `http://localhost:808x/health`

### 5. 停止服务

```bash
# 停止特定架构
docker-compose -f script/docker-compose.dev.yml --profile amd64 down

# 停止所有架构
docker-compose -f script/docker-compose.dev.yml --profile all down
```

## 🔧 高级用法

### 使用 Docker Buildx 构建多架构镜像

```bash
# 一次性构建所有架构 (需要 buildx 支持)
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -f script/Dockerfile.dev \
  -t lems:latest \
  .

# 构建并推送到 Registry
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -f script/Dockerfile.dev \
  -t registry.example.com/lems:latest \
  --push \
  .
```

### 清理旧的 PID 文件

如果遇到 PID 冲突问题,清理本地 PID 文件:

```bash
# 清理所有架构的 PID 文件
rm -f out/amd64/ems.pid out/arm64/ems.pid out/arm32/ems.pid
```

### 查看日志

```bash
# 查看 AMD64 容器日志
docker logs lems-amd64 -f

# 查看最近 100 行日志
docker logs lems-amd64 --tail 100
```

### 进入容器调试

```bash
# 进入 AMD64 容器
docker exec -it lems-amd64 sh

# 在容器内运行命令
docker exec lems-amd64 ./ems --help
```

## 📦 数据持久化

### 卷配置说明

每个架构使用独立的绑定挂载卷:

```yaml
volumes:
  vol-amd64:
    driver: local
    driver_opts:
      type: none
      o: bind
      device: ../out/amd64  # 挂载到 out/amd64/
```

### 输出目录结构

```
out/
├── amd64/          # AMD64 架构输出
│   ├── logs/       # 日志文件
│   ├── data/       # 数据库文件
│   ├── static/     # 静态文件
│   ├── ptdb/       # 时序数据库
│   └── ems.pid     # 进程 PID 文件
├── arm64/          # ARM64 架构输出
└── arm32/          # ARM32 架构输出
```

## ⚙️ 配置说明

### Dockerfile 关键配置

| 配置项 | 说明 | 值 |
|-------|------|-----|
| TARGETPLATFORM | 目标平台 | 动态 (linux/amd64, linux/arm64, linux/arm/v7) |
| TARGETARCH | 目标架构 | 动态 (amd64, arm64, arm) |
| Go 版本 | Go 语言版本 | 1.25.3 |
| C++ 编译器 | GCC 版本 | latest (Debian) |
| 运行时基础镜像 | Alpine 版本 | latest |
| LD_LIBRARY_PATH | C++ 库路径 | /app/lib (固化) |

### 启动命令

容器启动时自动执行:

```bash
sh -c 'rm -f ./out/ems.pid && ./ems --web=true --profile=prod'
```

说明:
1. `rm -f ./out/ems.pid` - 清理旧的 PID 文件,避免冲突
2. `--web=true` - 启用 Web 服务
3. `--profile=prod` - 使用生产环境配置

### 健康检查

- **检查方式**: `wget` 访问 `/health` 接口
- **检查间隔**: 30 秒
- **超时时间**: 10 秒
- **重试次数**: 3 次
- **启动等待**: 40 秒

## 🐛 常见问题

### 1. 构建失败: "gcc: error: unrecognized command-line option '-m64'"

**原因**: 使用了错误的构建平台 (BUILDPLATFORM vs TARGETPLATFORM)

**解决**: 已修复,Dockerfile 使用 `--platform=$TARGETPLATFORM` 确保编译器架构正确

### 2. 容器不断重启: PID 文件冲突

**原因**: 旧的 PID 文件导致程序认为已有进程在运行

**解决**:
```bash
# 清理 PID 文件
rm -f out/amd64/ems.pid

# 重启容器
docker-compose -f script/docker-compose.dev.yml --profile amd64 restart
```

### 3. 端口冲突

**原因**: 端口 8080/8081/8082 已被占用

**解决**: 修改 `docker-compose.dev.yml` 中的端口映射
```yaml
ports:
  - "9080:80"  # 修改为其他端口
```

### 4. Go 模块下载慢

**原因**: 网络问题或代理配置

**解决**: Dockerfile 已配置国内代理 `GOPROXY=https://goproxy.cn,direct`

### 5. C++ 库编译失败

**检查**:
```bash
# 查看 C++ 源代码
ls -la @cpp/hexlib/

# 检查 Makefile
cat @cpp/hexlib/Makefile
```

### 6. 清理构建缓存

```bash
# 清理所有未使用的镜像和容器
docker system prune -a

# 清理特定镜像
docker rmi lems:amd64 lems:arm64 lems:arm32

# 清理卷
docker volume prune
```

## 📊 性能优化

### 构建优化

1. **BuildKit 缓存**: 使用 `--mount=type=cache` 缓存 Go 模块和构建缓存
2. **分层复制**: 先复制依赖文件,再复制源代码,最大化缓存利用
3. **并行下载**: Go 模块并行下载

### 镜像大小优化

| 阶段 | 基础镜像 | 大小 | 说明 |
|------|---------|------|------|
| builder | gcc:latest | ~1.5GB | 仅构建阶段使用 |
| runtime | alpine:latest | ~5MB | 最小化运行时镜像 |
| 最终镜像 | - | ~50MB | 仅包含二进制和必要库 |

## 🔐 安全建议

1. **不使用 root 用户**: 生产环境建议添加非 root 用户运行
2. **只读文件系统**: 考虑使用 `--read-only` 挂载
3. **限制资源**: 使用 Docker 资源限制 (CPU/内存)
4. **定期更新**: 定期更新基础镜像和依赖

## 📚 技术栈

- **Go**: 1.25.3
- **C++**: GCC (Debian latest)
- **运行时**: Alpine Linux latest
- **框架**: GoFrame v2
- **构建工具**: Docker BuildKit
- **CGO**: 支持 C++ 库集成

## 🆚 版本差异

本配置与之前版本的主要改进:

| 改进项 | 之前 | 现在 |
|--------|------|------|
| 架构支持 | 仅 AMD64 | AMD64 + ARM64 + ARM32 |
| 镜像命名 | ems-plan:dev | lems:amd64/arm64/arm32 |
| Go 模块缓存 | 使用 Docker 卷 | 使用 BuildKit 缓存 |
| LD_LIBRARY_PATH | docker-compose 配置 | Dockerfile 固化 |
| PID 文件处理 | 手动清理 | 启动脚本自动清理 |
| 配置复用 | 重复定义 | YAML 锚点复用 |

## 📖 相关文档

- [项目主 README](../README.md) - 完整项目说明
- [CLAUDE.md](../CLAUDE.md) - 开发指南
- [CGO 集成规范](.cursor/rules/tech-cgo-integration_v1.0.mdc) - CGO 开发规范

## 🤝 技术支持

如果遇到问题,请检查:

1. ✅ Docker 版本 >= 20.10 (支持 BuildKit)
2. ✅ Docker Compose 版本 >= 2.0
3. ✅ 端口是否被占用 (8080/8081/8082)
4. ✅ 磁盘空间是否充足 (至少 5GB)
5. ✅ 网络连接是否正常

更多问题请在项目 Issues 中反馈。
