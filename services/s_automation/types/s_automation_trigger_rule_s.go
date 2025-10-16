package types

import (
	"fmt"
	"time"
)

// SAutomationDeviceCondition 设备触发条件结构体
type SAutomationDeviceCondition struct {
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
	From     string `json:"from" dc:"是从哪里去取值"`
	Rule     string `json:"rule" v:"required" dc:"规则表达式，如 P>30, Ia<100"`
}

// SAutomationTriggerCondition 自动化触发条件结构体
type SAutomationTriggerCondition struct {
	DeviceCondition *SAutomationDeviceCondition `json:"deviceCondition,omitempty" dc:"设备触发条件"`
	TimeCondition   *SAutomationTimeCondition   `json:"timeCondition,omitempty" dc:"时间触发条件"`
}

// SAutomationTimeCondition 时间触发条件结构体
type SAutomationTimeCondition struct {
	Hour       *int   `json:"hour,omitempty" v:"between:0,23" dc:"小时 (0-23)，nil 表示不限制"`
	Minute     *int   `json:"minute,omitempty" v:"between:0,59" dc:"分钟 (0-59)，nil 表示不限制"`
	DayOfWeek  *int   `json:"dayOfWeek,omitempty" v:"between:0,6" dc:"星期几 (0-6，0=周日)，nil 表示不限制"`
	DayOfMonth *int   `json:"dayOfMonth,omitempty" v:"between:1,31" dc:"月的第几天 (1-31)，nil 表示不限制"`
	Month      *int   `json:"month,omitempty" v:"between:1,12" dc:"月份 (1-12)，nil 表示不限制"`
	StartTime  string `json:"startTime,omitempty" dc:"开始时间 (HH:MM 格式)"`
	EndTime    string `json:"endTime,omitempty" dc:"结束时间 (HH:MM 格式)"`
}

// SAutomationTriggerConfig 自动化触发配置结构体
type SAutomationTriggerConfig struct {
	AnyMatch          []*SAutomationTriggerCondition `json:"anyMatch" v:"required|min-length:1" dc:"任意匹配条件（OR 逻辑）"`
	SubMatch          []*SAutomationTriggerCondition `json:"subMatch" dc:"子匹配条件"`
	SubMatchAll       *bool                          `json:"subMatchAll" dc:"子匹配是否全部满足"`
	ExecutionInterval int                            `json:"executionInterval" v:"min:0" dc:"执行间隔（秒），0表示实时执行"`
}

func (c *SAutomationTriggerCondition) IsDeviceCondition() bool {
	return c.DeviceCondition != nil && c.DeviceCondition.DeviceId != "" && c.DeviceCondition.Rule != ""
}

func (c *SAutomationTriggerCondition) IsTimeCondition() bool {
	return c.TimeCondition != nil
}

func (c *SAutomationTriggerCondition) HasAnyCondition() bool {
	return c.IsDeviceCondition() || c.IsTimeCondition()
}

func (c *SAutomationTriggerCondition) Validate() error {
	if !c.HasAnyCondition() {
		return fmt.Errorf("触发条件不能为空，必须指定设备条件或时间条件")
	}
	if c.IsDeviceCondition() {
		if err := c.DeviceCondition.Validate(); err != nil {
			return err
		}
	}
	if c.IsTimeCondition() {
		if err := c.TimeCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (tc *SAutomationTriggerConfig) Validate() error {
	if tc.ExecutionInterval < 0 {
		return fmt.Errorf("执行间隔不能为负数")
	}
	if len(tc.AnyMatch) == 0 {
		return fmt.Errorf("任意匹配条件不能为空")
	}
	for i, condition := range tc.AnyMatch {
		if err := condition.Validate(); err != nil {
			return fmt.Errorf("任意匹配条件[%d]验证失败: %v", i, err)
		}
	}
	for i, condition := range tc.SubMatch {
		if err := condition.Validate(); err != nil {
			return fmt.Errorf("子匹配条件[%d]验证失败: %v", i, err)
		}
	}
	return nil
}

func (tc *SAutomationTriggerConfig) IsRealTimeExecution() bool {
	return tc.ExecutionInterval == 0
}

func (tc *SAutomationTriggerConfig) GetExecutionInterval() int {
	return tc.ExecutionInterval
}

func (dc *SAutomationDeviceCondition) Validate() error {
	if dc.DeviceId == "" {
		return fmt.Errorf("设备ID不能为空")
	}
	if dc.Rule == "" {
		return fmt.Errorf("规则表达式不能为空")
	}
	return nil
}

func (tc *SAutomationTimeCondition) Validate() error {
	if tc.Hour != nil {
		if *tc.Hour < 0 || *tc.Hour > 23 {
			return fmt.Errorf("小时必须在 0-23 之间")
		}
	}
	if tc.Minute != nil {
		if *tc.Minute < 0 || *tc.Minute > 59 {
			return fmt.Errorf("分钟必须在 0-59 之间")
		}
	}
	if tc.DayOfWeek != nil {
		if *tc.DayOfWeek < 0 || *tc.DayOfWeek > 6 {
			return fmt.Errorf("星期几必须在 0-6 之间 (0=周日)")
		}
	}
	if tc.DayOfMonth != nil {
		if *tc.DayOfMonth < 1 || *tc.DayOfMonth > 31 {
			return fmt.Errorf("月的第几天必须在 1-31 之间")
		}
	}
	if tc.Month != nil {
		if *tc.Month < 1 || *tc.Month > 12 {
			return fmt.Errorf("月份必须在 1-12 之间")
		}
	}
	if tc.StartTime != "" && tc.EndTime != "" {
		if err := tc.validateTimeRange(); err != nil {
			return err
		}
	}
	return nil
}

func (tc *SAutomationTimeCondition) validateTimeRange() error {
	if _, err := time.Parse("15:04", tc.StartTime); err != nil {
		return fmt.Errorf("开始时间格式错误，应为 HH:MM 格式")
	}
	if _, err := time.Parse("15:04", tc.EndTime); err != nil {
		return fmt.Errorf("结束时间格式错误，应为 HH:MM 格式")
	}
	return nil
}

func (tc *SAutomationTimeCondition) IsTimeMatch(now time.Time) bool {
	if tc.Hour != nil && now.Hour() != *tc.Hour {
		return false
	}
	if tc.Minute != nil && now.Minute() != *tc.Minute {
		return false
	}
	if tc.DayOfWeek != nil && int(now.Weekday()) != *tc.DayOfWeek {
		return false
	}
	if tc.DayOfMonth != nil && now.Day() != *tc.DayOfMonth {
		return false
	}
	if tc.Month != nil && int(now.Month()) != *tc.Month {
		return false
	}
	if tc.StartTime != "" && tc.EndTime != "" {
		if !tc.isInTimeRange(now) {
			return false
		}
	}
	return true
}

func (tc *SAutomationTimeCondition) isInTimeRange(now time.Time) bool {
	nowTime := now.Format("15:04")
	if tc.StartTime <= tc.EndTime {
		return nowTime >= tc.StartTime && nowTime <= tc.EndTime
	}
	return nowTime >= tc.StartTime || nowTime <= tc.EndTime
}
