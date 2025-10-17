package p_energy_manage

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
