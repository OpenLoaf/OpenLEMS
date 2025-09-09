# Common 模块文件总结

本文档总结了 `common/` 目录下所有文件的功能和用途，用于后续重新归类使用。

## 目录结构概览

```
common/
├── c_alarm/                    # 告警管理模块
├── c_base/                     # 基础接口和类型定义
├── c_chart/                    # 图表数据结构
├── c_device/                   # 设备相关实现
├── c_func/                     # 聚合函数
├── c_log/                      # 日志接口
├── c_proto/                    # 协议相关
├── c_status/                   # 状态管理
├── c_timer/                    # 定时器功能
├── c_type/                     # 设备类型定义
├── go.mod                      # 模块依赖
├── go.sum                      # 依赖校验
├── c_device_manager_si.go      # 设备管理器单例
├── c_alarm_manager_si.go       # 告警管理器单例
└── c_storage_s.go              # 存储实例单例
```

## 文件详细总结

### 根目录文件

#### go.mod
- **功能**: Go 模块依赖管理文件
- **依赖**: github.com/pkg/errors, github.com/shockerli/cvt, gopkg.in/yaml.v3
- **用途**: 定义 common 模块的依赖关系

#### go.sum
- **功能**: Go 模块依赖校验文件
- **用途**: 确保依赖包的完整性和版本一致性

#### c_device_manager_si.go
- **功能**: 设备管理器单例接口
- **核心接口**: IDeviceManager
- **主要方法**: Start(), Shutdown(), GetDeviceById(), IteratorAllDevices() 等
- **用途**: 全局设备管理，提供设备注册、查询、遍历等功能

#### c_alarm_manager_si.go
- **功能**: 告警管理器单例接口
- **主要方法**: RegisterAlarmManager(), GetAlarmManager()
- **用途**: 全局告警管理，提供告警注册和获取功能

#### c_storage_s.go
- **功能**: 存储实例单例接口
- **主要方法**: RegisterStorageInstance(), GetStorageInstance()
- **用途**: 全局存储管理，提供存储实例注册和获取功能

### c_alarm/ - 告警管理模块

#### c_alarm_interface.go
- **功能**: 告警管理器接口定义
- **核心接口**: IAlarmManager
- **主要方法**: CreateAlarmHistory(), IsAlarmIgnored()
- **用途**: 定义告警管理的基本操作接口

#### c_alarm_manager.go
- **功能**: 告警管理器实现
- **用途**: 实现告警管理器的具体逻辑

### c_base/ - 基础接口和类型定义

#### 设备相关
- **c_base_device_i.go**: 设备基础接口定义 (IDevice)
- **c_base_device_config_s.go**: 设备配置结构体
- **c_base_driver_i.go**: 驱动接口定义 (IDriver)
- **c_base_driver_info_s.go**: 驱动信息结构体
- **c_base_driver_info_f.go**: 驱动信息相关函数
- **c_base_driver_service_s.go**: 驱动服务结构体

#### 协议相关
- **c_base_protocol_i_s.go**: 协议接口和结构体定义
- **c_base_protocol_config_s.go**: 协议配置结构体
- **c_base_protocol_task_i_s.go**: 协议任务接口和结构体
- **c_base_protocol_type_e.go**: 协议类型枚举

#### 存储相关
- **c_base_storage_i.go**: 存储接口定义 (IStorage)
- **c_base_storage_config_s.go**: 存储配置结构体
- **c_base_storage_type_e.go**: 存储类型枚举

#### 告警相关
- **c_base_alarm_i.go**: 告警接口定义 (IAlarm)
- **c_base_alarm_level_e.go**: 告警级别枚举
- **c_base_alarm_action_e.go**: 告警动作枚举
- **c_base_alarm_point_c.go**: 告警点常量定义

#### 策略相关
- **c_base_policy_i.go**: 策略接口定义 (IPolicy)

