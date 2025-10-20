package internal

import (
	"common/c_enum"
	"time"
)

// SPriceSegment 电价时段结构体
type SPriceSegment struct {
	StartTime string            `json:"startTime"` // 开始时间 "HH:MM" 格式
	EndTime   string            `json:"endTime"`   // 结束时间 "HH:MM" 格式
	PriceType c_enum.EPriceType `json:"priceType"` // 电价类型
	Price     float64           `json:"price"`     // 电价值
}

// SDateRange 日期范围配置
type SDateRange struct {
	StartDate  string `json:"startDate"`  // 开始日期 "YYYY-MM-DD" 格式
	EndDate    string `json:"endDate"`    // 结束日期 "YYYY-MM-DD" 格式
	IsLongTerm bool   `json:"isLongTerm"` // 是否长期有效
}

// STimeRange 时间范围配置
type STimeRange struct {
	Type         string `json:"type"`         // 时间范围类型：weekday/custom/monthly
	WeekdayType  string `json:"weekdayType"`  // 工作日类型：workday/weekend/all
	CustomDays   []int  `json:"customDays"`   // 自定义日期
	CustomMonths []int  `json:"customMonths"` // 自定义月份
}

// SPriceInfo 电价信息结构体
type SPriceInfo struct {
	Id            int              `json:"id"`            // 电价ID
	Description   string           `json:"description"`   // 电价描述
	Priority      int              `json:"priority"`      // 优先级
	Status        c_enum.EStatus   `json:"status"`        // 状态
	IsActive      bool             `json:"isActive"`      // 是否激活
	DateRange     *SDateRange      `json:"dateRange"`     // 日期范围
	TimeRange     *STimeRange      `json:"timeRange"`     // 时间范围
	PriceSegments []*SPriceSegment `json:"priceSegments"` // 电价时段
	RemoteId      *string          `json:"remoteId"`      // 远程电价ID
	CreatedAt     *time.Time       `json:"createdAt"`     // 创建时间
	UpdatedAt     *time.Time       `json:"updatedAt"`     // 更新时间
	CreatedBy     string           `json:"createdBy"`     // 创建人
}

// SCurrentPrice 当前电价信息（用于Storage保存）
type SCurrentPrice struct {
	PriceId   int     `json:"priceId"`   // 电价ID
	Price     float64 `json:"price"`     // 当前电价
	PriceType string  `json:"priceType"` // 电价类型
	Timestamp int64   `json:"timestamp"` // 时间戳
}
