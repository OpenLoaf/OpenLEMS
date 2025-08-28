package c_alarm

import "context"

// IAlarmManager 告警管理器接口
type IAlarmManager interface {
	// CreateAlarmHistory 创建告警历史记录
	CreateAlarmHistory(ctx context.Context, deviceId, point, level, title, detail string) error

	// IsAlarmIgnored 检查告警是否被忽略
	IsAlarmIgnored(ctx context.Context, deviceId, point string) (bool, error)
}
