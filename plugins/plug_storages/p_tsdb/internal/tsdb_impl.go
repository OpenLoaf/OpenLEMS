package internal

import (
	"common/c_base"
	"common/c_chart"
	"common/c_log"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/shockerli/cvt"

	"github.com/prometheus/prometheus/pkg/labels"
	promtsdb "github.com/prometheus/prometheus/tsdb"
)

const (
	DefaultDBPath     = "./out/ptdb"
	LabelNameMetric   = "__name__"
	LabelNameType     = "type" // device/protocol/system
	LabelNameID       = "id"   // deviceId/protocolId/measurement
	LabelNameField    = "field"
	LabelNameDeviceID = "device_id" // 设备ID
)

type promDB struct {
	ctx context.Context
	db  *promtsdb.DB
}

var (
	instance c_base.IStorage
	once     sync.Once
	mtx      sync.Mutex
)

func NewPromTSDB(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	once.Do(func() {
		mtx.Lock()
		defer mtx.Unlock()
		if instance != nil {
			return
		}

		basePath := DefaultDBPath
		if storageConfig != nil && storageConfig.Params != nil {
			if v, ok := storageConfig.Params["path"]; ok && v != "" {
				basePath = v
			}
		}

		// Use default options; caller controls retention via external cleanup or tsdb options if needed
		db, err := promtsdb.Open(filepath.Clean(basePath), nil, nil, nil)
		if err != nil {
			panic(errors.Errorf("打开 Prometheus TSDB 失败: %w", err.Error()))
		}

		instance = &promDB{ctx: ctx, db: db}
		c_log.BizInfof(ctx, "启动时序数据库！")
	})
	return instance
}

func (p *promDB) SaveDevices(deviceId string, deviceType c_base.EDeviceType, fields map[string]any) error {
	if len(fields) == 0 {
		return nil
	}
	ts := time.Now().UnixMilli()
	app := p.db.Appender()
	for field, value := range fields {
		// 统一转为 float64，如果失败写入字符串的 JSON 序列
		switch v := value.(type) {
		case int:
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: "ems_metric",
				LabelNameType:   string(c_base.StorageTypeDevice),
				LabelNameID:     deviceId,
				LabelNameField:  field,
			}), ts, float64(v))
			if err != nil {
				_ = app.Rollback()
				return err
			}
		case int32:
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: "ems_metric",
				LabelNameType:   string(c_base.StorageTypeDevice),
				LabelNameID:     deviceId,
				LabelNameField:  field,
			}), ts, float64(v))
			if err != nil {
				_ = app.Rollback()
				return err
			}
		case int64:
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: "ems_metric",
				LabelNameType:   string(c_base.StorageTypeDevice),
				LabelNameID:     deviceId,
				LabelNameField:  field,
			}), ts, float64(v))
			if err != nil {
				_ = app.Rollback()
				return err
			}
		case float32:
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: "ems_metric",
				LabelNameType:   string(c_base.StorageTypeDevice),
				LabelNameID:     deviceId,
				LabelNameField:  field,
			}), ts, float64(v))
			if err != nil {
				_ = app.Rollback()
				return err
			}
		case float64:
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: "ems_metric",
				LabelNameType:   string(c_base.StorageTypeDevice),
				LabelNameID:     deviceId,
				LabelNameField:  field,
			}), ts, v)
			if err != nil {
				_ = app.Rollback()
				return err
			}
		default:
			// 对非数值类型，将值 JSON 序列化后，附加一个 *_text 序列保存为 0/1
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: "ems_metric_text",
				LabelNameType:   string(c_base.StorageTypeDevice),
				LabelNameID:     deviceId,
				LabelNameField:  field,
			}), ts, 1)
			if err != nil {
				_ = app.Rollback()
				return err
			}
		}
	}
	return app.Commit()
}

