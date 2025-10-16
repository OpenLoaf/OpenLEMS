package internal

import (
	"fmt"
	"time"
)

// SAutomationDeviceCondition 设备触发条件结构体
type SAutomationDeviceCondition struct {
	DeviceId string `json:"deviceId"` // 设备ID
	From     string `json:"from"`     // 是从哪里去取值
	Rule     string `json:"rule"`     // 规则表达式，如 "P>30", "Ia<100"
}

// SAutomationTriggerCondition 自动化触发条件结构体
type SAutomationTriggerCondition struct {
	// 设备值触发条件
	DeviceCondition *SAutomationDeviceCondition `json:"deviceCondition,omitempty"` // 设备触发条件

	// 时间触发条件
	TimeCondition *SAutomationTimeCondition `json:"timeCondition,omitempty"` // 时间触发条件
}

// SAutomationTimeCondition 时间触发条件结构体
type SAutomationTimeCondition struct {
	// 基础时间条件
	Hour   *int `json:"hour,omitempty"`   // 小时 (0-23)，nil 表示不限制
	Minute *int `json:"minute,omitempty"` // 分钟 (0-59)，nil 表示不限制

	// 日期条件
	DayOfWeek  *int `json:"dayOfWeek,omitempty"`  // 星期几 (0-6，0=周日)，nil 表示不限制
	DayOfMonth *int `json:"dayOfMonth,omitempty"` // 月的第几天 (1-31)，nil 表示不限制

	// 月份条件
	Month *int `json:"month,omitempty"` // 月份 (1-12)，nil 表示不限制

	// 时间范围条件
	StartTime string `json:"startTime,omitempty"` // 开始时间 (HH:MM 格式)
	EndTime   string `json:"endTime,omitempty"`   // 结束时间 (HH:MM 格式)
}

// SAutomationTriggerConfig 自动化触发配置结构体
type SAutomationTriggerConfig struct {
	AnyMatch          []*SAutomationTriggerCondition `json:"anyMatch"`          // 任意匹配条件（OR 逻辑）
	SubMatch          []*SAutomationTriggerCondition `json:"subMatch"`          // 子匹配条件
	SubMatchAll       *bool                          `json:"subMatchAll"`       // 子匹配是否全部满足
	ExecutionInterval int                            `json:"executionInterval"` // 执行间隔（秒），0表示实时执行
}

// IsDeviceCondition 判断是否为设备值触发条件
func (c *SAutomationTriggerCondition) IsDeviceCondition() bool {
	return c.DeviceCondition != nil && c.DeviceCondition.DeviceId != "" && c.DeviceCondition.Rule != ""
	// && c.DeviceCondition.From != "" // 暂时不验证
}

// IsTimeCondition 判断是否为时间触发条件
func (c *SAutomationTriggerCondition) IsTimeCondition() bool {
	return c.TimeCondition != nil
}

// HasAnyCondition 判断是否至少有一种触发条件
func (c *SAutomationTriggerCondition) HasAnyCondition() bool {
	return c.IsDeviceCondition() || c.IsTimeCondition()
}

