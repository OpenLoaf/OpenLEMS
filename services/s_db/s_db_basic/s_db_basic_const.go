package s_db_basic

import (
	"common/c_enum"
	"math/rand"
	"time"
)

// generateRandomString 生成指定长度的随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

// 系统设置定义变量
var (
	// 根设备ID设置定义
	SystemSettingActiveDeviceRootId = &SSystemSettingDefine{
		Id:           "active_device_root_id",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "0",
		IsPublic:     true,
		Remark:       "根设备ID",
		FieldType:    c_enum.ESettingFieldTypeText,
	}

	// 激活的策略ID设置定义
	SystemSettingActivePolicyId = &SSystemSettingDefine{
		Id:           "active_policy_id",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "",
		IsPublic:     true,
		Remark:       "激活的策略ID",
		FieldType:    c_enum.ESettingFieldTypeText,
	}

	// 自动化任务轮询周期设置定义
	SystemSettingAutomationInternalMilliseconds = &SSystemSettingDefine{
		Id:           "automation_internal_milliseconds",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "1000",
		IsPublic:     true,
		Remark:       "自动化任务轮询周期（毫秒）",
		FieldType:    c_enum.ESettingFieldTypeNumber,
	}

	// 设备数据保留天数设置定义
	SystemSettingDeviceRetentionDays = &SSystemSettingDefine{
		Id:           "DeviceRetentionDays",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "100",
		IsPublic:     true,
		Remark:       "设备数据保留天数",
		FieldType:    c_enum.ESettingFieldTypeNumber,
	}

	// 系统数据保留天数设置定义
	SystemSettingSystemRetentionDays = &SSystemSettingDefine{
		Id:           "SystemRetentionDays",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "7",
		IsPublic:     true,
		Remark:       "系统数据保留天数",
		FieldType:    c_enum.ESettingFieldTypeNumber,
	}

	// 日志数据保留天数设置定义
	SystemSettingLogRetentionDays = &SSystemSettingDefine{
		Id:           "LogRetentionDays",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "30",
		IsPublic:     true,
		Remark:       "日志数据保留天数",
		FieldType:    c_enum.ESettingFieldTypeNumber,
	}

	// 系统调试日志开关设置定义
	SystemSettingSystemEnableDebugLog = &SSystemSettingDefine{
		Id:           "SystemEnableDebugLog",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "false",
		IsPublic:     true,
		Remark:       "启用系统调试日志",
		FieldType:    c_enum.ESettingFieldTypeBoolean,
	}

	// 软件许可证号设置定义
	SystemSettingLicenseKey = &SSystemSettingDefine{
		Id:           "license_key",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "",
		IsPublic:     false,
		Remark:       "软件许可证号",
		FieldType:    c_enum.ESettingFieldTypeText,
	}

	// 普通用户密码设置定义
	SystemSettingUserPassword = &SSystemSettingDefine{
		Id:           "user_password",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "1111111",
		IsPublic:     false,
		Remark:       "普通用户密码",
		FieldType:    c_enum.ESettingFieldTypeText,
	}

	// 管理员用户密码设置定义
	SystemSettingAdminPassword = &SSystemSettingDefine{
		Id:           "admin_password",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "888888",
		IsPublic:     false,
		Remark:       "管理员用户密码",
		FieldType:    c_enum.ESettingFieldTypeText,
	}

	// 密码长度设置定义
	SystemSettingPasswordLength = &SSystemSettingDefine{
		Id:           "password_length",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "6",
		IsPublic:     true,
		Remark:       "密码长度",
		FieldType:    c_enum.ESettingFieldTypeNumber,
	}

	// MQTT配置设置定义
	SystemSettingMqttConfigList = &SSystemSettingDefine{
		Id:           "mqtt_config_list",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "[]",
		IsPublic:     false,
		Remark:       "MQTT配置列表",
		FieldType:    c_enum.ESettingFieldTypeJsonArray,
	}

	// Modbus配置设置定义
	SystemSettingModbusConfig = &SSystemSettingDefine{
		Id:           "modbus_config",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "{}",
		IsPublic:     false,
		Remark:       "Modbus配置",
		FieldType:    c_enum.ESettingFieldTypeJson,
	}

	// 管理员会话超时时间（小时）
	SystemSettingSessionAdminTimeout = &SSystemSettingDefine{
		Id:           "session_admin_timeout",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "2",
		IsPublic:     false,
		Remark:       "管理员Session过期时间（小时）",
		FieldType:    c_enum.ESettingFieldTypeNumber,
	}

	// 普通用户会话超时时间（小时）
	SystemSettingSessionUserTimeout = &SSystemSettingDefine{
		Id:           "session_user_timeout",
		Group:        c_enum.ESettingGroupSystem,
		DefaultValue: "2",
		IsPublic:     false,
		Remark:       "普通用户Session过期时间（小时）",
		FieldType:    c_enum.ESettingFieldTypeNumber,
	}
)
