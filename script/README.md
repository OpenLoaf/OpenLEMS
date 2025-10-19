# EMS Plan Docker 开发环境

基于最佳实践的多阶段Docker构建，支持C++编译和Go应用开发。

## 文件说明

- `Dockerfile.dev` - 多阶段构建的Docker镜像配置
- `docker-compose.dev.yml` - 开发环境Docker Compose配置
- `README.md` - 使用说明文档

## 架构特点

### 多阶段构建

1. **C++ 编译阶段**: 使用 `gcc:latest` 编译 C++ 库
2. **Go 构建阶段**: 使用 `golang:1.25-alpine` 编译 Go 应用
3. **运行时阶段**: 使用 `alpine:latest` 轻量级运行时

### 优化特性

- **依赖缓存**: 利用 Docker 层缓存加速构建
- **模块化复制**: 分别复制 go.mod 文件利用缓存
- **CGO 支持**: 支持 C++ 库链接
- **健康检查**: 自动监控服务状态

## 快速开始

### 1. 启动开发环境

```bash
# 进入脚本目录
cd script

# 启动开发环境
docker-compose -f docker-compose.dev.yml up --build

# 后台运行
docker-compose -f docker-compose.dev.yml up -d --build
```

### 2. 访问应用

- **Web 管理界面**: http://localhost:80
- **健康检查**: http://localhost:80/health
- **API 文档**: http://localhost:80/api.json

### 3. 停止服务

```bash
# 停止服务
docker-compose -f docker-compose.dev.yml down

# 停止并删除数据卷
docker-compose -f docker-compose.dev.yml down -v
```

## 开发工作流

### 代码热重载

由于挂载了本地代码目录，修改代码后需要重启容器：

```bash
# 重启服务
docker-compose -f docker-compose.dev.yml restart

# 或者重新构建并启动
docker-compose -f docker-compose.dev.yml up --build
```

### 查看日志

```bash
# 查看实时日志
docker-compose -f docker-compose.dev.yml logs -f

# 查看特定服务日志
docker-compose -f docker-compose.dev.yml logs -f ems-dev
```

### 进入容器

```bash
# 进入运行中的容器
docker-compose -f docker-compose.dev.yml exec ems-dev bash

# 在容器内运行命令
docker-compose -f docker-compose.dev.yml exec ems-dev ./ems --help
```

## 性能优化

### 构建优化

- **分层缓存**: 先复制 go.mod 文件，再复制源代码
- **并行构建**: 利用多核 CPU 加速编译
- **精简镜像**: 运行时镜像仅包含必要依赖

### 运行时优化

- **静态链接**: 减少运行时依赖
- **Alpine 基础**: 最小化镜像大小
- **健康检查**: 自动故障恢复

## 配置说明

### 端口映射

- **80:80** - Web 服务端口（符合生产配置）

### 数据挂载

- **源代码目录** - 支持代码热重载
- **C++ 源代码** - `@cpp/` 目录
- **配置文件** - `manifest/config/` 目录
- **数据目录** - `out/` 目录（日志、数据库、PID 文件）
- **Go 模块缓存** - 加速依赖下载

### 环境变量

- `GO111MODULE=on` - 启用 Go 模块
- `GOPROXY=https://goproxy.cn,direct` - 使用国内代理
- `GOSUMDB=sum.golang.google.cn` - 使用国内校验数据库
- `CGO_ENABLED=1` - 启用 CGO
- `LD_LIBRARY_PATH=/app/lib` - C++ 库路径

### 启动参数

- `--web=true` - 启用 Web 服务
- `--profile=prod` - 使用生产配置

### 健康检查

- **检查间隔**: 30秒
- **超时时间**: 10秒
- **重试次数**: 3次
- **启动等待**: 40秒

## 常见问题

### 1. 构建失败

```bash
# 清理构建缓存
docker system prune -a

# 重新构建
docker-compose -f docker-compose.dev.yml build --no-cache
```

### 2. 端口冲突

修改 `docker-compose.dev.yml` 中的端口映射：

```yaml
ports:
  - "8080:80"  # 改为 8080:80
```

### 3. C++ 编译错误

确保 `@cpp` 目录包含完整的 C++ 源代码和 Makefile：

```bash
# 检查 C++ 代码
ls -la @cpp/
cat @cpp/Makefile
```

### 4. Go 模块问题

```bash
# 清理 Go 模块缓存
docker volume rm script_go-mod-cache

# 重新下载依赖
docker-compose -f docker-compose.dev.yml up --build
```

### 5. 依赖下载失败

如果 Go 模块下载失败，可以尝试：

```bash
# 清理 Go 模块缓存
docker-compose -f docker-compose.dev.yml exec ems-dev go clean -modcache

# 重新下载依赖
docker-compose -f docker-compose.dev.yml exec ems-dev go mod download
```

### 6. 数据库文件权限

如果 SQLite 数据库文件权限有问题：

```bash
# 检查数据库文件权限
ls -la ../out/data

# 调整权限
chmod 664 ../out/data
```

## 开发建议

### 1. 使用 IDE 开发

推荐使用支持 Docker 的 IDE（如 VS Code + Docker 扩展）进行开发，可以：
- 直接在容器内调试
- 实时查看日志
- 快速重启服务

### 2. 代码同步

由于使用了代码挂载，本地修改会立即反映到容器中，但需要重启容器才能生效。

### 3. 数据备份

定期备份 `out/` 目录中的数据：

```bash
# 备份数据目录
tar -czf ems-backup-$(date +%Y%m%d).tar.gz ../out/

# 恢复数据
tar -xzf ems-backup-20240101.tar.gz
```

## 技术栈

- **Go**: 1.25 (Alpine)
- **C++**: GCC (Ubuntu)
- **GoFrame**: v2
- **Docker**: 多阶段构建
- **Alpine Linux**: 轻量级运行时
- **CGO**: C++ 库集成

## 生产部署

本配置仅适用于开发环境。生产环境部署请参考项目根目录的 `Dockerfile` 和相关的生产部署文档。

## 技术支持

如果遇到问题，请检查：

1. Docker 和 Docker Compose 版本
2. 端口是否被占用
3. 文件权限是否正确
4. 网络连接是否正常
5. C++ 编译环境是否完整

更多信息请参考项目主 README.md 文件。