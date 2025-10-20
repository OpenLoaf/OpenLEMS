package internal

import (
	"common/c_enum"
	"common/c_log"
	"context"
	"sort"
	"time"
)

// GetActivePriceAtTime 获取指定时间激活的电价
func GetActivePriceAtTime(ctx context.Context, targetTime time.Time) (*SPriceInfo, error) {
	cache := GetPriceCache()

	// 获取所有启用的电价
	enabledPrices := cache.GetEnabledPrices()
	if len(enabledPrices) == 0 {
		c_log.Debug(ctx, "没有启用的电价配置")
		return nil, nil
	}

	// 过滤出在指定时间激活的电价
	var activePrices []*SPriceInfo
	for _, price := range enabledPrices {
		if isActiveAtTime(targetTime, price.DateRange, price.TimeRange) {
			activePrices = append(activePrices, price)
		}
	}

	if len(activePrices) == 0 {
		c_log.Debugf(ctx, "在时间 %s 没有激活的电价配置", targetTime.Format("2006-01-02 15:04:05"))
		return nil, nil
	}

	// 按优先级排序（数值越小优先级越高）
	sort.Slice(activePrices, func(i, j int) bool {
		return activePrices[i].Priority < activePrices[j].Priority
	})

	// 返回优先级最高的电价
	activePrice := activePrices[0]

	// 更新IsActive状态
	activePrice.IsActive = true

	c_log.Debugf(ctx, "找到激活的电价 - ID: %d, 描述: %s, 优先级: %d",
		activePrice.Id, activePrice.Description, activePrice.Priority)

	return activePrice, nil
}

// GetCurrentActivePrice 获取当前激活的电价
func GetCurrentActivePrice(ctx context.Context) (*SPriceInfo, error) {
	return GetActivePriceAtTime(ctx, time.Now())
}

// GetCurrentPriceSegment 获取当前时间对应的电价时段
func GetCurrentPriceSegment(ctx context.Context) (*SPriceSegment, error) {
	activePrice, err := GetCurrentActivePrice(ctx)
	if err != nil {
		return nil, err
	}

	if activePrice == nil {
		return nil, nil
	}

	now := time.Now()
	currentTime := now.Format("15:04")

	// 查找当前时间对应的电价时段
	for _, segment := range activePrice.PriceSegments {
		if isTimeInRange(currentTime, segment.StartTime, segment.EndTime) {
			c_log.Debugf(ctx, "当前电价时段 - 类型: %s, 价格: %.4f",
				segment.PriceType, segment.Price)
			return segment, nil
		}
	}

	c_log.Debugf(ctx, "当前时间 %s 没有匹配的电价时段", currentTime)
	return nil, nil
}

// isActiveAtTime 判断在给定时间点是否命中日期/时间范围
func isActiveAtTime(targetTime time.Time, dr *SDateRange, tr *STimeRange) bool {
	// 简化的时间范围判断逻辑
	if dr == nil && tr == nil {
		return true
	}

	// 日期范围判断
	if dr != nil {
		if !dr.IsLongTerm && dr.StartDate != "" {
			if startDate, err := time.Parse("2006-01-02", dr.StartDate); err == nil && targetTime.Before(startDate) {
				return false
			}
		}
		if !dr.IsLongTerm && dr.EndDate != "" {
			if endDate, err := time.Parse("2006-01-02", dr.EndDate); err == nil && targetTime.After(endDate) {
				return false
			}
		}
	}

	// 时间范围判断（简化版）
	if tr != nil {
		switch tr.Type {
		case "weekday":
			// 工作日判断
			if tr.WeekdayType == "workday" && (targetTime.Weekday() == time.Saturday || targetTime.Weekday() == time.Sunday) {
				return false
			}
			if tr.WeekdayType == "weekend" && targetTime.Weekday() != time.Saturday && targetTime.Weekday() != time.Sunday {
				return false
			}
		case "custom":
			// 自定义日期判断
			if len(tr.CustomDays) > 0 {
				day := targetTime.Day()
				found := false
				for _, customDay := range tr.CustomDays {
					if customDay == day {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
			// 自定义月份判断
			if len(tr.CustomMonths) > 0 {
				month := int(targetTime.Month())
				found := false
				for _, customMonth := range tr.CustomMonths {
					if customMonth == month {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
		case "monthly":
			// 月度判断（简化版，每月1日生效）
			if targetTime.Day() != 1 {
				return false
			}
		}
	}

	return true
}

// isTimeInRange 判断时间是否在指定范围内
func isTimeInRange(currentTime, startTime, endTime string) bool {
	// 解析时间
	current, err1 := time.Parse("15:04", currentTime)
	start, err2 := time.Parse("15:04", startTime)
	end, err3 := time.Parse("15:04", endTime)

	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}

	// 处理跨天的情况（如 23:00-07:00）
	if start.After(end) {
		// 跨天情况：当前时间 >= 开始时间 或 当前时间 <= 结束时间
		return !current.Before(start) || !current.After(end)
	} else {
		// 同天情况：开始时间 <= 当前时间 <= 结束时间
		return !current.Before(start) && !current.After(end)
	}
}

// RefreshAllPricesActiveStatus 刷新所有电价的激活状态
func RefreshAllPricesActiveStatus(ctx context.Context) error {
	cache := GetPriceCache()
	now := time.Now()

	// 获取所有电价
	allPrices := cache.GetAllPrices()

	// 更新激活状态
	for _, price := range allPrices {
		price.IsActive = isActiveAtTime(now, price.DateRange, price.TimeRange) && price.Status == c_enum.EStatusEnable
	}

	c_log.Debugf(ctx, "已刷新 %d 条电价的激活状态", len(allPrices))
	return nil
}
