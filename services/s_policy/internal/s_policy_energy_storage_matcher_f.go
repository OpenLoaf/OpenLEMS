package internal

import (
	"common/c_log"
	"context"
	"encoding/json"
	"s_db"
	"s_db/s_db_model"
	"sort"
	"time"

	"github.com/pkg/errors"
)

// GetActiveEnergyStorageTask 获取当前激活的储能定时任务
func GetActiveEnergyStorageTask(ctx context.Context) (*s_db_model.SEnergyStorageModel, error) {
	// 查询所有启用的储能定时任务
	allTasks, err := s_db.GetEnergyStorageService().GetEnabledEnergyStorages(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "查询启用的储能定时任务失败")
	}

	if len(allTasks) == 0 {
		return nil, errors.New("没有找到启用的储能定时任务")
	}

	// 过滤出匹配当前时间的任务
	now := time.Now()
	var matchedTasks []*s_db_model.SEnergyStorageModel

	for _, task := range allTasks {
		if isTaskActive(ctx, task, now) {
			matchedTasks = append(matchedTasks, task)
		}
	}

	if len(matchedTasks) == 0 {
		return nil, errors.New("没有找到匹配当前时间的储能定时任务")
	}

	// 按优先级（降序）和创建时间（降序）排序
	sort.Slice(matchedTasks, func(i, j int) bool {
		// 优先级高的在前
		if matchedTasks[i].Priority != matchedTasks[j].Priority {
			return matchedTasks[i].Priority > matchedTasks[j].Priority
		}
		// 创建时间晚的在前
		if matchedTasks[i].CreatedAt != nil && matchedTasks[j].CreatedAt != nil {
			return matchedTasks[i].CreatedAt.After(matchedTasks[j].CreatedAt)
		}
		return false
	})

	activeTask := matchedTasks[0]
	c_log.Infof(ctx, "匹配到储能定时任务: ID=%d, Name=%s, Priority=%d", activeTask.Id, activeTask.Name, activeTask.Priority)

	return activeTask, nil
}

// isTaskActive 判断任务是否在指定时间激活
func isTaskActive(ctx context.Context, task *s_db_model.SEnergyStorageModel, now time.Time) bool {
	// 解析日期范围
	var dateRange SDateRange
	if err := json.Unmarshal([]byte(task.DateRange), &dateRange); err != nil {
		c_log.Warningf(ctx, "解析任务 %d 的日期范围失败: %v", task.Id, err)
		return false
	}

	// 检查日期范围
	if !dateRange.IsLongTerm {
		// 检查开始日期
		if dateRange.StartDate != "" {
			startDate, err := time.Parse("2006-01-02", dateRange.StartDate)
			if err == nil && now.Before(startDate) {
				return false
			}
		}

		// 检查结束日期
		if dateRange.EndDate != "" {
			endDate, err := time.Parse("2006-01-02", dateRange.EndDate)
			if err == nil {
				// 结束日期应包含当天的23:59:59
				endDate = endDate.Add(24*time.Hour - time.Second)
				if now.After(endDate) {
					return false
				}
			}
		}
	}

	// 解析时间范围
	var timeRange STimeRange
	if err := json.Unmarshal([]byte(task.TimeRange), &timeRange); err != nil {
		c_log.Warningf(ctx, "解析任务 %d 的时间范围失败: %v", task.Id, err)
		return false
	}

	// 检查时间范围
	switch timeRange.Type {
	case ETimeRangeTypeWeekday:
		// 检查工作日/周末
		weekday := now.Weekday()
		isWeekend := weekday == time.Saturday || weekday == time.Sunday

		if timeRange.WeekdayType == EWeekdayTypeWorkday && isWeekend {
			return false
		}
		if timeRange.WeekdayType == EWeekdayTypeWeekend && !isWeekend {
			return false
		}

	case ETimeRangeTypeCustom:
		// 检查自定义时间段（使用CustomDays和CustomMonths字段）
		if len(timeRange.CustomDays) > 0 {
			currentDay := now.Day()
			dayMatched := false
			for _, day := range timeRange.CustomDays {
				if day == currentDay {
					dayMatched = true
					break
				}
			}
			if !dayMatched {
				return false
			}
		}

		if len(timeRange.CustomMonths) > 0 {
			currentMonth := int(now.Month())
			monthMatched := false
			for _, month := range timeRange.CustomMonths {
				if month == currentMonth {
					monthMatched = true
					break
				}
			}
			if !monthMatched {
				return false
			}
		}

	case ETimeRangeTypeMonthly:
		// 月度时间范围（使用CustomDays字段作为每月的日期）
		if len(timeRange.CustomDays) > 0 {
			currentDay := now.Day()
			dayMatched := false
			for _, day := range timeRange.CustomDays {
				if day == currentDay {
					dayMatched = true
					break
				}
			}
			if !dayMatched {
				return false
			}
		}
	}

	return true
}
