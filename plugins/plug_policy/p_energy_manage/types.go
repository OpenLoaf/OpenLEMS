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
	// 直接构造内部策略以使用其 IsActive 逻辑（类型别名可直接复用）
	s := &internal.SEnergyManageStrategy{
		DateRangeParsed: dr,
		TimeRangeParsed: tr,
	}
	return s.IsActive(now)
}