#### 元数据相关
- **c_base_meta_s.go**: 元数据结构体定义
- **c_base_meta_read_type_e.go**: 元数据读取类型枚举
- **c_base_meta_system_type_e.go**: 元数据系统类型枚举

#### 点数据相关
- **c_base_point_s.go**: 点数据结构体定义
- **c_base_point_s_test.go**: 点数据结构体测试文件

#### 配置相关
- **c_base_config_struct_fields_s.go**: 配置结构体字段定义
- **c_base_config_struct_fields_f.go**: 配置结构体字段相关函数
- **c_base_config_fields_component_type_e.go**: 配置字段组件类型枚举

#### 类型和常量
- **c_base_type_e.go**: 设备类型枚举 (EDeviceType)
- **c_base_const_c.go**: 基础常量定义
- **c_base_endianness_e.go**: 字节序枚举
- **c_base_energy_store_grid_e.go**: 储能电网枚举
- **c_base_setting_group_e.go**: 设置组枚举
- **c_base_server_state_e.go**: 服务器状态枚举

#### 字符串转换文件
- **c_base_alarm_action_e_string.go**: 告警动作枚举字符串转换
- **c_base_alarm_level_e_string.go**: 告警级别枚举字符串转换
- **c_base_config_fields_component_type_e_string.go**: 配置字段组件类型枚举字符串转换
- **c_base_endianness_e_string.go**: 字节序枚举字符串转换
- **c_base_server_state_e_string.go**: 服务器状态枚举字符串转换

#### 文档
- **DECODER_BYTES_README.md**: 字节解码器使用说明文档

### c_chart/ - 图表数据结构

#### c_chart_const.go
- **功能**: 图表相关常量定义
- **用途**: 定义图表类型、样式等常量

#### c_chart_data_s_f.go
- **功能**: 图表数据结构体和相关函数
- **核心结构**: ChartData
- **主要方法**: AddTimestamp(), AddSeries(), NewChartData()
- **用途**: 图表数据的创建和操作

#### c_chart_series_s_f.go
- **功能**: 图表系列结构体和相关函数
- **用途**: 图表数据系列的定义和操作

#### c_chart_xaxis_s_f.go
- **功能**: 图表X轴结构体和相关函数
- **用途**: 图表X轴的定义和操作

### c_device/ - 设备相关实现

#### c_device_real_s.go
- **功能**: 真实设备实现
- **核心结构**: SRealDeviceImpl[P]
- **主要方法**: NewRealDevice(), GetConfig(), GetProtocolStatus() 等
- **用途**: 真实设备的通用实现，支持泛型协议

#### c_device_real_with_protocol_f.go
- **功能**: 带协议的真实设备工厂函数
- **用途**: 创建带协议的真实设备实例

#### c_device_virtual_s.go
- **功能**: 虚拟设备实现
- **用途**: 虚拟设备的实现逻辑

#### c_device_virtual_f.go
- **功能**: 虚拟设备工厂函数
- **用途**: 创建虚拟设备实例

#### c_device_alarm_impl_s.go
- **功能**: 设备告警实现
- **用途**: 设备告警功能的具体实现

#### c_device_policy_s.go
- **功能**: 设备策略结构体
- **用途**: 设备策略的定义和实现

#### c_device_policy_param_s.go
- **功能**: 设备策略参数结构体
- **用途**: 设备策略参数的定义

### c_func/ - 聚合函数

#### c_func_aggregate_avg_f.go
- **功能**: 平均值聚合函数
- **核心函数**: avgAggregate[T], AggregateAvgInt, AggregateAvgFloat32 等
- **用途**: 提供各种数值类型的平均值计算

#### c_func_aggregate_avg_f_test.go
- **功能**: 平均值聚合函数测试
- **用途**: 测试平均值聚合函数的正确性

#### c_func_aggregate_max_f.go
- **功能**: 最大值聚合函数
- **用途**: 提供各种数值类型的最大值计算

