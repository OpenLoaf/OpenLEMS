package s_db_basic

const (
	SettingActiveDeviceRootIdKey = "active_device_root_id" // 根设备ID
	DefaultActiveDeviceRootId    = "0"                     // 默认根设备ID

	SettingActivePolicyIdKey = "active_policy_id" // 激活的策略ID

	SettingAutomationInternalMillisecondsKey = "automation_internal_milliseconds"
	DefaultAutomationInternalMilliseconds    = "1000"

	SettingDeviceRetentionDays = "DeviceRetentionDays"
	DefaultDeviceRetentionDays = "100" // 默认设备保存天数

	SettingSystemRetentionDays = "SystemRetentionDays"
	DefaultSystemRetentionDays = "7" // 默认系统数据保存天数

	SettingLogRetentionDays = "LogRetentionDays"
	DefaultLogRetentionDays = "30" // 默认

	SettingSystemEnableDebugLog = "SystemEnableDebugLog" // 启用系统调试日志
	DefaultSystemEnableDebugLog = "false"                // 默认不启用
)
