## tsdb 存储插件（Prometheus TSDB）

本模块基于 Prometheus TSDB 实现 `common/c_base.IStorage` 接口，用于设备、协议与系统指标的时间序列化存储与查询。实现参考 `plugins/plug_storages/pebbledb` 的业务语义，保证相同的行为与返回结构（`common/c_chart.ChartData`）。

### 目录结构

- `export_plug_storages_tsdb.go`：导出工厂函数 `NewStorageInstance`，供外部按配置构建存储实例。
- `internal/tsdb_impl.go`：IStorage 具体实现，封装 TSDB 打开、写入、查询与关闭。

### 依赖与安装

- 主要依赖：`github.com/prometheus/prometheus`（内部使用 `tsdb` 与 `pkg/labels`）。
- 已写入 `go.mod`，在模块目录执行以下命令拉取依赖：

```bash
cd plugins/plug_storages/tsdb
go mod tidy
```

### 配置项

- `SStorageConfig.Params["path"]`：TSDB 数据目录。默认 `./out/ptstdb`。

### 数据写入

- 写入统一映射为 Prometheus time series：
  - metric 名称：
    - 数值：`ems_metric`
    - 非数值：`ems_metric_text`（并附加 `ems_metric_text_len` 记录文本长度）
  - 标签：
    - `type`：`device` | `protocol` | `system`
    - `id`：设备ID/协议ID/measurement
    - `field`：点位名（keys）
  - 时间戳：毫秒（`time.Now().UnixMilli()`）

与 `pebbledb` 类似：

- `SaveDevices(deviceId, deviceType, fields)`：写入 `type=device,id=deviceId,field=<k>`。
- `SaveProtocolMetrics(protocolConfig, deviceConfig, metrics)`：写入 `type=protocol,id=protocolConfig.Id,field=<k>`。
- `SaveSystemMetrics(measurement, tags, metrics)`：写入 `type=system,id=measurement,field=<k>`。

### 数据查询

- `GetStorageData(storageType, id, pointKey, startTime, endTime, step)`：
  - 构造选择器：`__name__="ems_metric", type=<storageType>, id=<id>, field=<point>`。
  - 基于 `Querier(mint, maxt)` 查询指定时间范围。
  - 将多条曲线按时间戳对齐，输出 `c_chart.ChartData`：
    - `XAxis.Data`：毫秒时间戳字符串序列
    - `Series`：每个 `pointKey` 一条 `line` 类型曲线
  - `step`：毫秒步长过滤。以首条数据时间为锚点，仅保留 `ts` 满足 `ts >= nextAllowed` 的点，并将 `nextAllowed = ts + step`。

### 与 pebbledb 的差异

- pebbledb 是 KV 模式，键包含 `prefix/id/timestamp`，value 为 JSON；TSDB 为分系列、按标签维度索引的时间序列，写入/查询性能更友好于时间序列数据。
- 非数值字段在 TSDB 中不会直接保存原文，仅保存存在标记与长度指标，如需原文持久化可在上层引入对象存储或额外 KV。

### 使用示例

```go
// 创建实例
storage := tsdb.NewStorageInstance(ctx, &c_base.SStorageConfig{Params: map[string]string{"path": "./out/ptstdb"}})

// 写入
_ = storage.SaveDevices("dev-1", c_base.EDeviceAmmeter, map[string]any{"soc": 88.2, "status": "ok"})

// 查询
start := int(time.Now().Add(-time.Hour).UnixMilli())
end := int(time.Now().UnixMilli())
chart, _ := storage.GetStorageData(c_base.StorageTypeDevice, "dev-1", []string{"soc"}, &start, &end, 60000)
_ = chart
```

#### 单独查看数据库
```bash
prometheus --config.file=ptdb.yaml --storage.tsdb.path=ptdb/
```


### 注意事项

- Querier 只可见创建时已提交的数据；每次写入建议以 `Appender()` -> `Add()` -> `Commit()` 事务方式提交。
- 时间戳过旧或乱序的写入可能被 TSDB 拒绝（参见 Prometheus TSDB usage）。