#### c_func_aggregate_max_f_test.go
- **功能**: 最大值聚合函数测试
- **用途**: 测试最大值聚合函数的正确性

#### c_func_aggregate_min_f.go
- **功能**: 最小值聚合函数
- **用途**: 提供各种数值类型的最小值计算

#### c_func_aggregate_min_f_test.go
- **功能**: 最小值聚合函数测试
- **用途**: 测试最小值聚合函数的正确性

#### c_func_aggregate_sum_f.go
- **功能**: 求和聚合函数
- **用途**: 提供各种数值类型的求和计算

#### c_func_aggregate_sum_f_test.go
- **功能**: 求和聚合函数测试
- **用途**: 测试求和聚合函数的正确性

#### c_func_aggregate_equal_f.go
- **功能**: 相等性聚合函数
- **用途**: 提供数据相等性比较功能

#### c_func_aggregate_equal_f_test.go
- **功能**: 相等性聚合函数测试
- **用途**: 测试相等性聚合函数的正确性

#### c_base_meta_trigger.go
- **功能**: 元数据触发器
- **用途**: 元数据变化触发相关功能

#### example_test.go
- **功能**: 聚合函数使用示例
- **用途**: 展示聚合函数的使用方法

### c_log/ - 日志接口

#### c_log_interface.go
- **功能**: 日志接口定义
- **核心接口**: ILogger
- **核心结构**: LogLine, LogQueryParams, LogQueryResult
- **主要方法**: Debug(), Info(), Warning(), Error(), QueryLogs()
- **用途**: 定义统一的日志接口，支持结构化日志和日志查询

#### c_log_proxy.go
- **功能**: 日志代理实现
- **用途**: 实现日志接口的具体逻辑

### c_proto/ - 协议相关

#### c_modbus_protocol_i.go
- **功能**: Modbus协议接口定义
- **核心接口**: IModbusProtocol
- **主要方法**: ReadSingleSync(), ReadGroupSync(), WriteSingleRegister() 等
- **用途**: 定义Modbus协议的基本操作接口

#### c_modbus_config_s.go
- **功能**: Modbus配置结构体
- **用途**: Modbus协议的配置参数定义

#### c_modbus_config_s_test.go
- **功能**: Modbus配置结构体测试
- **用途**: 测试Modbus配置的正确性

#### c_modbus_task_s.go
- **功能**: Modbus任务结构体
- **用途**: Modbus协议任务的定义

#### c_modbus_read_function_e.go
- **功能**: Modbus读取功能枚举
- **用途**: 定义Modbus协议支持的各种读取功能

#### c_modbus_read_function_e_string.go
- **功能**: Modbus读取功能枚举字符串转换
- **用途**: 提供枚举值的字符串表示

### c_status/ - 状态管理

#### c_status_protocol_e.go
- **功能**: 协议状态枚举
- **枚举值**: EProtocolDisconnected, EProtocolConnecting, EProtocolConnected
- **用途**: 定义协议连接状态

#### c_status_protocol_e_string.go
- **功能**: 协议状态枚举字符串转换
- **用途**: 提供协议状态枚举的字符串表示

#### c_status_energy_store_e.go
- **功能**: 储能状态枚举
- **用途**: 定义储能设备的状态

#### c_status_energy_store_e_string.go
- **功能**: 储能状态枚举字符串转换
- **用途**: 提供储能状态枚举的字符串表示

#### c_status_pv_e.go
- **功能**: 光伏状态枚举
- **用途**: 定义光伏设备的状态

#### c_status_pv_e_string.go
- **功能**: 光伏状态枚举字符串转换
- **用途**: 提供光伏状态枚举的字符串表示

### c_timer/ - 定时器功能

#### timer.go
- **功能**: 定时器调度器实现
- **核心函数**: SetInterval(), SetTimeout()
- **核心结构**: scheduler, scheduledTask, taskHeap
- **用途**: 提供定时任务调度功能，支持周期性任务和一次性任务

