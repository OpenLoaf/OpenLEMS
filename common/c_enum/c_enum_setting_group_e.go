package c_enum

type ESettingGroup = string

const (
	ESettingGroupSystem ESettingGroup = "system" // 系统设置
	ESettingGroupPolicy ESettingGroup = "policy" // 策略设置
)
