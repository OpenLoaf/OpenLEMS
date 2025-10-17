package s_policy

import (
	"common"
	"context"
	"s_db/s_db_model"
	"s_policy/internal"
	"time"
)

// 导出内部类型
type SDateRange = internal.SDateRange
type STimeRange = internal.STimeRange
type SStrategyConfig = internal.SStrategyConfig

// NewPolicyManager 创建策略管理器实例
func NewPolicyManager(ctx context.Context) common.IPolicyManager {
	return internal.NewPolicyManager(ctx)
}

// GetActiveEnergyStorageTask 获取当前激活的储能定时任务
func GetActiveEnergyStorageTask(ctx context.Context) (*s_db_model.SEnergyStorageModel, error) {
	return internal.GetActiveEnergyStorageTask(ctx)
}

// GetActivePolicyId 获取当前激活的策略ID
func GetActivePolicyId(ctx context.Context) string {
	return common.GetPolicyManager().GetActivePolicyId()
}

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

// ExecuteStrategy 执行储能策略
func ExecuteStrategy(ctx context.Context, targetPower int32, essDeviceIds []string) error {
	return internal.ExecuteStrategy(ctx, targetPower, essDeviceIds)
}

// ValidateStrategy 验证储能策略配置
func ValidateStrategy(dateRange *SDateRange, timeRange *STimeRange, config *SStrategyConfig) error {
	return internal.ValidateStrategy(dateRange, timeRange, config)
}
