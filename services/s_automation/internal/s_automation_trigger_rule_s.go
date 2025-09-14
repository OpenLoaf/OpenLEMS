package internal

// SAutomationTriggerCondition 自动化触发条件结构体
type SAutomationTriggerCondition struct {
	DeviceId string `json:"deviceId"` // 设备ID
	Rule     string `json:"rule"`     // 规则表达式，如 "P>30", "Ia<100"
}

// SAutomationTriggerConfig 自动化触发配置结构体
type SAutomationTriggerConfig struct {
	AnyMatch    []*SAutomationTriggerCondition `json:"anyMatch"`    // 任意匹配条件（OR 逻辑）
	SubMatch    []*SAutomationTriggerCondition `json:"subMatch"`    // 子匹配条件
	SubMatchAll bool                           `json:"subMatchAll"` // 子匹配是否全部满足
}
