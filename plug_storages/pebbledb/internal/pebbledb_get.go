package internal

import (
	"common/c_base"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/cockroachdb/pebble"
)

type XAxis struct {
	Type string   `json:"type"`
	Data []string `json:"data"`
}

type SeriesItem struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	Data []string `json:"data"`
}

type ChartData struct {
	XAxis  XAxis        `json:"xAxis"`
	Series []SeriesItem `json:"series"`
}

// PebbleItem 表示单个数据项
type PebbleItem struct {
	Key       string `json:"key"`
	Value     []byte `json:"value"`
	Timestamp int    `json:"timestamp"`
}

func (p *Pebbledb) GetStorageData(storageType c_base.StorageType, id string, pointKey []string, startTime, endTime *int, page, pageSize int, sortOrder string, step int) (map[string]any, error) {
	switch storageType {
	case c_base.StorageTypeDevice:
		// 键名：device/{deviceId}/{timestamp}
		data, err := GetChartData(p.db.DeviceDb, fmt.Sprintf("device/%s", id), pointKey, startTime, endTime, step, "metrics")
		if err != nil {
			return nil, err
		}

		return map[string]any{
			"data": data,
		}, nil
	case c_base.StorageTypeProtocol:
		// 键名：protocol/{protocolId}/{timestamp}
		data, err := GetChartData(p.db.ProtocolDb, fmt.Sprintf("protocol/%s", id), pointKey, startTime, endTime, step, "metrics")
		log.Println("data", data)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"data": data,
		}, nil
	case c_base.StorageTypeSystem:
		// 键名：system/{measurement}/{timestamp}
		data, err := GetChartData(p.db.SystemDb, fmt.Sprintf("system/%s", id), pointKey, startTime, endTime, step, "metrics")
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"data": data,
		}, nil
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
		if page < 1 {
			page = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc" // 默认升序
	}

	// 构造查询范围的边界
	var lowerBound, upperBound []byte

	if prefix == "" {
		// prefix为空时，查询所有键
		if startTime != nil && *startTime != 0 {
			startKey := fmt.Sprintf("%d", *startTime)
			lowerBound = []byte(startKey)
		} else {
			lowerBound = nil // 不设置下界，从头开始
		}

		if endTime != nil && *endTime != 0 {
			endKey := fmt.Sprintf("%d", *endTime)
			upperBound = []byte(endKey + "z") // 添加后缀确保包含endKey
		} else {
			upperBound = nil // 不设置上界，查询到末尾
		}
	} else {
		// 有prefix的情况
		if startTime != nil && *startTime != 0 {
			startKey := fmt.Sprintf("%s/%d", prefix, *startTime)
			lowerBound = []byte(startKey)
		} else {
			// 左边界全开，从prefix开始
			lowerBound = []byte(prefix + "/")
		}

		if endTime != nil && *endTime != 0 {
			endKey := fmt.Sprintf("%s/%d", prefix, *endTime)
			upperBound = []byte(endKey + "z") // 添加后缀确保包含endKey
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

		if prefix == "" {
			// prefix为空时，尝试解析整个key作为时间戳，或者key中包含时间戳
			// 如果key包含"/"，取最后一部分作为时间戳
			if strings.Contains(keyStr, "/") {
				parts := strings.Split(keyStr, "/")
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
			if !strings.HasPrefix(keyStr, prefix+"/") {
				continue
			}

			// 提取时间戳并验证范围
			parts := strings.Split(keyStr, "/")
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
		if startTime != nil && *startTime != 0 && timestamp < *startTime {
			inRange = false
		}
		if endTime != nil && *endTime != 0 && timestamp > *endTime {
			inRange = false
		}

		// 检查时间间隔步长
		if inRange && step > 0 {
			// 如果设置了起始时间，从起始时间开始按step间隔过滤
			if startTime != nil && *startTime != 0 {
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
			key := make([]byte, len(iter.Key()))
			copy(key, iter.Key())

			value := make([]byte, len(iter.Value()))
			copy(value, iter.Value())

			allData = append(allData, PebbleItem{
				Key:       string(key),
				Value:     value,
				Timestamp: timestamp,
			})
		}
	}

	if err := iter.Error(); err != nil {
		return nil, 0, fmt.Errorf("iterator error: %w", err)
	}

	// 排序数据
	if sortOrder == "asc" {
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

	// 计算分页
	total := len(allData)

	if !needPaging {
		// 不分页，返回所有数据
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
func GetChartData(p *pebble.DB, prefix string, points []string, startTime, endTime *int, step int, tag string) (*ChartData, error) {

	// 构造查询范围的边界
	var lowerBound, upperBound []byte

	if prefix == "" {
		// prefix为空时，查询所有键
		if startTime != nil && *startTime != 0 {
			startKey := fmt.Sprintf("%d", *startTime)
			lowerBound = []byte(startKey)
		} else {
			lowerBound = nil // 不设置下界，从头开始
		}

		if endTime != nil && *endTime != 0 {
			endKey := fmt.Sprintf("%d", *endTime)
			upperBound = []byte(endKey + "z") // 添加后缀确保包含endKey
		} else {
			upperBound = nil // 不设置上界，查询到末尾
		}
	} else {
		// 有prefix的情况
		if startTime != nil && *startTime != 0 {
			startKey := fmt.Sprintf("%s/%d", prefix, *startTime)
			lowerBound = []byte(startKey)
		} else {
			// 左边界全开，从prefix开始
			lowerBound = []byte(prefix + "/")
		}

		if endTime != nil && *endTime != 0 {
			endKey := fmt.Sprintf("%s/%d", prefix, *endTime)
			upperBound = []byte(endKey + "z") // 添加后缀确保包含endKey
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
	var chartData ChartData
	xAxis := XAxis{
		Type: "category",
		Data: make([]string, 0),
	}
	series := make([]SeriesItem, len(points))
	for i, point := range points {
		series[i].Name = point
		series[i].Type = "line"
		series[i].Data = make([]string, 0)
	}

	for iter.First(); iter.Valid(); iter.Next() {
		keyStr := string(iter.Key())
		var timestampStr string
		var timestamp int
		var err error

		if prefix == "" {
			// prefix为空时，尝试解析整个key作为时间戳，或者key中包含时间戳
			// 如果key包含"/"，取最后一部分作为时间戳
			if strings.Contains(keyStr, "/") {
				parts := strings.Split(keyStr, "/")
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
			if !strings.HasPrefix(keyStr, prefix+"/") {
				continue
			}

			// 提取时间戳并验证范围
			parts := strings.Split(keyStr, "/")
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
		if startTime != nil && *startTime != 0 && timestamp < *startTime {
			inRange = false
		}
		if endTime != nil && *endTime != 0 && timestamp > *endTime {
			inRange = false
		}

		// 检查时间间隔步长
		if inRange && step > 0 {
			// 如果设置了起始时间，从起始时间开始按step间隔过滤
			if startTime != nil && *startTime != 0 {
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
			xAxis.Data = append(xAxis.Data, fmt.Sprintf("%d", timestamp))
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
			for i, point := range points {
				if pointData[point] != nil {
					series[i].Data = append(series[i].Data, fmt.Sprintf("%v", pointData[point]))
				} else {
					series[i].Data = append(series[i].Data, "")
				}
			}

		}
	}

	if err := iter.Error(); err != nil {
		return nil, fmt.Errorf("iterator error: %w", err)
	}
	chartData.XAxis = xAxis
	chartData.Series = series
	return &chartData, nil
}
