package s_db_basic

const (
	SettingActiveDeviceRootIdKey = "active_device_root_id" // 根设备ID
	DefaultActiveDeviceRootId    = "0"                     // 默认根设备ID

	SettingActivePolicyIdKey = "active_policy_id" // 激活的策略ID

	SettingDeviceRetentionDays = "DeviceRetentionDays"
	DefaultDeviceRetentionDays = "100" // 默认设备保存天数

	SettingSystemRetentionDays = "SystemRetentionDays"
	DefaultSystemRetentionDays = "7" // 默认系统数据保存天数

	SettingLogRetentionDays = "LogRetentionDays"
	DefaultLogRetentionDays = "30" // 默认
)
