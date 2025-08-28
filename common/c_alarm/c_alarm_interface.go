package c_alarm

import "context"

// IAlarmManager 告警管理器接口
type IAlarmManager interface {
	// CreateAlarmHistory 创建告警历史记录
	// @param ctx 上下文
	// @param deviceId 设备ID，用于标识告警来源设备
	// @param point 告警点位名称，如"temperature"、"voltage"等
	// @param level 告警等级，如"LOW"、"MEDIUM"、"HIGH"、"CRITICAL"
	// @param title 告警标题，简要描述告警内容
	// @param detail 告警详情，详细描述告警的具体信息
	// @return error 返回错误信息，成功时返回nil
	CreateAlarmHistory(ctx context.Context, deviceId, point, level, title, detail string) error

	// IsAlarmIgnored 检查告警是否被忽略
	// @param ctx 上下文
	// @param deviceId 设备ID，用于标识告警来源设备
	// @param point 告警点位名称，如"temperature"、"voltage"等
	// @return bool 返回true表示告警被忽略，false表示告警未被忽略
	// @return error 返回错误信息，成功时返回nil
	IsAlarmIgnored(ctx context.Context, deviceId, point string) (bool, error)
}
