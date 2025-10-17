package internal

// 时间范围类型（字符串枚举）
type ETimeRangeType string

const (
	ETimeRangeTypeWeekday ETimeRangeType = "weekday"
	ETimeRangeTypeCustom  ETimeRangeType = "custom"
	ETimeRangeTypeMonthly ETimeRangeType = "monthly"
)

// 工作日类型（字符串枚举）
type EWeekdayType string

const (
	EWeekdayTypeWorkday EWeekdayType = "workday"
	EWeekdayTypeWeekend EWeekdayType = "weekend"
	EWeekdayTypeAll     EWeekdayType = "all"
)

// SDateRange 日期范围配置
type SDateRange struct {
	StartDate  string `json:"startDate" v:"required|date"`
	EndDate    string `json:"endDate" v:"date"`
	IsLongTerm bool   `json:"isLongTerm" v:"required"`
}

// STimeRange 时间范围配置
type STimeRange struct {
	Type         ETimeRangeType `json:"type" v:"required|in:weekday,custom,monthly"`
	WeekdayType  EWeekdayType   `json:"weekdayType" v:"in:workday,weekend,all"`
	CustomDays   []int          `json:"customDays"`
	CustomMonths []int          `json:"customMonths"`
}

// SStrategyConfig 策略配置
type SStrategyConfig struct {
	Points [][2]int `json:"points"`
}
