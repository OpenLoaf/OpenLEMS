package c_alarm

import (
	"context"
	"time"
)

// IAlarmManager 告警管理器接口
type IAlarmManager interface {
	// CreateAlarmHistory 创建告警历史记录
	CreateAlarmHistory(ctx context.Context, deviceId, sourceDeviceId, point, level, title, detail string, triggerAt *time.Time) error

	// IsAlarmIgnored 检查告警是否被忽略
	IsAlarmIgnored(ctx context.Context, deviceId, sourceDeviceId, point string) (bool, error)
}
