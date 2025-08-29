package s_db_basic

import (
	"common/c_base"
	"context"
	"s_db/s_db_model"
	"time"
)

// IDeviceService 设备服务
type IDeviceService interface {
	GetEnableDeviceConfigsWithRecursion(ctx context.Context, parentId string) ([]*c_base.SDeviceConfig, error) // 获取所有设备列表
	GetRootDeviceId() string                                                                                   // 获取根设备ID
	GetAllDevices(ctx context.Context) ([]*s_db_model.SDeviceModel, error)                                     // 获取所有设备
	UpdateDevice(ctx context.Context, deviceId string, data map[string]interface{}) error
}

type ISettingService interface {
	GetAllSettings(ctx context.Context) ([]*s_db_model.SSettingModel, error)
	GetSettingValueByKey(ctx context.Context, name string, defaultValue ...string) string // 获取设置，如果获取不到，就设置为默认值
	SetSettingValueByName(ctx context.Context, name string, value string) error
}

type IProtocolService interface {
	GetProtocolList(ctx context.Context, type_ string) ([]*s_db_model.SProtocolModel, error)
	UpdateProtocol(ctx context.Context, protocolId string, data map[string]interface{}) error
	CreateProtocol(ctx context.Context, data map[string]interface{}) (string, error)
	DeleteProtocol(ctx context.Context, protocolId string) error
	GetAllProtocolConfigs(ctx context.Context) ([]*c_base.SProtocolConfig, error) // 获取协议列表
}

type IAlarmService interface {
	// 告警历史相关方法
	CreateAlarmHistory(ctx context.Context, deviceId, sourceDeviceId, point, level, title, detail string, triggerAt *time.Time) error
	GetAlarmHistoryByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SAlarmHistoryModel, error)
	GetAlarmHistoryByDeviceIdAndPoint(ctx context.Context, deviceId, point string) ([]*s_db_model.SAlarmHistoryModel, error)
	DeleteAlarmHistoryByDeviceId(ctx context.Context, deviceId string) error
	GetAllAlarmHistory(ctx context.Context) ([]*s_db_model.SAlarmHistoryModel, error)
	GetAlarmHistoryPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SAlarmHistoryModel, int, error)
	ClearAllAlarmHistory(ctx context.Context) error
	GetAlarmHistoryCount(ctx context.Context) int

	// 告警忽略相关方法
	CreateAlarmIgnore(ctx context.Context, deviceId, sourceDeviceId, point, pointName string) error
	GetAlarmIgnoreByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SAlarmIgnoreModel, error)
	IsAlarmIgnored(ctx context.Context, deviceId, sourceDeviceId, point string) (bool, error)
	DeleteAlarmIgnoreByDeviceId(ctx context.Context, deviceId string) error
	DeleteAlarmIgnoreByDeviceIdAndPoint(ctx context.Context, deviceId, point string) error
	GetAllAlarmIgnore(ctx context.Context) ([]*s_db_model.SAlarmIgnoreModel, error)
	GetAlarmIgnorePage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SAlarmIgnoreModel, int, error)
	GetAlarmIgnoreCount(ctx context.Context) int
}

type ILogService interface {
	CreateLog(ctx context.Context, logType, deviceId, level, content string) error
	GetLogByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SLogModel, error)
	GetLogByType(ctx context.Context, logType string) ([]*s_db_model.SLogModel, error)
	GetLogByLevel(ctx context.Context, level string) ([]*s_db_model.SLogModel, error)
	GetLogByDeviceIdAndType(ctx context.Context, deviceId, logType string) ([]*s_db_model.SLogModel, error)
	DeleteLogByDeviceId(ctx context.Context, deviceId string) error
	DeleteLogByType(ctx context.Context, logType string) error
	GetAllLog(ctx context.Context) ([]*s_db_model.SLogModel, error)
	GetLogPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SLogModel, int, error)
	ClearAllLog(ctx context.Context) error
	GetLogCount(ctx context.Context) (int, error)
}
