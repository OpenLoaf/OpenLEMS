package p_energy_manage

import (
	"p_energy_manage/internal"
	"time"
)

// SDateRange 日期范围配置
type SDateRange struct {
	StartDate  string `json:"startDate" v:"required|date"`
	EndDate    string `json:"endDate" v:"date"`
	IsLongTerm bool   `json:"isLongTerm" v:"required"`
}

// STimeRange 时间范围配置
type STimeRange struct {
	Type         string `json:"type" v:"required|in:weekday,custom,monthly"`
	WeekdayType  string `json:"weekdayType" v:"in:workday,weekend,all"`
	CustomDays   []int  `json:"customDays"`
	CustomMonths []int  `json:"customMonths"`
}

// SStrategyConfig 策略配置
type SStrategyConfig struct {
	SocMinRatio              float64  `json:"socMinRatio" v:"required|between:0,100"`
	SocMaxRatio              float64  `json:"socMaxRatio" v:"required|between:0,100"`
	EnableHealthOptimization bool     `json:"enableHealthOptimization" v:"required"`
	MonthlyChargeDay         int      `json:"monthlyChargeDay" v:"between:1,28"`
	Points                   [][2]int `json:"points"`
}

// IsActive 判断在给定时间点是否命中日期/时间范围
// 与内部实现一致：日期范围 + 时间范围 全部满足才生效
func IsActive(now time.Time, dr *SDateRange, tr *STimeRange) bool {
	// 转换为内部类型并复用内部判断逻辑
	var idr *internal.SDateRange
	if dr != nil {
		idr = &internal.SDateRange{StartDate: dr.StartDate, EndDate: dr.EndDate, IsLongTerm: dr.IsLongTerm}
	}
	var itr *internal.STimeRange
	if tr != nil {
		itr = &internal.STimeRange{Type: tr.Type, WeekdayType: tr.WeekdayType, CustomDays: tr.CustomDays, CustomMonths: tr.CustomMonths}
	}

	// 直接构造内部策略以使用其 IsActive 逻辑
	s := &internal.SEnergyManageStrategy{
		DateRangeParsed: idr,
		TimeRangeParsed: itr,
	}
	return s.IsActive(now)
}
