# @cpp: MPC卡尔曼滤波器 C++ 库 + Go 绑定

本模块包含用于微电网预测的 MPC 卡尔曼滤波器的 C++ 实现，以及使用 `c-for-go` 生成的 Go 绑定。

## 功能特性

- **MPC 卡尔曼滤波器**: 针对微电网负载/光伏预测优化
- **状态空间模型**: [值, 速度, 加速度] 带趋势衰减
- **不确定性量化**: 提供预测置信区间
- **在线学习**: 持续适应新测量数据
- **C++ 性能**: 高速计算与 Go 集成

## 目录结构

- `include/hexlib/hexlib.h`: 包含 MPC 函数的 C API 头文件
- `src/mpc_kalman.cpp`: MPC 卡尔曼滤波器实现
- `src/hexlib.cpp`: 基础工具函数
- `CMakeLists.txt` 和 `Makefile`: 构建脚本
- `bindings/`: c-for-go 配置
- `hexlib/`: 生成的 Go 绑定

## API 函数

### 核心 MPC 函数

```c
// 从历史数据创建预测器
void* mpc_create_predictor(const double* historical_data, int data_count);

// 生成多步预测及不确定性
PredictionArray* mpc_predict(void* predictor, int prediction_steps);

// 使用新测量值更新
void mpc_update_measurement(void* predictor, double measurement);

// 获取当前滤波状态
PredictionResult mpc_get_current_state(void* predictor);

// 清理资源
void mpc_free_predictor(void* predictor);
void mpc_free_prediction_array(PredictionArray* array);
```

## 使用方法

### 构建库:
```bash
cd @cpp
make build
```

### 生成 Go 绑定:
```bash
cd @cpp
make generate
```

### Go 集成示例:
```go
// 使用历史数据创建预测器
historicalData := []float64{50.0, 50.0, 45.0, 40.0, 35.0}
predictor := hexlib.MpcCreatePredictor(&historicalData[0], len(historicalData))
defer hexlib.MpcFreePredictor(predictor)

// 获取预测结果
predictions := hexlib.MpcPredict(predictor, 10)
defer hexlib.MpcFreePredictionArray(predictions)

// 使用新测量值更新
hexlib.MpcUpdateMeasurement(predictor, 38.0)
```

## 与 @internal/ 的集成

生成的 Go 绑定可以在 `@internal/` 目录中导入和使用，用于 EMS 预测服务。MPC 卡尔曼滤波器提供：

1. **实时预测**: 适用于 MPC 控制循环
2. **不确定性量化**: 风险感知决策制定
3. **自适应学习**: 持续模型改进
4. **高性能**: C++ 计算与 Go 便利性

完美适用于需要可靠短期预测的微电网能源管理系统。