func (p *promDB) SaveProtocolMetrics(protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, metrics map[string]any) error {
	if len(metrics) == 0 {
		return nil
	}
	ts := time.Now().UnixMilli()
	app := p.db.Appender()
	for field, v := range metrics {
		val, err := cvt.Float64E(v)
		if err != nil {
			// 同 SaveDevices 的处理
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric:   "ems_metric_text",
				LabelNameType:     string(c_base.StorageTypeProtocol),
				LabelNameID:       protocolConfig.Id,
				LabelNameField:    field,
				LabelNameDeviceID: deviceConfig.Id,
			}), ts, 1)
			if err != nil {
				_ = app.Rollback()
				return err
			}
			continue
		}
		_, err = app.Add(labels.FromMap(map[string]string{
			LabelNameMetric:   "ems_metric",
			LabelNameType:     string(c_base.StorageTypeProtocol),
			LabelNameID:       protocolConfig.Id,
			LabelNameField:    field,
			LabelNameDeviceID: deviceConfig.Id,
		}), ts, val)
		if err != nil {
			_ = app.Rollback()
			return err
		}
	}
	return app.Commit()
}

func (p *promDB) SaveSystemMetrics(measurement string, tags map[string]string, metrics map[string]any) error {
	if len(metrics) == 0 {
		return nil
	}
	ts := time.Now().UnixMilli()
	app := p.db.Appender()
	for field, v := range metrics {
		val, err := cvt.Float64E(v)
		if err != nil {
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: "ems_metric_text",
				LabelNameType:   string(c_base.StorageTypeSystem),
				LabelNameID:     measurement,
				LabelNameField:  field,
			}), ts, 1)
			if err != nil {
				_ = app.Rollback()
				return err
			}
			continue
		}
		_, err = app.Add(labels.FromMap(map[string]string{
			LabelNameMetric: "ems_metric",
			LabelNameType:   string(c_base.StorageTypeSystem),
			LabelNameID:     measurement,
			LabelNameField:  field,
		}), ts, val)
		if err != nil {
			_ = app.Rollback()
			return err
		}
	}
	return app.Commit()
}

func (p *promDB) GetStorageData(storageType c_base.StorageType, id string, pointKey []string, startTime, endTime *int, step int) (*c_chart.ChartData, error) {
	// 查询 ems_metric 和 ems_metric_text 系列
	mint := int64(0)
	maxt := int64(1<<63 - 1)
	if startTime != nil && *startTime > 0 {
		mint = int64(*startTime)
	}
	if endTime != nil && *endTime > 0 {
		maxt = int64(*endTime)
	}

	q, err := p.db.Querier(context.Background(), mint, maxt)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// 构建返回结构
	chart := c_chart.NewChartData(len(pointKey))
	seriesMap := make(map[string]*c_chart.Series)
	for _, k := range pointKey {
		seriesMap[k] = c_chart.NewSeries(k, c_chart.ChartTypeLine, 0)
	}

	// 逐点位查询并合并时间戳
	// 收集所有样本 (timestamp -> map[field]value)
	type sample struct {
		ts  int64
		val float64
	}
	data := make(map[int64]map[string]float64)

	for _, k := range pointKey {
		// 选择器: __name__="ems_metric", type=storageType, id=id, field=k
		matchers := []*labels.Matcher{
			labels.MustNewMatcher(labels.MatchEqual, LabelNameMetric, "ems_metric"),
			labels.MustNewMatcher(labels.MatchEqual, LabelNameType, string(storageType)),
			labels.MustNewMatcher(labels.MatchEqual, LabelNameID, id),
			labels.MustNewMatcher(labels.MatchEqual, LabelNameField, k),
		}
		ss, _, err := q.Select(false, nil, matchers...)
		if err != nil {
			return nil, err
		}
		for ss.Next() {
			s := ss.At()
			it := s.Iterator()
			for it.Next() {
				t, v := it.At()
				if _, ok := data[t]; !ok {
					data[t] = make(map[string]float64)
				}
				data[t][k] = v
			}
		}
		if err := ss.Err(); err != nil {
			return nil, err
		}
	}

	// 排序时间戳
	var timestamps []int64
	for ts := range data {
		timestamps = append(timestamps, ts)
	}
	sort.Slice(timestamps, func(i, j int) bool { return timestamps[i] < timestamps[j] })

	// 步长过滤（按毫秒）
	var nextAllowed int64
	hasAnchor := false
	for _, ts := range timestamps {
		if step > 0 {
			if !hasAnchor {
				hasAnchor = true
				nextAllowed = ts + int64(step)
			} else if ts < nextAllowed {
				continue
			} else {
				nextAllowed = ts + int64(step)
			}
		}
		chart.AddTimestamp(ts)
		for _, k := range pointKey {
			if v, ok := data[ts][k]; ok {
				seriesMap[k].AppendData(fmt.Sprintf("%v", v))
			} else {
				seriesMap[k].AppendData("")
			}
		}
	}

	for _, k := range pointKey {
		chart.AddSeries(*seriesMap[k])
	}
	return chart, nil
}

