package internal

import (
	"context"
	"encoding/json"
	"time"

	"common/c_log"
)

// SStandardFormatter Standard格式数据格式化器
type SStandardFormatter struct{}

// StandardData Standard格式的数据结构
type StandardData struct {
	SN   string         `json:"sn"`   // 设备序列号（使用设备ID）
	Time int64          `json:"time"` // 时间戳（毫秒）
	Data map[string]any `json:"data"` // 设备点位数据
}

// Format 格式化设备数据为标准格式
// 输出格式：{"sn":"xxxx","time":时间戳,"data":{"pointA":数值,"pointB":数值}}
func (f *SStandardFormatter) Format(ctx context.Context, deviceId string, deviceData map[string]any, systemNumber string) ([]byte, error) {
	// 创建标准格式数据
	data := StandardData{
		SN:   deviceId,               // 使用设备ID作为序列号
		Time: time.Now().UnixMilli(), // 当前时间戳（毫秒）
		Data: deviceData,             // 设备点位数据
	}

	// 序列化为JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		c_log.Errorf(ctx, "序列化设备数据失败: 设备ID=%s, 错误=%v", deviceId, err)
		return nil, err
	}

	c_log.Debugf(ctx, "格式化设备数据成功: 设备ID=%s, 数据长度=%d", deviceId, len(jsonData))
	return jsonData, nil
}

// GetTopicTemplate 返回topic模板
// Topic: lems/{system_number}/info
func (f *SStandardFormatter) GetTopicTemplate() string {
	return "lems/{system_number}/info"
}
