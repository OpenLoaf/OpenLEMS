package internal

import "context"

// IDataFormatter 数据格式化器接口，支持扩展多种格式
type IDataFormatter interface {
	// Format 格式化设备数据
	// deviceId: 设备ID
	// deviceData: 设备点位数据
	// systemNumber: 系统序列号
	// 返回: 格式化后的JSON字节数组
	Format(ctx context.Context, deviceId string, deviceData map[string]any, systemNumber string) ([]byte, error)

	// GetTopicTemplate 返回topic模板，如 "lems/{system_number}/info"
	GetTopicTemplate() string
}
