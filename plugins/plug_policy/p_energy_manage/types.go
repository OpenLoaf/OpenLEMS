package p_energy_manage

import (
	"p_energy_manage/internal"
	"time"
)

// SDateRange 日期范围配置（类型别名到 internal，避免重复定义）
type SDateRange = internal.SDateRange

// STimeRange 时间范围配置（类型别名到 internal，避免重复定义）
type STimeRange = internal.STimeRange

// SStrategyConfig 策略配置（类型别名到 internal，避免重复定义）
type SStrategyConfig = internal.SStrategyConfig

// IsActive 判断在给定时间点是否命中日期/时间范围
// 与内部实现一致：日期范围 + 时间范围 全部满足才生效
func IsActive(now time.Time, dr *SDateRange, tr *STimeRange) bool {
	// 简化的时间范围判断逻辑
	if dr == nil && tr == nil {
		return true
	}

	// 日期范围判断
	if dr != nil {
		if !dr.IsLongTerm && dr.StartDate != "" {
			if startDate, err := time.Parse("2006-01-02", dr.StartDate); err == nil && now.Before(startDate) {
				return false
			}
		}
		if !dr.IsLongTerm && dr.EndDate != "" {
			if endDate, err := time.Parse("2006-01-02", dr.EndDate); err == nil && now.After(endDate) {
				return false
			}
		}
	}

	// 时间范围判断（简化版）
	if tr != nil {
		switch tr.Type {
		case internal.ETimeRangeTypeWeekday:
			// 工作日判断
			if tr.WeekdayType == internal.EWeekdayTypeWorkday && (now.Weekday() == time.Saturday || now.Weekday() == time.Sunday) {
				return false
			}
			if tr.WeekdayType == internal.EWeekdayTypeWeekend && now.Weekday() != time.Saturday && now.Weekday() != time.Sunday {
				return false
			}
		}
	}

	return true
}
