package internal

import (
	"c_window_counter/public"
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

	"github.com/prometheus/prometheus/pkg/labels"
	promtsdb "github.com/prometheus/prometheus/tsdb"
)

const (
	DefaultDBPath = "./out/ptdb"
)

type promDB struct {
	ctx           context.Context
	db            *promtsdb.DB
	sampleCounter public.IWindowCounter // 滑动窗口计数器，用于统计每秒存储样本数
	statusCache   map[string]any        // 状态类型缓存，key: deviceId+pointKey
	statusMutex   sync.RWMutex          // 保护状态缓存的读写锁
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
			panic(errors.Errorf("打开 Prometheus TSDB 失败: %v", err))
		}

		// 创建滑动窗口计数器，用于统计每秒存储样本数
		// 使用1分钟窗口，60个桶，每个桶代表1秒
		sampleCounter := public.NewQPSWindowCounter()

		instance = &promDB{
			ctx:           ctx,
			db:            db,
			sampleCounter: sampleCounter,
			statusCache:   make(map[string]any),
		}
		c_log.BizInfof(ctx, "启动时序数据库！")
	})
	return instance
}

func (p *promDB) SaveDevices(deviceId string, pointValues []*c_base.SPointValue) error {
	if len(pointValues) == 0 {
		return nil
	}

	ts := time.Now().UnixMilli()
	app := p.db.Appender()
	var sampleCount int64 = 0

	for _, pv := range pointValues {
		if pv == nil || pv.IPoint == nil {
			continue
		}

		pointKey := pv.GetKey()
		value := pv.GetValue()

		// 检查是否是状态类型（有 ValueExplain）
		if len(pv.GetValueExplain()) > 0 {
			// 状态类型：检查缓存
			cacheKey := deviceId + ":" + pointKey

			p.statusMutex.RLock()
			cachedValue, exists := p.statusCache[cacheKey]
			p.statusMutex.RUnlock()

			// 状态未变化，跳过保存
			if exists && cachedValue == value {
				continue
			}

			// 状态变化，更新缓存
			p.statusMutex.Lock()
			p.statusCache[cacheKey] = value
			p.statusMutex.Unlock()
		}

		// 保存数值类型数据
		if numericValue, ok := convertToFloat64(value); ok {
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: MetricNameEmsMetric,
				LabelNameType:   string(c_base.StorageTypeDevice),
				LabelNameID:     deviceId,
				LabelNameField:  pointKey,
			}), ts, numericValue)
			if err != nil {
				_ = app.Rollback()
				return err
			}
			sampleCount++
		}
	}

	if sampleCount > 0 {
		p.sampleCounter.IncrementBy(sampleCount)
	}

	return app.Commit()
}

func (p *promDB) SaveProtocolMetrics(protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, metrics map[string]any) error {
	if len(metrics) == 0 {
		return nil
	}
	ts := time.Now().UnixMilli()
	app := p.db.Appender()

	var sampleCount int64 = 0 // 统计本次存储的样本数量
	for field, value := range metrics {
		// 只保存数值类型数据
		if numericValue, ok := convertToFloat64(value); ok {
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric:   MetricNameEmsMetric,
				LabelNameType:     string(c_base.StorageTypeProtocol),
				LabelNameID:       protocolConfig.Id,
				LabelNameField:    field,
				LabelNameDeviceID: deviceConfig.Id,
			}), ts, numericValue)
			if err != nil {
				_ = app.Rollback()
				return err
			}
			sampleCount++
		}
		// 非数值类型数据将被忽略
	}

	// 更新滑动窗口计数器
	if sampleCount > 0 {
		p.sampleCounter.IncrementBy(sampleCount)
	}

	return app.Commit()
}

func (p *promDB) SaveSystemMetrics(measurement string, tags map[string]string, metrics map[string]any) error {
	if len(metrics) == 0 {
		return nil
	}
	ts := time.Now().UnixMilli()
	app := p.db.Appender()

	var sampleCount int64 = 0 // 统计本次存储的样本数量
	for field, value := range metrics {
		// 只保存数值类型数据
		if numericValue, ok := convertToFloat64(value); ok {
			_, err := app.Add(labels.FromMap(map[string]string{
				LabelNameMetric: MetricNameEmsMetric,
				LabelNameType:   string(c_base.StorageTypeSystem),
				LabelNameID:     measurement,
				LabelNameField:  field,
			}), ts, numericValue)
			if err != nil {
				_ = app.Rollback()
				return err
			}
			sampleCount++
		}
		// 非数值类型数据将被忽略
	}

	// 更新滑动窗口计数器
	if sampleCount > 0 {
		p.sampleCounter.IncrementBy(sampleCount)
	}

	return app.Commit()
}

