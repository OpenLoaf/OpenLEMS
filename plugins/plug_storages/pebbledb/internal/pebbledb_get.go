package internal

import (
	"common/c_base"
	"common/c_chart"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/cockroachdb/pebble"
)

// 排序相关常量
const (
	SortOrderAsc  = "asc"  // 升序
	SortOrderDesc = "desc" // 降序
)

// 存储前缀常量
const (
	DevicePrefix   = "device"   // 设备前缀
	ProtocolPrefix = "protocol" // 协议前缀
	SystemPrefix   = "system"   // 系统前缀
)

// 数据标签常量
const (
	MetricsTag = "metrics" // 指标标签
)

// 其他常量
const (
	KeySeparator   = "/" // 键分隔符
	BoundarySuffix = "z" // 边界后缀
	EmptyString    = ""  // 空字符串
	ZeroTimestamp  = 0   // 零时间戳
)

// PebbleItem 表示单个数据项
type PebbleItem struct {
	Key       string `json:"key"`
	Value     []byte `json:"value"`
	Timestamp int64  `json:"timestamp"` // 使用int64避免时间戳溢出
}

func (p *Pebbledb) GetStorageData(storageType c_base.StorageType, id string, pointKey []string, startTime, endTime *int, step int) (*c_chart.ChartData, error) {
	switch storageType {
	case c_base.StorageTypeDevice:
		// 键名：device/{deviceId}/{timestamp}
		data, err := GetChartData(p.DeviceDb, fmt.Sprintf("%s%s%s", DevicePrefix, KeySeparator, id), pointKey, startTime, endTime, step, MetricsTag)
		if err != nil {
			return nil, err
		}

		return data, nil
	case c_base.StorageTypeProtocol:
		// 键名：protocol/{protocolId}/{timestamp}
		data, err := GetChartData(p.ProtocolDb, fmt.Sprintf("%s%s%s", ProtocolPrefix, KeySeparator, id), pointKey, startTime, endTime, step, MetricsTag)
		log.Println("data", data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case c_base.StorageTypeSystem:
		// 键名：system/{measurement}/{timestamp}
		data, err := GetChartData(p.SystemDb, fmt.Sprintf("%s%s%s", SystemPrefix, KeySeparator, id), pointKey, startTime, endTime, step, MetricsTag)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, nil
	}
}

// GetPagesByTimeRange 根据时间范围查询数据（key格式：prefix/timestamp）
// prefix: 键前缀
// startTime: 开始时间戳，nil表示左边界全开
// endTime: 结束时间戳，nil表示右边界全开
// page: 页码（从1开始）
// pageSize: 每页大小
// sortOrder: 排序方式 ("asc" 或 "desc")
// step: 时间间隔（毫秒），如60000表示按分钟查询，<=0表示不按间隔过滤
// 返回分页结果和总数
func GetPagesByTimeRange(p *pebble.DB, prefix string, startTime, endTime *int, page, pageSize int, sortOrder string, step int) ([]PebbleItem, int, error) {
	// 验证参数
	needPaging := true
	if page <= 0 && pageSize <= 0 {
		// 如果都没传或都为0，则不分页，返回所有数据
		needPaging = false
	} else {
		// 如果传了分页参数，则设置默认值
		if page < c_chart.DefaultPage {
			page = c_chart.DefaultPage
		}
		if pageSize < c_chart.DefaultPage {
			pageSize = c_chart.DefaultPageSize
		}
	}

	if sortOrder != SortOrderAsc && sortOrder != SortOrderDesc {
		sortOrder = SortOrderAsc // 默认升序
	}

	// 构造查询范围的边界
	var lowerBound, upperBound []byte

	if prefix == EmptyString {
		// prefix为空时，查询所有键
		if startTime != nil && *startTime != ZeroTimestamp {
			startKey := fmt.Sprintf("%d", *startTime)
			lowerBound = []byte(startKey)
		} else {
			lowerBound = nil // 不设置下界，从头开始
		}

		if endTime != nil && *endTime != ZeroTimestamp {
			endKey := fmt.Sprintf("%d", *endTime)
			upperBound = []byte(endKey + BoundarySuffix) // 添加后缀确保包含endKey
		} else {
			upperBound = nil // 不设置上界，查询到末尾
		}
	} else {
		// 有prefix的情况
		if startTime != nil && *startTime != ZeroTimestamp {
			startKey := fmt.Sprintf("%s%s%d", prefix, KeySeparator, *startTime)
			lowerBound = []byte(startKey)
		} else {
			// 左边界全开，从prefix开始
			lowerBound = []byte(prefix + KeySeparator)
		}

		if endTime != nil && *endTime != ZeroTimestamp {
			endKey := fmt.Sprintf("%s%s%d", prefix, KeySeparator, *endTime)
			upperBound = []byte(endKey + BoundarySuffix) // 添加后缀确保包含endKey
		} else {
			// 右边界全开，不设置upperBound，通过前缀的下一个字典序字符串作为上界
			prefixBytes := []byte(prefix)
			upperBound = make([]byte, len(prefixBytes))
			copy(upperBound, prefixBytes)
			// 将最后一个字符+1，创建前缀的上界
			upperBound[len(upperBound)-1]++
		}
	}

	iter, err := p.NewIter(&pebble.IterOptions{
		LowerBound: lowerBound,
		UpperBound: upperBound,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create iterator: %w", err)
	}
	defer iter.Close()

	// 收集所有匹配的数据
	var allData []PebbleItem

	for iter.First(); iter.Valid(); iter.Next() {
		keyStr := string(iter.Key())
		var timestampStr string
		var timestamp int
		var err error

		if prefix == EmptyString {
			// prefix为空时，尝试解析整个key作为时间戳，或者key中包含时间戳
			// 如果key包含"/"，取最后一部分作为时间戳
			if strings.Contains(keyStr, KeySeparator) {
				parts := strings.Split(keyStr, KeySeparator)
				timestampStr = parts[len(parts)-1]
			} else {
				// 整个key作为时间戳
				timestampStr = keyStr
			}

			timestamp, err = strconv.Atoi(timestampStr)
			if err != nil {
				continue // 跳过无效的时间戳
			}
		} else {
			// 检查key是否符合预期格式
			if !strings.HasPrefix(keyStr, prefix+KeySeparator) {
				continue
			}

			// 提取时间戳并验证范围
			parts := strings.Split(keyStr, KeySeparator)
			if len(parts) < 2 {
				continue
			}

			timestampStr = parts[len(parts)-1]
			timestamp, err = strconv.Atoi(timestampStr)
			if err != nil {
				continue // 跳过无效的时间戳
			}
		}

		// 检查时间戳是否在范围内
		inRange := true
		if startTime != nil && *startTime != ZeroTimestamp && timestamp < *startTime {
			inRange = false
		}
		if endTime != nil && *endTime != ZeroTimestamp && timestamp > *endTime {
			inRange = false
		}

		// 步长过滤改为在排序之后进行，避免与起始时间对齐导致的误筛

		if inRange {
			key := make([]byte, len(iter.Key()))
			copy(key, iter.Key())

			value := make([]byte, len(iter.Value()))
			copy(value, iter.Value())

			allData = append(allData, PebbleItem{
				Key:       string(key),
				Value:     value,
				Timestamp: int64(timestamp),
			})
		}
	}

	if err := iter.Error(); err != nil {
		return nil, 0, fmt.Errorf("iterator error: %w", err)
	}

	// 排序数据
	if sortOrder == SortOrderAsc {
		// 升序排列（时间戳从小到大）
		sort.Slice(allData, func(i, j int) bool {
			return allData[i].Timestamp < allData[j].Timestamp
		})
	} else {
		// 降序排列（时间戳从大到小）
		sort.Slice(allData, func(i, j int) bool {
			return allData[i].Timestamp > allData[j].Timestamp
		})
	}

	// 基于第一条数据的时间戳进行步长筛选
	if step > 0 && len(allData) > 0 {
		var filtered []PebbleItem
		lastTs := allData[0].Timestamp
		filtered = append(filtered, allData[0])
		step64 := int64(step)
		if sortOrder == SortOrderAsc {
			for i := 1; i < len(allData); i++ {
				if allData[i].Timestamp >= lastTs+step64 {
					filtered = append(filtered, allData[i])
					lastTs = allData[i].Timestamp
				}
			}
		} else { // 降序
			for i := 1; i < len(allData); i++ {
				if allData[i].Timestamp <= lastTs-step64 {
					filtered = append(filtered, allData[i])
					lastTs = allData[i].Timestamp
				}
			}
		}
		allData = filtered
	}

	// 计算分页
	total := len(allData)

	if !needPaging {
		// 不分页，返回所有数据（已应用步长筛选）
		return allData, total, nil
	}

	// 进行分页
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if startIndex >= total {
		// 页码超出范围，返回空结果
		return []PebbleItem{}, total, nil
	}

	if endIndex > total {
		endIndex = total
	}

	// 构建分页结果
	result := allData[startIndex:endIndex]

	return result, total, nil
}

// GetChartData根据时间范围查询数据（key格式：prefix/timestamp）
// prefix: 键前缀
// points: 点位列表
// startTime: 开始时间戳，nil表示左边界全开
// endTime: 结束时间戳，nil表示右边界全开
// step: 时间间隔（毫秒），如60000表示按分钟查询，<=0表示不按间隔过滤
// tag: 标签， 获取到数据后，根据tag过滤数据
// 返回分页结果和总数
func GetChartData(p *pebble.DB, prefix string, points []string, startTime, endTime *int, step int, tag string) (*c_chart.ChartData, error) {

	// 构造查询范围的边界
	var lowerBound, upperBound []byte

	if prefix == EmptyString {
		// prefix为空时，查询所有键
		if startTime != nil && *startTime != ZeroTimestamp {
			startKey := fmt.Sprintf("%d", *startTime)
			lowerBound = []byte(startKey)
		} else {
			lowerBound = nil // 不设置下界，从头开始
		}

		if endTime != nil && *endTime != ZeroTimestamp {
			endKey := fmt.Sprintf("%d", *endTime)
			upperBound = []byte(endKey + BoundarySuffix) // 添加后缀确保包含endKey
		} else {
			upperBound = nil // 不设置上界，查询到末尾
		}
	} else {
		// 有prefix的情况
		if startTime != nil && *startTime != ZeroTimestamp {
			startKey := fmt.Sprintf("%s%s%d", prefix, KeySeparator, *startTime)
			lowerBound = []byte(startKey)
		} else {
			// 左边界全开，从prefix开始
			lowerBound = []byte(prefix + KeySeparator)
		}

		if endTime != nil && *endTime != ZeroTimestamp {
			endKey := fmt.Sprintf("%s%s%d", prefix, KeySeparator, *endTime)
			upperBound = []byte(endKey + BoundarySuffix) // 添加后缀确保包含endKey
		} else {
			// 右边界全开，不设置upperBound，通过前缀的下一个字典序字符串作为上界
			prefixBytes := []byte(prefix)
			upperBound = make([]byte, len(prefixBytes))
			copy(upperBound, prefixBytes)
			// 将最后一个字符+1，创建前缀的上界
			upperBound[len(upperBound)-1]++
		}
	}

	iter, err := p.NewIter(&pebble.IterOptions{
		LowerBound: lowerBound,
		UpperBound: upperBound,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create iterator: %w", err)
	}
	defer iter.Close()

	// 收集所有匹配的数据
	chartData := c_chart.NewChartData(len(points))

	// 创建数据系列映射，用于数据收集过程
	seriesMap := make(map[string]*c_chart.Series)
	for _, point := range points {
		series := c_chart.NewSeries(point, c_chart.ChartTypeLine, 0) // 容量设为0，让其自动扩容
		seriesMap[point] = series
	}

	for iter.First(); iter.Valid(); iter.Next() {
		keyStr := string(iter.Key())
		var timestampStr string
		var timestamp int
		var err error

		if prefix == EmptyString {
			// prefix为空时，尝试解析整个key作为时间戳，或者key中包含时间戳
			// 如果key包含"/"，取最后一部分作为时间戳
			if strings.Contains(keyStr, KeySeparator) {
				parts := strings.Split(keyStr, KeySeparator)
				timestampStr = parts[len(parts)-1]
			} else {
				// 整个key作为时间戳
				timestampStr = keyStr
			}

			timestamp, err = strconv.Atoi(timestampStr)
			if err != nil {
				continue // 跳过无效的时间戳
			}
		} else {
			// 检查key是否符合预期格式
			if !strings.HasPrefix(keyStr, prefix+KeySeparator) {
				continue
			}

			// 提取时间戳并验证范围
			parts := strings.Split(keyStr, KeySeparator)
			if len(parts) < 2 {
				continue
			}

			timestampStr = parts[len(parts)-1]
			timestamp, err = strconv.Atoi(timestampStr)
			if err != nil {
				continue // 跳过无效的时间戳
			}
		}

		// 检查时间戳是否在范围内
		inRange := true
		if startTime != nil && *startTime != ZeroTimestamp && timestamp < *startTime {
			inRange = false
		}
		if endTime != nil && *endTime != ZeroTimestamp && timestamp > *endTime {
			inRange = false
		}

		// 检查时间间隔步长
		if inRange && step > 0 {
			// 如果设置了起始时间，从起始时间开始按step间隔过滤
			if startTime != nil && *startTime != ZeroTimestamp {
				if (timestamp-*startTime)%step != 0 {
					inRange = false
				}
			} else {
				// 如果没有设置起始时间，则按时间戳对step取模
				if timestamp%step != 0 {
					inRange = false
				}
			}
		}

		if inRange {
			chartData.AddTimestamp(int64(timestamp))
			value, err := iter.ValueAndErr()
			if err != nil {
				continue
			}
			var data map[string]any
			err = json.Unmarshal(value, &data)
			if err != nil {
				log.Println("unmarshal value error", err)
				continue
			}
			if data[tag] == nil {
				continue
			}

			pointData := data[tag].(map[string]any)
			for _, point := range points {
				if pointData[point] != nil {
					seriesMap[point].AppendData(fmt.Sprintf("%v", pointData[point]))
				} else {
					seriesMap[point].AppendData("")
				}
			}

		}
	}

	if err := iter.Error(); err != nil {
		return nil, fmt.Errorf("iterator error: %w", err)
	}

	// 将收集到的数据系列添加到图表数据中
	for _, point := range points {
		chartData.AddSeries(*seriesMap[point])
	}

	return chartData, nil
}