// Validate 验证触发条件
func (c *SAutomationTriggerCondition) Validate() error {
	// 必须至少有一种触发条件
	if !c.HasAnyCondition() {
		return fmt.Errorf("触发条件不能为空，必须指定设备条件或时间条件")
	}

	// 验证设备条件
	if c.IsDeviceCondition() {
		if err := c.DeviceCondition.Validate(); err != nil {
			return err
		}
	}

	// 验证时间条件
	if c.IsTimeCondition() {
		if err := c.TimeCondition.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Validate 验证触发配置
func (tc *SAutomationTriggerConfig) Validate() error {
	// 验证执行间隔
	if tc.ExecutionInterval < 0 {
		return fmt.Errorf("执行间隔不能为负数")
	}

	// 验证任意匹配条件
	if len(tc.AnyMatch) == 0 {
		return fmt.Errorf("任意匹配条件不能为空")
	}

	for i, condition := range tc.AnyMatch {
		if err := condition.Validate(); err != nil {
			return fmt.Errorf("任意匹配条件[%d]验证失败: %v", i, err)
		}
	}

	// 验证子匹配条件
	for i, condition := range tc.SubMatch {
		if err := condition.Validate(); err != nil {
			return fmt.Errorf("子匹配条件[%d]验证失败: %v", i, err)
		}
	}

	return nil
}

// IsRealTimeExecution 判断是否为实时执行
func (tc *SAutomationTriggerConfig) IsRealTimeExecution() bool {
	return tc.ExecutionInterval == 0
}

// GetExecutionInterval 获取执行间隔
func (tc *SAutomationTriggerConfig) GetExecutionInterval() int {
	return tc.ExecutionInterval
}

// Validate 验证设备条件
func (dc *SAutomationDeviceCondition) Validate() error {
	if dc.DeviceId == "" {
		return fmt.Errorf("设备ID不能为空")
	}
	//if dc.From == "" {
	//	return fmt.Errorf("数据来源不能为空")
	//}
	if dc.Rule == "" {
		return fmt.Errorf("规则表达式不能为空")
	}
	return nil
}

// Validate 验证时间条件
func (tc *SAutomationTimeCondition) Validate() error {
	// 验证小时
	if tc.Hour != nil {
		if *tc.Hour < 0 || *tc.Hour > 23 {
			return fmt.Errorf("小时必须在 0-23 之间")
		}
	}

	// 验证分钟
	if tc.Minute != nil {
		if *tc.Minute < 0 || *tc.Minute > 59 {
			return fmt.Errorf("分钟必须在 0-59 之间")
		}
	}

	// 验证星期几
	if tc.DayOfWeek != nil {
		if *tc.DayOfWeek < 0 || *tc.DayOfWeek > 6 {
			return fmt.Errorf("星期几必须在 0-6 之间 (0=周日)")
		}
	}

	// 验证月的第几天
	if tc.DayOfMonth != nil {
		if *tc.DayOfMonth < 1 || *tc.DayOfMonth > 31 {
			return fmt.Errorf("月的第几天必须在 1-31 之间")
		}
	}

	// 验证月份
	if tc.Month != nil {
		if *tc.Month < 1 || *tc.Month > 12 {
			return fmt.Errorf("月份必须在 1-12 之间")
		}
	}

	// 验证时间范围
	if tc.StartTime != "" && tc.EndTime != "" {
		if err := tc.validateTimeRange(); err != nil {
			return err
		}
	}

	return nil
}

// validateTimeRange 验证时间范围格式
func (tc *SAutomationTimeCondition) validateTimeRange() error {
	// 验证开始时间格式
	if _, err := time.Parse("15:04", tc.StartTime); err != nil {
		return fmt.Errorf("开始时间格式错误，应为 HH:MM 格式")
	}

	// 验证结束时间格式
	if _, err := time.Parse("15:04", tc.EndTime); err != nil {
		return fmt.Errorf("结束时间格式错误，应为 HH:MM 格式")
	}

	return nil
}

// IsTimeMatch 判断当前时间是否匹配时间条件
func (tc *SAutomationTimeCondition) IsTimeMatch(now time.Time) bool {
	// 检查小时
	if tc.Hour != nil && now.Hour() != *tc.Hour {
		return false
	}

	// 检查分钟
	if tc.Minute != nil && now.Minute() != *tc.Minute {
		return false
	}

	// 检查星期几
	if tc.DayOfWeek != nil && int(now.Weekday()) != *tc.DayOfWeek {
		return false
	}

	// 检查月的第几天
	if tc.DayOfMonth != nil && now.Day() != *tc.DayOfMonth {
		return false
	}

	// 检查月份
	if tc.Month != nil && int(now.Month()) != *tc.Month {
		return false
	}

	// 检查时间范围
	if tc.StartTime != "" && tc.EndTime != "" {
		if !tc.isInTimeRange(now) {
			return false
		}
	}

	return true
}

// isInTimeRange 判断是否在时间范围内
func (tc *SAutomationTimeCondition) isInTimeRange(now time.Time) bool {
	nowTime := now.Format("15:04")

	// 如果开始时间小于结束时间，正常范围判断
	if tc.StartTime <= tc.EndTime {
		return nowTime >= tc.StartTime && nowTime <= tc.EndTime
	}

	// 如果开始时间大于结束时间，跨天范围判断
	return nowTime >= tc.StartTime || nowTime <= tc.EndTime
}