func (p *promDB) GetStorageData(storageType c_base.StorageType, id string, pointKey []string, startTime, endTime *int64, step int) (*c_chart.ChartData, error) {
	// 查询 ems_metric 系列（仅数值类型数据）
	mint := int64(0)
	maxt := int64(1<<63 - 1)
	if startTime != nil && *startTime > 0 {
		mint = *startTime
	}
	if endTime != nil && *endTime > 0 {
		maxt = *endTime
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
	data := make(map[int64]map[string]float64)

	for _, k := range pointKey {
		// 选择器: __name__="ems_metric", type=storageType, id=id, field=k
		matchers := []*labels.Matcher{
			labels.MustNewMatcher(labels.MatchEqual, LabelNameMetric, MetricNameEmsMetric),
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
		chart.AddSeries(seriesMap[k])
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

	// 获取数据库中真正的第一条数据时间
	oldestTime, err := p.getOldestTimestamp()
	if err != nil {
		c_log.BizInfof(p.ctx, "获取最老时间戳失败: %v", err)
		// 如果查询失败，回退到使用Head的MinTime
		if headStats.MinTime > 0 {
			oldestTime := time.UnixMilli(headStats.MinTime)
			stats.OldestTimestamp = &oldestTime
		}
	} else {
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
	if stats.OldestTimestamp != nil && headStats.MaxTime > 0 {
		oldestTimeMs := stats.OldestTimestamp.UnixMilli()
		stats.RetentionTime = (headStats.MaxTime - oldestTimeMs) / 1000 // 转换为秒
	}

	// 计算平均每个序列占用数据大小
	if stats.TotalSeries > 0 && stats.StorageSize > 0 {
		stats.AvgSeriesSize = float64(stats.StorageSize) / float64(stats.TotalSeries)
	}

	// 使用滑动窗口计数器获取每秒存储样本数
	stats.SamplesPerSecond = p.sampleCounter.GetQPS()

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
		labels.MustNewMatcher(labels.MatchEqual, LabelNameMetric, MetricNameEmsMetric),
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

// getOldestTimestamp 获取数据库中真正的第一条数据时间
func (p *promDB) getOldestTimestamp() (time.Time, error) {
	// 创建一个查询器，查询所有时间范围的数据
	q, err := p.db.Querier(context.Background(), 0, time.Now().UnixMilli())
	if err != nil {
		return time.Time{}, err
	}
	defer q.Close()

	// 查询所有ems_metric系列来找到最早的时间戳
	matchers := []*labels.Matcher{
		labels.MustNewMatcher(labels.MatchEqual, LabelNameMetric, MetricNameEmsMetric),
	}

	ss, _, err := q.Select(false, nil, matchers...)
	if err != nil {
		return time.Time{}, err
	}

	var oldestTimestamp int64 = -1
	for ss.Next() {
		s := ss.At()
		it := s.Iterator()
		if it.Next() {
			t, _ := it.At()
			if oldestTimestamp == -1 || t < oldestTimestamp {
				oldestTimestamp = t
			}
		}
	}

	if err := ss.Err(); err != nil {
		return time.Time{}, err
	}

	if oldestTimestamp == -1 {
		return time.Time{}, fmt.Errorf("未找到任何数据")
	}

	return time.UnixMilli(oldestTimestamp), nil
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

// convertToFloat64 将各种数值类型（包括指针类型）转换为 float64
// 支持的类型：int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64
// 以及对应的指针类型：*int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64
// 返回值：转换后的 float64 值和是否转换成功
func convertToFloat64(value any) (float64, bool) {
	if value == nil {
		return 0, false
	}

	switch v := value.(type) {
	// 基本数值类型
	case int:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case uint:
		return float64(v), true
	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true
	case uint64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	case bool:
		if v {
			return 1, true
		} else {
			return 0, true
		}
	// 指针类型
	case *int:
		if v != nil {
			return float64(*v), true
		}
	case *int8:
		if v != nil {
			return float64(*v), true
		}
	case *int16:
		if v != nil {
			return float64(*v), true
		}
	case *int32:
		if v != nil {
			return float64(*v), true
		}
	case *int64:
		if v != nil {
			return float64(*v), true
		}
	case *uint:
		if v != nil {
			return float64(*v), true
		}
	case *uint8:
		if v != nil {
			return float64(*v), true
		}
	case *uint16:
		if v != nil {
			return float64(*v), true
		}
	case *uint32:
		if v != nil {
			return float64(*v), true
		}
	case *uint64:
		if v != nil {
			return float64(*v), true
		}
	case *float32:
		if v != nil {
			return float64(*v), true
		}
	case *float64:
		if v != nil {
			return *v, true
		}
	case *bool:
		if *v {
			return 1, true
		} else {
			return 0, true
		}
	}

	return 0, false
}
