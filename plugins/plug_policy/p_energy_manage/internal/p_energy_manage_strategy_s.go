package internal

import (
	"encoding/json"
	"time"

	"s_db/s_db_model"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/pkg/errors"
)

// SEnergyManageStrategy 储能策略包装结构体
type SEnergyManageStrategy struct {
	*s_db_model.SEnergyStorageStrategyModel
	DateRangeParsed  *SDateRange
	TimeRangeParsed  *STimeRange
	ConfigParsed     *SStrategyConfig
	EssDeviceIdsList []string
}

// IsActive 判断策略是否在当前时间生效
func (s *SEnergyManageStrategy) IsActive(now time.Time) bool {
	// 1. 检查日期范围
	if !s.isInDateRange(now) {
		return false
	}

	// 2. 检查时间范围
	if !s.isInTimeRange(now) {
		return false
	}

	return true
}

// isInDateRange 日期范围判断
func (s *SEnergyManageStrategy) isInDateRange(now time.Time) bool {
	if s.DateRangeParsed == nil {
		return true // 没有日期限制，默认生效
	}

	// 长期有效
	if s.DateRangeParsed.IsLongTerm {
		return true
	}

	// 检查开始日期
	if s.DateRangeParsed.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", s.DateRangeParsed.StartDate)
		if err != nil {
			return false
		}
		if now.Before(startDate) {
			return false
		}
	}

	// 检查结束日期
	if s.DateRangeParsed.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", s.DateRangeParsed.EndDate)
		if err != nil {
			return false
		}
		if now.After(endDate) {
			return false
		}
	}

	return true
}

// isInTimeRange 时间范围判断（处理三种类型）
func (s *SEnergyManageStrategy) isInTimeRange(now time.Time) bool {
	if s.TimeRangeParsed == nil {
		return true // 没有时间限制，默认生效
	}

	switch s.TimeRangeParsed.Type {
	case "weekday":
		return s.isWeekdayMatch(now)
	case "custom":
		return s.isCustomDayMatch(now)
	case "monthly":
		return s.isMonthlyMatch(now)
	default:
		return true // 未知类型，默认生效
	}
}

// isWeekdayMatch 工作日/周末匹配
func (s *SEnergyManageStrategy) isWeekdayMatch(now time.Time) bool {
	weekday := now.Weekday()
	isWeekend := weekday == time.Saturday || weekday == time.Sunday

	switch s.TimeRangeParsed.WeekdayType {
	case "workday":
		return !isWeekend
	case "weekend":
		return isWeekend
	case "all":
		return true
	default:
		return true
	}
}

// isCustomDayMatch 自定义周内日匹配
func (s *SEnergyManageStrategy) isCustomDayMatch(now time.Time) bool {
	if len(s.TimeRangeParsed.CustomDays) == 0 {
		return true
	}

	weekday := int(now.Weekday()) // 0=周日, 1=周一, ..., 6=周六
	for _, day := range s.TimeRangeParsed.CustomDays {
		if day == weekday {
			return true
		}
	}
	return false
}

// isMonthlyMatch 指定月份匹配
func (s *SEnergyManageStrategy) isMonthlyMatch(now time.Time) bool {
	if len(s.TimeRangeParsed.CustomMonths) == 0 {
		return true
	}

	month := int(now.Month())
	for _, m := range s.TimeRangeParsed.CustomMonths {
		if m == month {
			return true
		}
	}
	return false
}

// ParseStrategy 解析策略模型为储能策略
func ParseStrategy(model *s_db_model.SEnergyStorageStrategyModel) (*SEnergyManageStrategy, error) {
	strategy := &SEnergyManageStrategy{
		SEnergyStorageStrategyModel: model,
	}

	// 解析日期范围
	if model.DateRange != "" {
		var dateRange SDateRange
		if err := gjson.DecodeTo(model.DateRange, &dateRange); err != nil {
			return nil, errors.Wrap(err, "解析日期范围失败")
		}
		strategy.DateRangeParsed = &dateRange
	}

	// 解析时间范围
	if model.TimeRange != "" {
		var timeRange STimeRange
		if err := gjson.DecodeTo(model.TimeRange, &timeRange); err != nil {
			return nil, errors.Wrap(err, "解析时间范围失败")
		}
		strategy.TimeRangeParsed = &timeRange
	}

	// 解析策略配置
	if model.Config != "" {
		var config SStrategyConfig
		if err := gjson.DecodeTo(model.Config, &config); err != nil {
			return nil, errors.Wrap(err, "解析策略配置失败")
		}
		strategy.ConfigParsed = &config
	}

	// 解析设备ID列表
	if model.EssDeviceIds != "" {
		var deviceIds []string
		if err := json.Unmarshal([]byte(model.EssDeviceIds), &deviceIds); err != nil {
			return nil, errors.Wrap(err, "解析设备ID列表失败")
		}
		strategy.EssDeviceIdsList = deviceIds
	}

	return strategy, nil
}
