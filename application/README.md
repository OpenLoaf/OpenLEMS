# EMS Plan Application

## GoFrame API 生成方法总结

本项目基于 GoFrame v2.9.0 框架开发，以下是 GoFrame 中生成 API 的完整方法总结。

### 1. API 生成命令概览

GoFrame 提供了强大的代码生成工具，主要包含以下命令：

```bash
# 生成控制器接口规范
gf gen ctrl

# 生成数据访问层
gf gen dao

# 生成服务层接口
gf gen service

# 生成枚举维护
gf gen enums

# 生成协议缓冲区文件
gf gen pb

# 生成数据表PB
gf gen pbentity
```

### 2. API 接口规范生成 (gen ctrl)

#### 2.1 基本用法

```bash
# 生成控制器接口规范
gf gen ctrl

# 指定源目录和目标目录
gf gen ctrl -s internal/controller -d api
```

#### 2.2 生成的文件结构

生成的 API 文件遵循以下结构：

```
api/
├── telemetry/
│   ├── telemetry.go          # 接口定义
│   └── v1/
│       ├── telemetry_get.go  # 请求/响应结构体
│       └── telemetry_description.go
```

#### 2.3 接口定义示例

```go
// api/telemetry/telemetry.go
package telemetry

import (
    "context"
    "application/api/telemetry/v1"
)

type ITelemetryV1 interface {
    GetTelemetryDescription(ctx context.Context, req *v1.GetTelemetryDescriptionReq) (res *v1.GetTelemetryDescriptionRes, err error)
    GetTelemetryGet(ctx context.Context, req *v1.GetTelemetryGetReq) (res *v1.GetTelemetryGetRes, err error)
}
```

#### 2.4 请求/响应结构体示例

```go
// api/telemetry/v1/telemetry_get.go
package v1

import (
    "github.com/gogf/gf/v2/frame/g"
)

type GetTelemetryGetReq struct {
    g.Meta       `path:"/telemetry/get" method:"get" tags:"遥测相关" summary:"获取某个遥测设备的当前值（用于测试）"`
    DeviceId     string `json:"deviceId,omitempty" v:"required|length:1,32#请输入设备Key|设备Key长度为:min到:max位" dc:"设备Key"`
    TelemetryKey string `json:"telemetryKey,omitempty" v:"required|length:1,32#请输入遥测Key|遥测Key长度为:min到:max位" dc:"遥测Key"`
}

type GetTelemetryGetRes struct {
    TestJoinKey      string `json:"testJoinKey" dc:"测试联合Key"`
    DeviceId         string `json:"deviceId" dc:"设备Key"`
    TelemetryKey     string `json:"telemetryKey" dc:"遥测Key"`
    TelemetryKeyName string `json:"telemetryKeyName" dc:"遥测名称"`
    Value            any    `json:"value" dc:"遥测值"`
}
```

### 3. 服务层生成 (gen service)

#### 3.1 基本用法

```bash
# 生成服务层接口
gf gen service

# 指定源目录和目标目录
gf gen service -s internal/logic -d internal/service
```

#### 3.2 设计背景

- 提供一种代码管理方式，通过具体模块实现直接生成模块接口定义
- 简化业务逻辑实现与接口分离的实现
- 降低模块方法与接口定义的重复操作

#### 3.3 使用步骤

1. **引入配置文件**
   - 使用提供的 `watchers.xml` 配置文件（Goland IDE）
   - 或配置 VS Code 的 RunOnSave 插件

2. **编写业务逻辑代码**
   ```go
   // internal/logic/user/user.go
   package user
   
   type sUser struct{}
   
   func init() {
       Service = New()
   }
   
   func New() *sUser {
       return &sUser{}
   }
   
   func (s *sUser) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
       // 业务逻辑实现
       return
   }
   ```

3. **生成接口及服务注册文件**
   ```bash
   gf gen service
   ```

4. **服务实现注入**
   ```go
   // internal/logic/user/user.go
   var localUser IUser
   
   func User() IUser {
       if localUser == nil {
           panic("implement not found for interface IUser, forgot register?")
       }
       return localUser
   }
   
   func RegisterUser(i IUser) {
       localUser = i
   }
   ```

5. **启动文件中引用接口实现注册**
   ```go
   // main.go
   import (
       _ "application/internal/logic"
       _ "application/internal/service"
   )
   ```

### 4. 控制器实现

#### 4.1 控制器结构

```go
// internal/controller/telemetry/telemetry_new.go
package telemetry

import (
    "application/api/telemetry"
)

type ControllerV1 struct{}

func NewV1() telemetry.ITelemetryV1 {
    return &ControllerV1{}
}
```

#### 4.2 方法实现

```go
// internal/controller/telemetry/telemetry_v1_get_telemetry_get.go
package telemetry

import (
    "application/api/telemetry/v1"
    "context"
)

func (c *ControllerV1) GetTelemetryGet(ctx context.Context, req *v1.GetTelemetryGetReq) (res *v1.GetTelemetryGetRes, err error) {
    // 业务逻辑实现
    return &v1.GetTelemetryGetRes{
        TestJoinKey:      req.DeviceId + ":" + req.TelemetryKey,
        DeviceId:         req.DeviceId,
        TelemetryKey:     req.TelemetryKey,
        TelemetryKeyName: "遥测名称",
        Value:            "遥测值",
    }, nil
}
```

### 5. 路由注册

#### 5.1 服务注册

```go
// internal/cmd/cmd_web.go
func startWeb(ctx context.Context) *ghttp.Server {
    s := g.Server()
    
    s.Group("/api", func(group *ghttp.RouterGroup) {
        group.Middleware(
            ghttp.MiddlewareCORS,
            // 其他中间件...
        )
        
        // 注册控制器
        group.Bind(telemetry.NewV1())
        group.Bind(device.NewV1())
        group.Bind(driver.NewV1())
        // 更多控制器...
    })
    
    return s
}
```

### 6. 自动模式配置

#### 6.1 Goland IDE 配置

使用提供的 `watchers.xml` 配置文件，在文件修改时自动生成接口文件。

#### 6.2 VS Code 配置

```json
{
    "emeraldwalk.runonsave": {
        "commands": [
            {
                "match": ".*logic.*go",
                "isAsync": true,
                "cmd": "gf gen service"
            }
        ]
    }
}
```

### 7. 注意事项

#### 7.1 结构体命名规范

- 控制器结构体应以 `Controller` 开头
- 服务结构体应以 `s` 开头，如 `sUser`
- 接口结构体应以 `I` 开头，如 `IUser`

#### 7.2 文件组织

- API 定义放在 `api/` 目录
- 控制器实现放在 `internal/controller/` 目录
- 业务逻辑放在 `internal/logic/` 目录
- 服务接口放在 `internal/service/` 目录

#### 7.3 版本管理

- 使用 `v1/`, `v2/` 等目录管理 API 版本
- 每个版本包含独立的请求/响应结构体

### 8. 常用命令

```bash
# 生成所有代码
gf gen all

# 运行项目（热编译）
gf run

# 构建项目
gf build

# 交叉编译
gf build -a "linux,amd64" -s "application"

# 打包资源
gf pack
```

### 9. 参考文档

- [GoFrame 官方文档](https://goframe.org/)
- [API 生成文档](https://goframe.org/docs/cli/gen-service)
- [接口规范生成](https://goframe.org/quick/scaffold-api-definition)

---

本项目展示了完整的 GoFrame API 开发流程，从接口定义到控制器实现，再到服务注册，形成了一个完整的微服务架构。
