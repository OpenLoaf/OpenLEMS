package s_db_basic

import "common/c_enum"

// 系统设置定义变量
var (
	// 根设备ID设置定义
	SystemSettingActiveDeviceRootId = &SSystemSettingDefine{
		Id:           "active_device_root_id",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "0",
		IsPublic:     true,
		Remark:       "根设备ID",
	}

	// 激活的策略ID设置定义
	SystemSettingActivePolicyId = &SSystemSettingDefine{
		Id:           "active_policy_id",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "",
		IsPublic:     true,
		Remark:       "激活的策略ID",
	}

	// 自动化任务轮询周期设置定义
	SystemSettingAutomationInternalMilliseconds = &SSystemSettingDefine{
		Id:           "automation_internal_milliseconds",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "1000",
		IsPublic:     true,
		Remark:       "自动化任务轮询周期（毫秒）",
	}

	// 设备数据保留天数设置定义
	SystemSettingDeviceRetentionDays = &SSystemSettingDefine{
		Id:           "DeviceRetentionDays",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "100",
		IsPublic:     true,
		Remark:       "设备数据保留天数",
	}

	// 系统数据保留天数设置定义
	SystemSettingSystemRetentionDays = &SSystemSettingDefine{
		Id:           "SystemRetentionDays",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "7",
		IsPublic:     true,
		Remark:       "系统数据保留天数",
	}

	// 日志数据保留天数设置定义
	SystemSettingLogRetentionDays = &SSystemSettingDefine{
		Id:           "LogRetentionDays",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "30",
		IsPublic:     true,
		Remark:       "日志数据保留天数",
	}

	// 系统调试日志开关设置定义
	SystemSettingSystemEnableDebugLog = &SSystemSettingDefine{
		Id:           "SystemEnableDebugLog",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "false",
		IsPublic:     true,
		Remark:       "启用系统调试日志",
	}

	// 软件许可证号设置定义
	SystemSettingLicenseKey = &SSystemSettingDefine{
		Id:           "license_key",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "",
		IsPublic:     false,
		Remark:       "软件许可证号",
	}

	// 普通用户密码设置定义
	SystemSettingUserPassword = &SSystemSettingDefine{
		Id:           "user_password",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "1111111",
		IsPublic:     false,
		Remark:       "普通用户密码",
	}

	// 管理员用户密码设置定义
	SystemSettingAdminPassword = &SSystemSettingDefine{
		Id:           "admin_password",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "888888",
		IsPublic:     false,
		Remark:       "管理员用户密码",
	}

	// 密码长度设置定义
	SystemSettingPasswordLength = &SSystemSettingDefine{
		Id:           "password_length",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "6",
		IsPublic:     true,
		Remark:       "密码长度",
	}
)