func (p *promDB) GetStorageStats() (*c_base.StorageStats, error) {
	stats := &c_base.StorageStats{}

	// 获取数据库头部信息
	head := p.db.Head()
	if head == nil {
		return stats, fmt.Errorf("无法获取数据库头部信息")
	}

	// 获取头部统计信息
	headStats := head.Stats("")
	stats.TotalSeries = int64(headStats.NumSeries)
	// 注意：Prometheus TSDB的Stats结构体没有NumSamples字段，我们通过查询来估算
	stats.TotalSamples = p.estimateTotalSamples()

	// 转换时间戳为time.Time类型
	if headStats.MinTime > 0 {
		oldestTime := time.UnixMilli(headStats.MinTime)
		stats.OldestTimestamp = &oldestTime
	}
	if headStats.MaxTime > 0 {
		newestTime := time.UnixMilli(headStats.MaxTime)
		stats.NewestTimestamp = &newestTime
	}

	// 计算存储大小（通过数据库目录大小估算）
	storageSize, err := p.calculateStorageSize()
	if err != nil {
		c_log.BizInfof(p.ctx, "计算存储大小失败: %v", err)
		stats.StorageSize = -1 // 表示无法获取
	} else {
		stats.StorageSize = storageSize
	}

	// 计算数据保留时间（秒）
	if headStats.MinTime > 0 && headStats.MaxTime > 0 {
		stats.RetentionTime = (headStats.MaxTime - headStats.MinTime) / 1000 // 转换为秒
	}

	// 计算平均每个序列占用数据大小
	if stats.TotalSeries > 0 && stats.StorageSize > 0 {
		stats.AvgSeriesSize = float64(stats.StorageSize) / float64(stats.TotalSeries)
	}

	// 计算每秒存储样本数
	if stats.RetentionTime > 0 && stats.TotalSamples > 0 {
		stats.SamplesPerSecond = float64(stats.TotalSamples) / float64(stats.RetentionTime)
	}

	// 计算存储大小（MB）
	stats.StorageSizeMB = float64(stats.StorageSize) / (1024 * 1024)

	// 计算数据保留时间（小时）
	stats.RetentionHours = float64(stats.RetentionTime) / 3600

	c_log.Debugf(p.ctx, "获取存储统计信息: 序列数=%d, 样本数=%d, 存储大小=%.2fMB, 保留时间=%.2f小时, 平均序列大小=%.2f字节, 每秒样本数=%.2f",
		stats.TotalSeries, stats.TotalSamples, stats.StorageSizeMB, stats.RetentionHours, stats.AvgSeriesSize, stats.SamplesPerSecond)

	return stats, nil
}

// estimateTotalSamples 估算总样本数量
func (p *promDB) estimateTotalSamples() int64 {
	// 通过查询所有ems_metric系列来估算样本数量
	head := p.db.Head()
	if head == nil {
		return 0
	}

	// 创建查询器
	q, err := p.db.Querier(context.Background(), 0, time.Now().UnixMilli())
	if err != nil {
		return 0
	}
	defer q.Close()

	// 查询所有ems_metric系列
	matchers := []*labels.Matcher{
		labels.MustNewMatcher(labels.MatchEqual, LabelNameMetric, "ems_metric"),
	}

	ss, _, err := q.Select(false, nil, matchers...)
	if err != nil {
		return 0
	}

	var totalSamples int64
	for ss.Next() {
		s := ss.At()
		it := s.Iterator()
		for it.Next() {
			totalSamples++
		}
	}

	return totalSamples
}

// calculateStorageSize 计算存储目录大小
func (p *promDB) calculateStorageSize() (int64, error) {
	// 获取数据库路径
	dbPath := p.db.Dir()
	if dbPath == "" {
		return 0, fmt.Errorf("无法获取数据库路径")
	}

	// 遍历目录计算总大小
	var totalSize int64
	err := filepath.Walk(dbPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	return totalSize, err
}

func (p *promDB) Close() {
	if p.db != nil {
		_ = p.db.Close()
		c_log.BizInfof(p.ctx, "关闭时序数据库！")
	}
}
