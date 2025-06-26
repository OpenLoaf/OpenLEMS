package internal

import (
	"common/c_base"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceDataResult 设备数据查询结果
type DeviceDataResult struct {
	DeviceID   string                 `json:"device_id"`
	DeviceType string                 `json:"device_type"`
	Timestamp  int64                  `json:"timestamp"`
	Fields     map[string]interface{} `json:"fields"`
}

// ProtocolMetricsResult 协议指标数据查询结果
type ProtocolMetricsResult struct {
	DeviceID        string                 `json:"device_id"`
	DeviceName      string                 `json:"device_name"`
	ProtocolID      string                 `json:"protocol_id"`
	ProtocolAddress string                 `json:"protocol_address"`
	ProtocolType    string                 `json:"protocol_type"`
	Timestamp       int64                  `json:"timestamp"`
	Metrics         map[string]interface{} `json:"metrics"`
}

// SystemMetricsResult 系统指标数据查询结果
type SystemMetricsResult struct {
	Measurement string                 `json:"measurement"`
	Tags        map[string]string      `json:"tags"`
	Timestamp   int64                  `json:"timestamp"`
	Metrics     map[string]interface{} `json:"metrics"`
}

// GetDeviceData 根据设备ID获取设备数据
func (p *Pebbledb) GetDeviceData(deviceID string, limit int) ([]*DeviceDataResult, error) {
	if limit <= 0 {
		limit = 100 // 默认限制100条
	}

	var results []*DeviceDataResult
	// 修复：使用通配符来匹配所有设备类型下的设备ID
	prefix := fmt.Sprintf("devices/")

	iter, _ := p.db.NewIter(&pebble.IterOptions{
		LowerBound: []byte(prefix),
		UpperBound: []byte(fmt.Sprintf("%s~", prefix)),
	})
	defer iter.Close()

	for iter.First(); iter.Valid() && len(results) < limit; iter.Next() {
		// 检查键名是否包含目标设备ID
		key := string(iter.Key())
		// 键名格式：devices/{deviceType}/{deviceId}/{timestamp}
		parts := strings.Split(key, "/")
		if len(parts) < 3 || parts[2] != deviceID {
			continue
		}

		var data map[string]interface{}
		if err := json.Unmarshal(iter.Value(), &data); err != nil {
			g.Log().Errorf(p.ctx, "解析设备数据失败: %v, key: %s", err, string(iter.Key()))
			continue
		}

		result := &DeviceDataResult{
			DeviceID:   fmt.Sprintf("%v", data["device_id"]),
			DeviceType: fmt.Sprintf("%v", data["device_type"]),
			Timestamp:  int64(data["timestamp"].(float64)),
			Fields:     data["fields"].(map[string]interface{}),
		}
		results = append(results, result)
	}

	// 按时间戳倒序排列（最新的在前）
	sort.Slice(results, func(i, j int) bool {
		return results[i].Timestamp > results[j].Timestamp
	})

	return results, nil
}

// GetDeviceDataByType 根据设备类型获取设备数据
func (p *Pebbledb) GetDeviceDataByType(deviceType c_base.EDeviceType, limit int) ([]*DeviceDataResult, error) {
	if limit <= 0 {
		limit = 100
	}

	var results []*DeviceDataResult
	prefix := fmt.Sprintf("devices/%s/", string(deviceType))

	iter, _ := p.db.NewIter(&pebble.IterOptions{
		LowerBound: []byte(prefix),
		UpperBound: []byte(fmt.Sprintf("%s~", prefix)),
	})
	defer iter.Close()

	for iter.First(); iter.Valid() && len(results) < limit; iter.Next() {
		var data map[string]interface{}
		if err := json.Unmarshal(iter.Value(), &data); err != nil {
			g.Log().Errorf(p.ctx, "解析设备数据失败: %v, key: %s", err, string(iter.Key()))
			continue
		}

		result := &DeviceDataResult{
			DeviceID:   fmt.Sprintf("%v", data["device_id"]),
			DeviceType: fmt.Sprintf("%v", data["device_type"]),
			Timestamp:  int64(data["timestamp"].(float64)),
			Fields:     data["fields"].(map[string]interface{}),
		}
		results = append(results, result)
	}

	// 按时间戳倒序排列
	sort.Slice(results, func(i, j int) bool {
		return results[i].Timestamp > results[j].Timestamp
	})

	return results, nil
}

// GetDeviceDataByTimeRange 根据时间范围获取设备数据
func (p *Pebbledb) GetDeviceDataByTimeRange(deviceID string, startTime, endTime time.Time, limit int) ([]*DeviceDataResult, error) {
	if limit <= 0 {
		limit = 1000
	}

	var results []*DeviceDataResult
	// 修复：使用通配符来匹配所有设备类型下的设备ID
	prefix := fmt.Sprintf("devices/")

	iter, _ := p.db.NewIter(&pebble.IterOptions{
		LowerBound: []byte(prefix),
		UpperBound: []byte(fmt.Sprintf("%s~", prefix)),
	})
	defer iter.Close()

	startTimestamp := startTime.Unix()
	endTimestamp := endTime.Unix()

	for iter.First(); iter.Valid() && len(results) < limit; iter.Next() {
		// 检查键名是否包含目标设备ID
		key := string(iter.Key())
		// 键名格式：devices/{deviceType}/{deviceId}/{timestamp}
		parts := strings.Split(key, "/")
		if len(parts) < 3 || parts[2] != deviceID {
			continue
		}

		var data map[string]interface{}
		if err := json.Unmarshal(iter.Value(), &data); err != nil {
			continue
		}

		timestamp := int64(data["timestamp"].(float64))
		if timestamp >= startTimestamp && timestamp <= endTimestamp {
			result := &DeviceDataResult{
				DeviceID:   fmt.Sprintf("%v", data["device_id"]),
				DeviceType: fmt.Sprintf("%v", data["device_type"]),
				Timestamp:  timestamp,
				Fields:     data["fields"].(map[string]interface{}),
			}
			results = append(results, result)
		}
	}

	// 按时间戳排序
	sort.Slice(results, func(i, j int) bool {
		return results[i].Timestamp > results[j].Timestamp
	})

	return results, nil
}

// GetProtocolMetrics 获取协议指标数据
func (p *Pebbledb) GetProtocolMetrics(deviceID string, limit int) ([]*ProtocolMetricsResult, error) {
	if limit <= 0 {
		limit = 100
	}

	var results []*ProtocolMetricsResult
	prefix := fmt.Sprintf("protocol_metrics/%s", deviceID)

	iter, _ := p.db.NewIter(&pebble.IterOptions{
		LowerBound: []byte(prefix),
		UpperBound: []byte(fmt.Sprintf("%s~", prefix)),
	})
	defer iter.Close()

	for iter.First(); iter.Valid() && len(results) < limit; iter.Next() {
		var data map[string]interface{}
		if err := json.Unmarshal(iter.Value(), &data); err != nil {
			g.Log().Errorf(p.ctx, "解析协议指标数据失败: %v, key: %s", err, string(iter.Key()))
			continue
		}

		result := &ProtocolMetricsResult{
			DeviceID:        fmt.Sprintf("%v", data["device_id"]),
			DeviceName:      fmt.Sprintf("%v", data["device_name"]),
			ProtocolID:      fmt.Sprintf("%v", data["protocol_id"]),
			ProtocolAddress: fmt.Sprintf("%v", data["protocol_address"]),
			ProtocolType:    fmt.Sprintf("%v", data["protocol_type"]),
			Timestamp:       int64(data["timestamp"].(float64)),
			Metrics:         data["metrics"].(map[string]interface{}),
		}
		results = append(results, result)
	}

	// 按时间戳倒序排列
	sort.Slice(results, func(i, j int) bool {
		return results[i].Timestamp > results[j].Timestamp
	})

	return results, nil
}

// GetSystemMetrics 获取系统指标数据
func (p *Pebbledb) GetSystemMetrics(measurement string, limit int) ([]*SystemMetricsResult, error) {
	if limit <= 0 {
		limit = 100
	}

	var results []*SystemMetricsResult
	prefix := fmt.Sprintf("system_metrics/%s", measurement)

	iter, _ := p.db.NewIter(&pebble.IterOptions{
		LowerBound: []byte(prefix),
		UpperBound: []byte(fmt.Sprintf("%s~", prefix)),
	})
	defer iter.Close()

	for iter.First(); iter.Valid() && len(results) < limit; iter.Next() {
		var data map[string]interface{}
		if err := json.Unmarshal(iter.Value(), &data); err != nil {
			g.Log().Errorf(p.ctx, "解析系统指标数据失败: %v, key: %s", err, string(iter.Key()))
			continue
		}

		result := &SystemMetricsResult{
			Measurement: fmt.Sprintf("%v", data["measurement"]),
			Timestamp:   int64(data["timestamp"].(float64)),
			Metrics:     data["metrics"].(map[string]interface{}),
		}

		// 安全处理tags字段
		if tagsInterface, ok := data["tags"]; ok {
			if tagsMap, ok := tagsInterface.(map[string]interface{}); ok {
				result.Tags = make(map[string]string)
				for k, v := range tagsMap {
					result.Tags[k] = fmt.Sprintf("%v", v)
				}
			}
		}

		results = append(results, result)
	}

	// 按时间戳倒序排列
	sort.Slice(results, func(i, j int) bool {
		return results[i].Timestamp > results[j].Timestamp
	})

	return results, nil
}

// GetAllKeys 获取所有键名（用于调试）
func (p *Pebbledb) GetAllKeys(prefix string, limit int) ([]string, error) {
	if limit <= 0 {
		limit = 1000
	}

	var keys []string
	var iterOptions *pebble.IterOptions

	if prefix != "" {
		iterOptions = &pebble.IterOptions{
			LowerBound: []byte(prefix),
			UpperBound: []byte(fmt.Sprintf("%s~", prefix)),
		}
	}

	iter, _ := p.db.NewIter(iterOptions)
	defer iter.Close()

	for iter.First(); iter.Valid() && len(keys) < limit; iter.Next() {
		keys = append(keys, string(iter.Key()))
	}

	return keys, nil
}

// GetDataByKey 根据完整键名获取数据（用于调试）
func (p *Pebbledb) GetDataByKey(key string) (map[string]interface{}, error) {
	value, closer, err := p.db.Get([]byte(key))
	if err != nil {
		if err == pebble.ErrNotFound {
			return nil, gerror.Newf("数据不存在: %s", key)
		}
		return nil, gerror.Wrapf(err, "获取数据失败: %s", key)
	}
	defer closer.Close()

	var data map[string]interface{}
	if err := json.Unmarshal(value, &data); err != nil {
		return nil, gerror.Wrapf(err, "解析数据失败: %s", key)
	}

	return data, nil
}

// GetStats 获取数据库统计信息
func (p *Pebbledb) GetStats() map[string]interface{} {
	stats := make(map[string]interface{})

	// 获取各类型数据的数量
	deviceCount := p.countKeysWithPrefix("devices/")
	protocolCount := p.countKeysWithPrefix("protocol_metrics/")
	systemCount := p.countKeysWithPrefix("system_metrics/")

	stats["device_data_count"] = deviceCount
	stats["protocol_metrics_count"] = protocolCount
	stats["system_metrics_count"] = systemCount
	stats["total_count"] = deviceCount + protocolCount + systemCount
	stats["db_path"] = p.pebbleConfig.Path

	return stats
}

// countKeysWithPrefix 计算指定前缀的键数量
func (p *Pebbledb) countKeysWithPrefix(prefix string) int {
	count := 0
	iter, _ := p.db.NewIter(&pebble.IterOptions{
		LowerBound: []byte(prefix),
		UpperBound: []byte(fmt.Sprintf("%s~", prefix)),
	})
	defer iter.Close()

	for iter.First(); iter.Valid(); iter.Next() {
		count++
	}

	return count
}

// GetLatestDeviceData 获取设备的最新数据
func (p *Pebbledb) GetLatestDeviceData(deviceID string) (*DeviceDataResult, error) {
	results, err := p.GetDeviceData(deviceID, 1)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, gerror.Newf("没有找到设备数据: %s", deviceID)
	}

	return results[0], nil
}

// ParseTimestampFromKey 从键名中解析时间戳（辅助函数）
func ParseTimestampFromKey(key string) (int64, error) {
	parts := strings.Split(key, "/")
	if len(parts) < 4 {
		return 0, gerror.Newf("无效的键格式: %s", key)
	}

	timestampStr := parts[len(parts)-1]
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return 0, gerror.Wrapf(err, "解析时间戳失败: %s", timestampStr)
	}

	return timestamp, nil
}
