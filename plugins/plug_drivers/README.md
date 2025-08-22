# EMS Plan 驱动构建工具

这是一个用于构建EMS Plan项目中所有驱动插件的Makefile工具。

## 功能特性

- ✅ **单个驱动构建**: 支持构建指定的单个驱动
- ✅ **批量驱动构建**: 支持一次性构建所有驱动
- ✅ **自定义输出目录**: 支持指定驱动文件保存路径
- ✅ **彩色日志输出**: 提供清晰的颜色编码日志信息
- ✅ **驱动状态检查**: 查看所有驱动的构建状态
- ✅ **清理功能**: 支持清理单个或所有驱动
- ✅ **错误处理**: 完善的错误检查和提示
- ✅ **版本管理**: 自动获取驱动版本信息
- ✅ **文件信息显示**: 构建成功后显示文件路径和大小（单个和批量构建都支持）
- ✅ **失败原因显示**: 构建失败后显示详细错误信息

## 使用方法

### 查看帮助信息
```bash
make help
```

### 查看可用驱动列表
```bash
make list
```

### 查看驱动构建状态
```bash
make status
```

### 构建单个驱动
```bash
make project=<驱动名称>
# 例如：
make project=gpio_basic

# 指定自定义输出目录
make project=<驱动名称> outdir=<路径>
# 例如：
make project=gpio_basic outdir=./custom_output
```

### 构建所有驱动
```bash
make all
# 或者
make build-all
```

### 清理单个驱动
```bash
make clean project=<驱动名称>
# 例如：
make clean project=gpio_basic
```

### 清理所有驱动
```bash
make clean-all
```

## 可用驱动列表

当前支持的驱动包括：

- `ammeter_acrel_10r` - 安科瑞10R电表驱动
- `bms_lnxall` - 协能BMS驱动
- `bms_pylon_tech_us108` - 派能科技US108 BMS驱动
- `ess_boost_gold` - 高特EMS驱动
- `ess_boost_lnxall` - 协能EMS驱动
- `ess_pylon_checkwatt` - 派能科技CheckWatt EMS驱动
- `fire_control` - 消防控制驱动
- `gpio_basic` - 基础GPIO驱动
- `pcs_elecod_mac` - 亿兰科PCS驱动
- `pcs_enjoy_basic` - 享能基础PCS驱动
- `pcs_lnxall` - 协能PCS驱动
- `pcs_star_charge_100E` - 星星充电100E PCS驱动
- `sess_basic` - 基础储能站驱动

## 构建输出

### 输出目录
- 驱动文件: `/Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/drivers/`
- 可执行文件: `{项目名}/bin/`

### 文件命名规则
- 驱动文件: `{项目名}_{版本号}.driver`
- 可执行文件: `{项目名}_{版本号}`

## 配置选项

### 环境变量
- `OutDir`: 输出目录路径 (默认: `/Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/drivers`)
- `BuildBin`: 是否构建可执行文件 (默认: 1)

### 构建参数
- `project`: 指定要构建的驱动项目名称
- `outdir`: 指定自定义输出目录路径

## 日志输出说明

- 🟢 **[SUCCESS]**: 操作成功
- 🔴 **[ERROR]**: 操作失败
- 🟡 **[WARN]**: 警告信息
- 🔵 **[INFO]**: 一般信息
- 🟣 **[BUILD]**: 构建过程信息

## 错误处理

Makefile包含完善的错误检查：

1. **项目参数检查**: 确保指定了有效的项目名称
2. **目录存在检查**: 验证驱动目录是否存在
3. **文件存在检查**: 验证main.go文件是否存在
4. **构建结果检查**: 验证构建过程是否成功

## 示例

### 构建单个驱动
```bash
$ make project=gpio_basic
[BUILD] 开始构建驱动: gpio_basic
构建信息:
  项目: gpio_basic
  构建时间: 2025-08-22 15:46:41
  提交哈希: 04e2409b5c95bb1a290012e42123eee2a8b63982
  输出目录: ./test_output

创建输出目录: ./test_output
获取版本信息...
  版本: v1.0.0
  文件名: gpio_basic_v1.0.0

构建插件文件...
[SUCCESS] 插件构建成功
  文件: ./test_output/gpio_basic_v1.0.0.driver (12M)
复制到输出目录...
构建可执行文件...
[SUCCESS] 可执行文件构建成功
  文件: gpio_basic/bin/gpio_basic_v1.0.0 (6.8M)
[BUILD] 驱动构建完成: gpio_basic
=== 构建完成 ===
```

### 批量构建
```bash
$ make all
[INFO] 开始批量构建所有驱动...
构建信息:
  构建时间: 2025-08-22 15:53:46
  提交哈希: 04e2409b5c95bb1a290012e42123eee2a8b63982
  驱动数量: 14

----------------------------------------
[SUCCESS] ammeter_acrel_10r 构建成功
  文件: /Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/drivers/ammeter_acrel_10r_v1.0.0.driver (11M)
----------------------------------------
[SUCCESS] bms_lnxall 构建成功
  文件: /Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/drivers/bms_lnxall_v1.0.0.driver (11M)
----------------------------------------
[ERROR] pcs_elecod_mdc 构建失败
  原因: [ERROR] main.go 文件不存在: pcs_elecod_mdc/main.go
----------------------------------------
[ERROR] pcs_enjoy_basic 构建失败
  原因: [ERROR] 插件构建失败
  原因: 
----------------------------------------
批量构建完成!
  成功: 12
  失败: 2
  总计: 14
```

## 注意事项

1. 确保在`plugins/plug_drivers`目录下运行Makefile
2. 确保Go环境已正确配置
3. 确保所有依赖包已安装
4. 构建前会自动创建必要的输出目录
5. 版本信息通过`go run main.go version`命令获取

## 故障排除

### 常见问题

1. **"project is not set"错误**
   - 解决方案: 使用`make project=<驱动名称>`或`make all`

2. **"驱动目录不存在"错误**
   - 解决方案: 检查驱动名称是否正确，使用`make list`查看可用驱动

3. **"main.go文件不存在"错误**
   - 解决方案: 检查驱动目录结构，确保存在main.go文件

4. **构建失败**
   - 解决方案: 检查Go环境、依赖包和代码语法错误

### 调试技巧

1. 使用`make status`查看当前构建状态
2. 使用`make list`查看所有可用驱动
3. 检查输出目录中的文件是否正确生成
4. 查看构建日志中的详细错误信息
