package s_db_basic

import (
	"common/c_base"
	"common/c_enum"
	"context"
	"s_db/s_db_model"
	"time"
)

// IDeviceService 设备服务
type IDeviceService interface {
	GetDeviceConfigsWithRecursion(ctx context.Context, parentId string) ([]*c_base.SDeviceConfig, error) // 获取所有设备列表（包括enabled=false）
	GetAllDevices(ctx context.Context) ([]*s_db_model.SDeviceModel, error)                               // 获取所有设备
	GetDeviceById(ctx context.Context, id string) (*s_db_model.SDeviceModel, error)                      // 根据ID获取设备
	UpdateDevice(ctx context.Context, deviceId string, data map[string]interface{}) error
}

type ISettingService interface {
	GetAllSettings(ctx context.Context) ([]*s_db_model.SSettingModel, error)
	GetAllSettingsByGroup(ctx context.Context, group string) ([]*s_db_model.SSettingModel, error)                     // 根据分组获取所有设置
	GetSettingById(ctx context.Context, id string) (*s_db_model.SSettingModel, error)                                 // 根据ID获取设置详情
	GetSettingValueById(ctx context.Context, id string) string                                                        // 获取设置，如果获取不到，返回空字符串
	GetSettingValueByIdWithDefaultValue(ctx context.Context, id, group, defaultValue string, remark ...string) string // 获取设置，如果获取不到，就设置为默认值
	SetSettingValueById(ctx context.Context, id string, value string) error
	GetRootDeviceId(ctx context.Context) string                                        // 获取根设备ID
	GetRootPolicyId(ctx context.Context) string                                        // 获取激活的策略ID
	GetPublicEnabledSettings(ctx context.Context) ([]*s_db_model.SSettingModel, error) // 获取公开且启用的设置
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
	CreateAlarmHistory(ctx context.Context, deviceId, sourceDeviceId string, meta c_base.IPoint, level c_enum.EAlarmLevel, detail string, triggerAt time.Time) error
	GetAlarmHistoryByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SAlarmHistoryModel, error)
	GetAlarmHistoryByDeviceIdAndPoint(ctx context.Context, deviceId, point string) ([]*s_db_model.SAlarmHistoryModel, error)
	DeleteAlarmHistoryByDeviceId(ctx context.Context, deviceId string) error
	GetAllAlarmHistory(ctx context.Context) ([]*s_db_model.SAlarmHistoryModel, error)
	GetAlarmHistoryPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SAlarmHistoryModel, int, error)
	ClearAllAlarmHistory(ctx context.Context) error
	GetAlarmHistoryCount(ctx context.Context, deviceId string) int

	// 告警忽略相关方法
	CreateAlarmIgnore(ctx context.Context, deviceId, sourceDeviceId, point, pointName string) error
	GetAlarmIgnoreByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SAlarmIgnoreModel, error)
	IsAlarmIgnored(ctx context.Context, deviceId, sourceDeviceId, point string) (bool, error)
	DeleteAlarmIgnoreByDeviceId(ctx context.Context, deviceId string) error
	DeleteAlarmIgnoreByDeviceIdAndPoint(ctx context.Context, deviceId, point string) error
	GetAllAlarmIgnore(ctx context.Context) ([]*s_db_model.SAlarmIgnoreModel, error)
	GetAlarmIgnorePage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SAlarmIgnoreModel, int, error)
	GetAlarmIgnoreCount(ctx context.Context, deviceId string) int
}

type ILogService interface {
	CreateLog(ctx context.Context, logType, deviceId, level, content string) error
	GetLogByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SLogModel, error)
	GetLogByType(ctx context.Context, logType string) ([]*s_db_model.SLogModel, error)
	GetLogByLevel(ctx context.Context, level string) ([]*s_db_model.SLogModel, error)
	GetLogByDeviceIdAndType(ctx context.Context, deviceId, logType string) ([]*s_db_model.SLogModel, error)
	DeleteLogByDeviceId(ctx context.Context, deviceId string) error
	DeleteLogByType(ctx context.Context, logType string) error
	DeleteLogByFilters(ctx context.Context, filters map[string]interface{}) (int, error)
	GetAllLog(ctx context.Context) ([]*s_db_model.SLogModel, error)
	GetLogPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SLogModel, int, error)
	ClearAllLog(ctx context.Context) error
	GetLogCount(ctx context.Context) (int, error)
}

type IAutomationService interface {
	// 基础 CRUD 方法
	CreateAutomation(ctx context.Context, name string, startTime, endTime *time.Time, timeRangeType, timeRangeValue, triggerRule, executeRule string, executionInterval int) (int, error)
	GetAutomationById(ctx context.Context, id int) (*s_db_model.SAutomationModel, error)
	UpdateAutomation(ctx context.Context, id int, data map[string]interface{}) error
	DeleteAutomation(ctx context.Context, id int) error

	// 查询方法
	GetAllAutomations(ctx context.Context) ([]*s_db_model.SAutomationModel, error)
	GetAutomationsByTimeRangeType(ctx context.Context, timeRangeType string) ([]*s_db_model.SAutomationModel, error)
	GetEnabledAutomations(ctx context.Context) ([]*s_db_model.SAutomationModel, error)
	GetAutomationsByFilters(ctx context.Context, deviceId string, filters map[string]interface{}) ([]*s_db_model.SAutomationModel, error)
	GetAutomationPage(ctx context.Context, page, pageSize int, deviceId string, filters map[string]interface{}) ([]*s_db_model.SAutomationModel, int, error)

	// 统计方法
	ClearAllAutomations(ctx context.Context) error
	GetAutomationCount(ctx context.Context) (int, error)
}
