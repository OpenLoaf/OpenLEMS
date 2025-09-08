# TSDB 统计功能说明

## 概述

本模块为 Prometheus TSDB 存储插件添加了统计功能，可以获取数据库的详细统计信息。

## 新增功能

### 1. 存储统计接口

在 `IStorage` 接口中新增了 `GetStorageStats()` 方法：

```go
// GetStorageStats 获取存储统计信息
GetStorageStats() (*StorageStats, error)
```

### 2. 统计信息结构体

定义了 `StorageStats` 结构体来存储统计信息：

```go
type StorageStats struct {
    TotalSeries     int64   `json:"total_series"`     // 总时间序列数量
    TotalSamples    int64   `json:"total_samples"`    // 总样本数量
    StorageSize     int64   `json:"storage_size"`     // 存储大小（字节）
    OldestTimestamp int64   `json:"oldest_timestamp"` // 最老数据时间戳
    NewestTimestamp int64   `json:"newest_timestamp"` // 最新数据时间戳
    RetentionTime   int64   `json:"retention_time"`   // 数据保留时间（秒）
}
```

## 使用方法

### 基本用法

```go
// 创建存储实例
storage := p_tsdb.NewStorageInstance(ctx, storageConfig)

// 获取统计信息
stats, err := storage.GetStorageStats()
if err != nil {
    log.Fatal(err)
}

// 输出统计信息
fmt.Printf("总时间序列数量: %d\n", stats.TotalSeries)
fmt.Printf("总样本数量: %d\n", stats.TotalSamples)
fmt.Printf("存储大小: %d 字节\n", stats.StorageSize)
```

### JSON 格式输出

```go
statsJSON, err := json.MarshalIndent(stats, "", "  ")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%s\n", statsJSON)
```

## 统计信息说明

### TotalSeries (总时间序列数量)
- 表示数据库中存储的唯一时间序列数量
- 通过 Prometheus TSDB 的 Head.Stats() 方法获取
- 每个唯一的标签组合对应一个时间序列

### TotalSamples (总样本数量)
- 表示数据库中存储的样本点总数
- 通过查询所有 `ems_metric` 系列来估算
- 包括所有数值类型的指标数据

### StorageSize (存储大小)
- 表示数据库文件占用的磁盘空间大小（字节）
- 通过遍历数据库目录计算所有文件大小
- 如果计算失败，返回 -1

### OldestTimestamp (最老数据时间戳)
- 表示数据库中最早数据的时间戳（毫秒）
- 通过 Prometheus TSDB 的 Head.Stats() 方法获取

### NewestTimestamp (最新数据时间戳)
- 表示数据库中最新数据的时间戳（毫秒）
- 通过 Prometheus TSDB 的 Head.Stats() 方法获取

### RetentionTime (数据保留时间)
- 表示数据的时间跨度（秒）
- 计算公式：`(NewestTimestamp - OldestTimestamp) / 1000`
- 用于了解数据的时效性

## 示例代码

参考 `example_stats.go` 文件中的完整示例。

## 注意事项

1. **性能考虑**：统计查询可能涉及大量数据扫描，建议在低峰期执行
2. **样本数量估算**：TotalSamples 是通过查询估算的，可能不是精确值
3. **存储大小计算**：StorageSize 通过文件系统遍历计算，对于大型数据库可能较慢
4. **时间戳格式**：所有时间戳均为毫秒级 Unix 时间戳

## API 兼容性

- 新增的统计功能不会影响现有的存储和查询功能
- 所有现有接口保持不变
- 统计方法为可选功能，不影响核心业务流程