### c_type/ - 设备类型定义

#### 电池管理系统 (BMS)
- **type_bms_i.go**: BMS设备接口定义 (IBms, IBmsBasic)
- **type_bms_e.go**: BMS状态枚举
- **type_bms_e_string.go**: BMS状态枚举字符串转换

#### 电表相关
- **type_ammeter_i.go**: 电表设备接口定义

#### 充电相关
- **type_charge_i.go**: 充电设备接口定义

#### 制冷相关
- **type_cooling_ac_i.go**: 制冷空调设备接口定义
- **type_cooling_i_e.go**: 制冷设备状态枚举
- **type_cooling_liquid_i.go**: 液冷设备接口定义

#### 储能相关
- **type_engery_store_i.go**: 储能设备接口定义
- **type_station_energy_store_i_v.go**: 总站储能设备接口定义

#### 消防相关
- **type_fire_i.go**: 消防设备接口定义

#### 发电机相关
- **type_generator_i.go**: 发电机设备接口定义

#### GPIO相关
- **type_gpio_i_e.go**: GPIO设备状态枚举

#### 温湿度相关
- **type_humiture_i.go**: 温湿度设备接口定义

#### 负载相关
- **type_load_i.go**: 负载设备接口定义

#### PCS相关
- **type_pcs_i_e.go**: PCS设备状态枚举

#### 光伏相关
- **type_pv_i.go**: 光伏设备接口定义

#### 总站入口相关
- **type_station_entrance_i_v.go**: 总站入口设备接口定义

## 模块依赖关系

### 核心依赖
- **c_base**: 所有其他模块的基础，提供核心接口和类型定义
- **c_device**: 依赖 c_base，实现设备相关功能
- **c_proto**: 依赖 c_base，实现协议相关功能
- **c_alarm**: 依赖 c_base，实现告警相关功能

### 功能模块
- **c_func**: 依赖 c_base，提供聚合函数
- **c_chart**: 独立模块，提供图表数据结构
- **c_log**: 独立模块，提供日志接口
- **c_timer**: 独立模块，提供定时器功能
- **c_status**: 独立模块，提供状态枚举
- **c_type**: 依赖 c_base，提供设备类型定义

### 单例管理
- **c_device_manager_si.go**: 设备管理器单例
- **c_alarm_manager_si.go**: 告警管理器单例
- **c_storage_s.go**: 存储实例单例

## 重新归类建议

### 按功能分类
1. **基础层** (c_base): 核心接口、类型、枚举
2. **设备层** (c_device, c_type): 设备实现和类型定义
3. **协议层** (c_proto): 协议相关功能
4. **业务层** (c_alarm, c_func, c_chart): 业务逻辑和数据处理
5. **基础设施层** (c_log, c_timer, c_status): 基础设施功能
6. **管理层** (单例文件): 全局管理功能

### 按职责分类
1. **接口定义**: 所有 *_i.go 文件
2. **结构体定义**: 所有 *_s.go 文件
3. **枚举定义**: 所有 *_e.go 文件
4. **函数实现**: 所有 *_f.go 文件
5. **常量定义**: 所有 *_c.go 文件
6. **测试文件**: 所有 *_test.go 文件
7. **字符串转换**: 所有 *_string.go 文件

### 按模块大小分类
1. **核心模块**: c_base (最大，包含最多文件)
2. **中等模块**: c_device, c_type, c_func
3. **小型模块**: c_alarm, c_chart, c_log, c_proto, c_status, c_timer
4. **单例文件**: 根目录下的单例管理文件

## 总结

`common/` 模块是一个设计良好的共享库，包含了能源管理系统的核心接口、类型定义和基础功能。模块结构清晰，职责分明，为整个系统提供了坚实的基础。建议在重新归类时保持现有的模块划分，同时可以考虑将一些相关的功能进行合并或拆分，以提高代码的可维护性和可读性。
